package app

import (
	"github.com/S3B4SZ17/Web_Algo/controllers"
	"github.com/S3B4SZ17/Web_Algo/middleware"
)

func mapUrls() {

	public := router.Group("/api")
	public.GET("/ping", controllers.Ping)
	public.GET("/callback-gl", controllers.CallBackFromGoogle)
	public.GET("/login-gl", controllers.HandleGoogleLogin)
	public.GET("/home", controllers.Home)

	protected := router.Group("/api/authorized")
	protected.Use(middleware.Oauth2AuthMiddleware())
	protected.POST("/two_sums", controllers.TwoSums)
	protected.GET("/userinfo", controllers.GetUserInfo)
	protected.POST("/sendEmail", controllers.SendEmail)
	protected.POST("/reverse_linkedList", controllers.ReverseLinkedList)
	protected.POST("/reverse", controllers.Reverse)
	protected.POST("/solve_algo", controllers.Solve_algorithm)
	protected.GET("/end_session", controllers.RemoveUserFromSessions)
}
