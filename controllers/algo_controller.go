package controllers

import (
	"net/http"

	"github.com/S3B4SZ17/Web_Algo/algorithms"
	"github.com/gin-gonic/gin"
)

func TwoNums(c *gin.Context){
	list1 := &algorithms.List{}
	list1.Add(2)
	list1.Add(4)
	list1.Add(3)

	list2 := &algorithms.List{}
	list2.Add(5)
	list2.Add(6)
	list2.Add(4)
	list2.Add(5)

	res1 := list1.PrintList()
	res2 := list2.PrintList()

	sumList := algorithms.SumLists(list1, list2)
	sum := sumList.PrintList()

	c.JSON(http.StatusOK, gin.H{
		"list1": res1,
		"list2": res2,
		"sum": sum,
	})

}