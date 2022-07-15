package app

import (
	"os"

	"github.com/gin-gonic/gin"
)
var (
	router = gin.Default()
)
func StartApp(){
	mapUrls()
	
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8181"
	}

	router.Run(":" + httpPort)
}