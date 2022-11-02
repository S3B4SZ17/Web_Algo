package controllers

import (
	"github.com/S3B4SZ17/Web_Algo/management"
	"github.com/S3B4SZ17/Web_Algo/services"
	"github.com/gin-gonic/gin"
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
	oauthConfGl.Scopes = []string{gmail.GmailReadonlyScope, "https://www.googleapis.com/auth/userinfo.email"}
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
