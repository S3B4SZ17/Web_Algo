package services

import (
	"github.com/S3B4SZ17/Web_Algo/algorithms"
	pbReverse "github.com/S3B4SZ17/Web_Algo/proto/reverseNumber"
)

func Reverse_service(number *pbReverse.Number) (res *pbReverse.Reverse, err error) {
	res = &pbReverse.Reverse{}
	response := algorithms.CallReverse(number.Number)
	res.Reverse = response 

	return res, err
}