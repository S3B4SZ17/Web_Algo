package server

import (
	"context"

	mgt "github.com/S3B4SZ17/Web_Algo/management"
	pb "github.com/S3B4SZ17/Web_Algo/proto/reverseLinkedlist"
	"github.com/S3B4SZ17/Web_Algo/services"
)

type ReverseLinkedListServer struct {
	pb.UnimplementedReverseLinkedListServer
}

func (s *ReverseLinkedListServer) ReverseLinkedList(ctx context.Context, in *pb.LinkedList) (res *pb.ReverseList, err error) {
	// mgt.Info.Printf("Received: Number = %v ", in.Number)

	res, err = services.ReverseLinkedList_service(in)
	if err != nil {
		mgt.Log.Error("[Error] An error occurred while making the reverseLinkedList: " + err.Error())
		return
	}

	return
}
