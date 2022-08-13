package server

import (
	"context"

	mgt "github.com/S3B4SZ17/Web_Algo/management"
	pb "github.com/S3B4SZ17/Web_Algo/proto/addTwoNumbers"
	"github.com/S3B4SZ17/Web_Algo/services"
)
type AddTwoNumbersServer struct {
	pb.UnimplementedAddTwoNumbersServer
}

func (s *AddTwoNumbersServer) AddTwoNumbers (ctx context.Context, in *pb.ListReq) (res *pb.ListSum, err error) {
	mgt.Info.Printf("Received: List 1 %v and List 2 %v", in.ListVal1, in.ListVal2)
	
	res, err = services.GetTwoSumsResult_Service(in); if err != nil {
		mgt.Error.Printf("[Error] An error occurred while making the sum: %v", err)
		return
	}
	
	return 
}
