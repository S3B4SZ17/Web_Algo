package controllers

import (
	"net/http"

	mgt "github.com/S3B4SZ17/Web_Algo/management"
	pbReverse "github.com/S3B4SZ17/Web_Algo/proto/reverseNumber"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func Reverse(c *gin.Context){
	var number *pbReverse.Number

	// Set up a connection to the AddTwoNumbers server.
	conn, err := grpc.Dial(host + ":" +gRPCListener, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		mgt.Error.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pbReverse.NewReverseNumberClient(conn)

	//using BindJson method to serialize body with struct
	if err := c.ShouldBindWith(&number, binding.JSON);err!=nil{
   		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(415, gin.H{"errcode": 415, "description": "Bad data sent. Only use whole numbers."})
   		return
	}

	res, err := client.ReverseNumber(c, number); if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(415, gin.H{"errcode": 415, "description": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}