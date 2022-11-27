package middleware

import (
	"net/http"

	"github.com/S3B4SZ17/Web_Algo/services"
	"github.com/gin-gonic/gin"
)

func Oauth2AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := services.ValidateToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
