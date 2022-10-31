package controllers

import (
	"net/http"

	"github.com/S3B4SZ17/Web_Algo/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Solve_algorithm(c *gin.Context) {
	var file *services.AlgoFile

	//using BindJson method to serialize body with struct
	if err := c.ShouldBindWith(&file, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(415, gin.H{"errcode": 415, "description": "Bad data sent. Only use whole numbers."})
		return
	}

	valid, err := services.CompareResult(file.File)
	var res *services.AlgoResponse
	if err != nil {
		res = &services.AlgoResponse{Valid: valid, Message: err.Error()}
	} else {
		res = &services.AlgoResponse{Valid: valid, Message: "Resolviste el algoritmo!!"}
	}

	c.JSON(http.StatusOK, res)
}
