package controllers

import (
	"net/http"

	mgt "github.com/S3B4SZ17/Web_Algo/management"
	pbRevList "github.com/S3B4SZ17/Web_Algo/proto/reverseLinkedlist"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ReverseLinkedList(c *gin.Context) {
	var linkedList *pbRevList.LinkedList

	// Set up a connection to the AddTwoNumbers server.
	conn, err := grpc.Dial(host+":"+gRPCListener, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		mgt.Log.Fatal("Did not connect: " + err.Error())
	}
	defer conn.Close()
	client := pbRevList.NewReverseLinkedListClient(conn)

	//using BindJson method to serialize body with struct
	if err := c.ShouldBindWith(&linkedList, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(415, gin.H{"errcode": 415, "description": "Bad data sent. Only use whole numbers."})
		return
	}

	res, err := client.ReverseLinkedList(c, linkedList)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(415, gin.H{"errcode": 415, "description": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}
