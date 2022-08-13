package server

import (
	"context"

	mgt "github.com/S3B4SZ17/Web_Algo/management"
	pb "github.com/S3B4SZ17/Web_Algo/proto/reverseNumber"
	"github.com/S3B4SZ17/Web_Algo/services"
)
type ReverseNumberServer struct {
	pb.UnimplementedReverseNumberServer
}

func (s *ReverseNumberServer) ReverseNumber (ctx context.Context, in *pb.Number) (res *pb.Reverse, err error) {
	mgt.Info.Printf("Received: Number = %v ", in.Number)
	
	res, err = services.Reverse_service(in); if err != nil {
		mgt.Error.Printf("[Error] An error occurred while making the sum: %v", err)
		return
	}
	
	return 
}