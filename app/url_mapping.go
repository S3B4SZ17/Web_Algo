package app

import (
	"github.com/S3B4SZ17/Web_Algo/controllers"
)

func mapUrls(){
	router.GET("/ping", controllers.Ping)
	router.GET("/home", controllers.Home)
	router.POST("/two_sums", controllers.TwoSums)
}