package app

import (
	"net"
	"os"

	mgt "github.com/S3B4SZ17/Web_Algo/management"
	pb "github.com/S3B4SZ17/Web_Algo/proto/addTwoNumbers"
	"github.com/S3B4SZ17/Web_Algo/server"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)
var (
	host = "localhost"
	addTwoNumbersPort = "50051"
	router *gin.Engine
	GIN_MODE string
)
func StartApp(){

	// Start the AddTwoNumbers server
	go StartAddTwoSumServer()

	// Start the HTTP server for the application
	StartHTTPServer()
}

func StartHTTPServer() {
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

	mgt.Info.Printf("Starting application on port %v and in %v mode\n", httpPort, gin_mode)
	router = gin.Default()
	mapUrls()

	router.Run(":" + httpPort)
}

func StartAddTwoSumServer(){
	mgt.Info.Printf("Start AddTwoNumbersServer on port %v", addTwoNumbersPort)
	
	listener, err := net.Listen("tcp", ":"+addTwoNumbersPort)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterAddTwoNumbersServer(srv, &server.AddTwoNumbersServer{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		mgt.Error.Fatalf("An error occurred while serving: %v", e)
	}
}