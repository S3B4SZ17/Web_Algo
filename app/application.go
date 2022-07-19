package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)
var (
	router *gin.Engine
	GIN_MODE string
)
func StartApp(){

	gin_mode := os.Getenv(GIN_MODE)
	if gin_mode == "" {
		gin_mode = "release"
		os.Setenv(GIN_MODE, gin_mode)
		gin.SetMode(gin.ReleaseMode)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8181"
	}

	fmt.Printf("Starting application on port %v and in %v mode\n", httpPort, gin_mode)
	router = gin.Default()
	mapUrls()

	router.Run(":" + httpPort)
}