package controllers

import (
	"net/http"

	"github.com/S3B4SZ17/Web_Algo/algorithms"
	"github.com/S3B4SZ17/Web_Algo/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func TwoSums(c *gin.Context){
	var list1vals *[]algorithms.ListVals

	//using BindJson method to serialize body with struct
	if err := c.ShouldBindBodyWith(&list1vals, binding.JSON);err!=nil{
   		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(415, gin.H{"errcode": 415, "description": "Post Data Err"})
   		return
	}

	res, err := services.GetTwoSumsResult_Service(list1vals); if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(415, gin.H{"errcode": 415, "description": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}