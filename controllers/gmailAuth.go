package controllers

import (
	"net/http"

	"github.com/S3B4SZ17/Web_Algo/management"
	"github.com/S3B4SZ17/Web_Algo/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

var (
	oauthConfGl        = &oauth2.Config{}
	oauthStateStringGl = ""
)

/*
InitializeOAuthGoogle Function
*/
func InitializeOAuthGoogle() {
	oauthConfGl.ClientID = viper.GetString("google.client_id")
	oauthConfGl.ClientSecret = viper.GetString("google.client_secret")
	oauthConfGl.RedirectURL = viper.GetString("google.redirect_uri")
	oauthConfGl.Scopes = []string{gmail.GmailReadonlyScope, "https://www.googleapis.com/auth/userinfo.email", gmail.GmailSendScope}
	oauthConfGl.Endpoint = google.Endpoint
	oauthStateStringGl = viper.GetString("oauthStateString")
	management.Log.Info("Oauth Config", zap.String("redirect_uri", oauthConfGl.RedirectURL))
}

/*
HandleGoogleLogin Function
*/
func HandleGoogleLogin(c *gin.Context) {
	InitializeOAuthGoogle()
	services.HandleLogin(c, oauthConfGl, &oauthStateStringGl)
}

/*
CallBackFromGoogle Function
*/
func CallBackFromGoogle(c *gin.Context) {
	management.Log.Info("Callback-gl..")
	services.AuthenticateUser(c, oauthConfGl, &oauthStateStringGl)
}

func RemoveUserFromSessions(c *gin.Context) {
	email := c.Request.Header.Get("user_email")
	management.Log.Info("Trying to remove user" + email)
	services.RemoveUserFromSessions(email)
}

func SendEmail(c *gin.Context) {

	var email_to *services.EmailTo

	//using BindJson method to serialize body with struct
	if err := c.ShouldBindWith(&email_to, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(415, gin.H{"errcode": 415, "description": "Post Data Err"})
		return
	}

	response, err := services.EmailSrv(c, oauthConfGl, &oauthStateStringGl, email_to)

	if err != nil {
		response = &services.EmailResponse{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
