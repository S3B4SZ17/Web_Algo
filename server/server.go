package server

import (
	"net"

	pbAddTwoNum "github.com/S3B4SZ17/Web_Algo/proto/addTwoNumbers"
	pbReverseLinkedList "github.com/S3B4SZ17/Web_Algo/proto/reverseLinkedlist"
	pbReverse "github.com/S3B4SZ17/Web_Algo/proto/reverseNumber"
	"google.golang.org/grpc"
)

func RegisterServices(s *grpc.Server) {
	pbReverse.RegisterReverseNumberServer(s, &ReverseNumberServer{})
	pbAddTwoNum.RegisterAddTwoNumbersServer(s, &AddTwoNumbersServer{})
	pbReverseLinkedList.RegisterReverseLinkedListServer(s, &ReverseLinkedListServer{})

}

func StartServer(s *grpc.Server, l net.Listener) error {
	return s.Serve(l)
}
