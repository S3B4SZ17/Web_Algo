package services

import (
	"github.com/S3B4SZ17/Web_Algo/algorithms"
)

func GetTwoSumsResult_Service(list1vals *[]algorithms.ListVals) *algorithms.Response {
	list1 := &algorithms.List{}
	list2 := &algorithms.List{}

	list1.AddFromList(&(*list1vals)[0])
	list2.AddFromList(&(*list1vals)[1])

	res1 := list1.PrintList()
	res2 := list2.PrintList()

	sumList := algorithms.SumLists(list1, list2)
	sum := sumList.PrintList()

	response := &algorithms.Response{ List1: res1, List2: res2, Sum: sum}

	return response
}