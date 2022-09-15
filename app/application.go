package app

import (
	"fmt"
	"net"
	"os"
	"time"

	mgt "github.com/S3B4SZ17/Web_Algo/management"
	"github.com/S3B4SZ17/Web_Algo/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)
var (
	host = "localhost"
	gRPCListener = "50051"
	router *gin.Engine
)
func StartApp(){

	// Start the gRPC server
	go StartgRPCServer()

	// Start the HTTP server for the application
	StartHTTPServer()
}

func StartHTTPServer() {
	gin_mode := os.Getenv("GIN_MODE")
	fmt.Println(gin_mode)
	if gin_mode == "" {
		gin_mode = "release"
		os.Setenv("GIN_MODE", gin_mode)
		gin.SetMode(gin.ReleaseMode)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8081"
	}

	mgt.Info.Printf("Starting application on port %v and in %v mode\n", httpPort, gin_mode)
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH","POST", "DELETE", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers","*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
  	}))
	mapUrls()

	router.Run(":" + httpPort)
}

func StartgRPCServer(){
	mgt.Info.Printf("Start gRPCListener on port %v", gRPCListener)
	
	listener, err := net.Listen("tcp", ":"+gRPCListener)
	if err != nil {
		mgt.Error.Printf(err.Error())
	}

	srv := grpc.NewServer()
	server.RegisterServices(srv)
	
	if e := server.StartServer(srv, listener); e != nil {
		mgt.Error.Fatalf("An error occurred while serving: %v", e)
	}
}