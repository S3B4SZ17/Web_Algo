package services

import (
	"github.com/S3B4SZ17/Web_Algo/algorithms"
)

func Reverse_service(number int32) (res int32, err error) {

	res = algorithms.CallReverse(number)

	return res, err
}