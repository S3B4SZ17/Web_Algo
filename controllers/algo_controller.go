package controllers

import (
	"fmt"
	"net/http"

	"github.com/S3B4SZ17/Web_Algo/algorithms"
	"github.com/S3B4SZ17/Web_Algo/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func TwoSums(c *gin.Context){
	// item1 := c.PostForm("list1")
	// item2 := c.PostFormMap("list2")
	
	var list1vals *[]algorithms.ListVals
	// var list2vals algorithms.ListVals

	//using BindJson method to serialize body with struct
	if err := c.ShouldBindBodyWith(&list1vals, binding.JSON);err!=nil{
   		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(415, gin.H{"errcode": 415, "description": "Post Data Err"})
   		return
	}

	// //using BindJson method to serialize body with struct
	// if err := c.ShouldBindBodyWith(&list2vals, binding.JSON);err!=nil{
   	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	c.JSON(415, gin.H{"errcode": 415, "description": "Post Data Err"})
   	// 	return
	// }

	fmt.Printf("list1: %v\n", list1vals)
	// fmt.Printf("list2: %v\n", list2vals)

	res := services.GetTwoSumsResult_Service(list1vals)

	c.JSON(http.StatusOK, res)

}