package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/S3B4SZ17/Web_Algo/db"
	"github.com/S3B4SZ17/Web_Algo/management"
	"github.com/S3B4SZ17/Web_Algo/models"
	pbEmail "github.com/S3B4SZ17/Web_Algo/proto/email_user"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2"
)

func SendEmail(message *pbEmail.EmailMessage) (res *pbEmail.EmailResponse, err error) {

	return res, err
}

// func GetAuthenticatedUser() (res *pbEmail.EmailUser, err error) {

// 	return res, err
// }

/*
HandleLogin Function
*/
func HandleLogin(c *gin.Context, oauthConf *oauth2.Config, oauthStateString *string) {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)

	if err != nil {
		management.Log.Error("Parse: " + err.Error())
	}
	management.Log.Info(URL.String())

	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", *oauthStateString)
	parameters.Add("access_type", "offline")
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	management.Log.Info(url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func AuthenticateUser(c *gin.Context, oauthConfGl *oauth2.Config, oauthStateStringGl *string) {
	compareState(c, oauthConfGl, oauthStateStringGl)
	checkCode(c, oauthConfGl, oauthStateStringGl)

}

func checkCode(c *gin.Context, oauthConfGl *oauth2.Config, oauthStateStringGl *string) *oauth2.Token {
	management.Log.Info("Getting code from Oauth2")
	code := c.Request.FormValue("code")

	if code == "" {
		management.Log.Warn("Code not found..")
		reason := c.Request.FormValue("error_reason")
		if reason == "user_denied" {
			management.Log.Error("User has denied Permission..")
			front_end_url := url.URL{Path: viper.GetString("front_end_url") + "/loginerror"}
			c.Redirect(http.StatusUnauthorized, front_end_url.RequestURI())
		}

	} else {
		token, err := oauthConfGl.Exchange(c, code)
		if err != nil {
			management.Log.Error("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
			front_end_url := url.URL{Path: viper.GetString("front_end_url") + "/loginerror"}
			c.Redirect(http.StatusTemporaryRedirect, front_end_url.RequestURI())
		}

		fmt.Printf("Token: %v\n", token)
		user, _, _ := ExtractUser(&token.AccessToken)
		if user == nil {
			front_end_url := url.URL{Path: viper.GetString("front_end_url") + "/loginerror"}
			c.Redirect(http.StatusTemporaryRedirect, front_end_url.RequestURI())
		}
		saveToken(token, user.GetEmail())
		IsTokenValid(user.GetEmail())
		c.SetCookie("token", url.QueryEscape(token.AccessToken), 1000, "/authorized", c.Request.URL.Hostname(), false, false)
		front_end_url := url.URL{Path: viper.GetString("front_end_url") + "/authorized/welcome"}
		c.Redirect(http.StatusTemporaryRedirect, front_end_url.RequestURI())
		return token
	}
	return nil
}

func saveToken(token *oauth2.Token, email string) {

	filter := bson.D{{Key: "email", Value: email}}
	opts := options.Update().SetUpsert(true)
	user_collection := db.GetCollection(db.DB, "users")
	user := &models.User{Email: email, Token: token.AccessToken, Expiry: token.Expiry, RefreshToken: token.RefreshToken}
	update := bson.M{"$set": user}
	result, err := user_collection.UpdateOne(context.TODO(), filter, update, opts)
	// check for errors in the insertion
	if err != nil {
		management.Log.Panic(err.Error())
	}
	// display the id of the newly inserted object
	fmt.Println(result.UpsertedID)

}

func IsTokenValid(email string) bool {
	filter := bson.D{{Key: "email", Value: email}}
	user_collection := db.GetCollection(db.DB, "users")
	var user models.User
	err := user_collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			management.Log.Warn("Your query did not match any documents")
			return false
		}
		management.Log.Error(err.Error())
	}

	t_now := time.Now()
	valid := t_now.Before(user.Expiry)
	if valid {
		management.Log.Info("The token is valid!")
		return true
	} else {
		management.Log.Warn("Token is expired!")
		return false
	}
}

// // this map stores the users sessions. For larger scale applications, you can use a database or cache for this purpose
// var sessions = map[string]session{}

// // each session contains the username of the user and the time at which it expires
// type session struct {
// 	username string
// 	expiry   time.Time
// }

// // we'll use this method later to determine if the session has expired
// func (s session) isExpired() bool {
// 	return s.expiry.Before(time.Now())
// }

func emailSrv(c *gin.Context, oauthConfGl *oauth2.Config, oauthStateStringGl *string, token *oauth2.Token) {
	gmail_client := oauthConfGl.Client(c, token)
	gmail, _ := GetGmailService(c, gmail_client)
	GetGmailLabels(gmail)
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		management.Log.Error("Get: " + err.Error() + "\n")
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	defer resp.Body.Close()

	email_user := &pbEmail.EmailUser{}

	response, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(response, &email_user)
	// Check your errors!
	if err != nil {
		management.Log.Fatal(err.Error())
	}
}

func compareState(c *gin.Context, oauthConfGl *oauth2.Config, oauthStateStringGl *string) {
	state := c.Request.FormValue("state")
	management.Log.Info(state)
	if state != *oauthStateStringGl {
		management.Log.Error("invalid oauth state, expected " + *oauthStateStringGl + ", got " + state + "\n")
		front_end_url := url.URL{Path: viper.GetString("front_end_url") + "/error"}
		c.Redirect(http.StatusUnauthorized, front_end_url.RequestURI())
	}
}
