package services

import (
	"github.com/S3B4SZ17/Web_Algo/algorithms"
	pb "github.com/S3B4SZ17/Web_Algo/proto/addTwoNumbers"
)

func GetTwoSumsResult_Service(list1vals *pb.ListReq) (response *pb.ListSum, err error) {
	list1 := &algorithms.List{}
	list2 := &algorithms.List{}

	err = list1.AddFromList(list1vals.ListVal1); if err != nil {
		return nil,err
	}
	err = list2.AddFromList(list1vals.ListVal2); if err != nil {
		return nil,err
	}

	res1 := list1.PrintList()
	res2 := list2.PrintList()

	sumList := algorithms.SumLists(list1, list2)
	sum := sumList.PrintList()

	response = &pb.ListSum{ List1: res1, List2: res2, Sum: sum}

	return response, err
}