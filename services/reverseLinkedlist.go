package services

import (
	"github.com/S3B4SZ17/Web_Algo/algorithms"
	mgt "github.com/S3B4SZ17/Web_Algo/management"
	pbRevList "github.com/S3B4SZ17/Web_Algo/proto/reverseLinkedlist"
)

func ReverseLinkedList_service(linkedList *pbRevList.LinkedList) (res *pbRevList.ReverseList, err error) {
	res = &pbRevList.ReverseList{}
	list := &algorithms.List{}
	mgt.Log.Info(linkedList.String())
	list.CallReverseLinkedList(linkedList)
	res.LinkedList = list.PrintIntList()

	return res, err
}
