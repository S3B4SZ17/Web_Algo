package controllers

import (
	"net/http"

	mgt "github.com/S3B4SZ17/Web_Algo/management"
	pbAddTwoNum "github.com/S3B4SZ17/Web_Algo/proto/addTwoNumbers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	host = "localhost"
	gRPCListener = "50051"
)

func TwoSums(c *gin.Context){
	var listVals *pbAddTwoNum.ListReq


	// Set up a connection to the AddTwoNumbers server.
	conn, err := grpc.Dial(host + ":" +gRPCListener, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		mgt.Error.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pbAddTwoNum.NewAddTwoNumbersClient(conn)
	

	//using BindJson method to serialize body with struct
	if err := c.ShouldBindWith(&listVals, binding.JSON);err!=nil{
   		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(415, gin.H{"errcode": 415, "description": "Post Data Err"})
   		return
	}

	response, err := client.AddTwoNumbers(c, listVals); if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(415, gin.H{"errcode": 415, "description": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

}