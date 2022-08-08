package controllers

import (
	"net/http"

	"github.com/S3B4SZ17/Web_Algo/algorithms"
	"github.com/S3B4SZ17/Web_Algo/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Reverse(c *gin.Context){
	var reverse algorithms.Reverse

	//using BindJson method to serialize body with struct
	if err := c.ShouldBindBodyWith(&reverse, binding.JSON);err!=nil{
   		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(415, gin.H{"errcode": 415, "description": "Bad data sent. Only use whole numbers."})
   		return
	}

	res, err := services.Reverse_service(reverse.Reverse); if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(415, gin.H{"errcode": 415, "description": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}