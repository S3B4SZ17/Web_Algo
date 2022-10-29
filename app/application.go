package app

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/S3B4SZ17/Web_Algo/config"
	"github.com/S3B4SZ17/Web_Algo/db"
	"github.com/S3B4SZ17/Web_Algo/management"
	mgt "github.com/S3B4SZ17/Web_Algo/management"
	"github.com/S3B4SZ17/Web_Algo/server"
	"github.com/S3B4SZ17/Web_Algo/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	host         = "localhost"
	gRPCListener = "50051"
	router       *gin.Engine
)

func StartApp(config *config.Config) {

	// Start the gRPC server
	go StartgRPCServer()
	db.IntializeDB()
	services.IsTokenValid("sebaszh17@gmail.com")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	ended := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		management.Log.Info("Application interrupted ...")
		db.CloseClientDB()
		os.Exit(1)
		ended <- true
	}()

	// Start the HTTP server for the application
	StartHTTPServer(config)
	defer db.CloseClientDB()
}

func StartHTTPServer(config *config.Config) {
	gin_mode := os.Getenv("GIN_MODE")
	fmt.Println(gin_mode)
	if gin_mode == "" {
		gin_mode = "release"
		os.Setenv("GIN_MODE", gin_mode)
		gin.SetMode(gin.ReleaseMode)
	}
	httpPort := config.Http_server.HttpPort
	if httpPort == "" {
		httpPort = "8181"
	}

	mgt.Log.Info("Starting application ...")
	router = gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.Http_server.Cors.List_hosts,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	mapUrls()

	router.Run(":" + httpPort)
}

func StartgRPCServer() {
	mgt.Log.Info("Start gRPCListener on port ", zap.String("port", gRPCListener))

	listener, err := net.Listen("tcp", ":"+gRPCListener)
	if err != nil {
		mgt.Log.Error(err.Error())
	}

	srv := grpc.NewServer()
	server.RegisterServices(srv)

	if e := server.StartServer(srv, listener); e != nil {
		mgt.Log.Error("An error occurred while serving: " + e.Error())
	}
}
