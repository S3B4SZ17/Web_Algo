package algorithms

func CallReverse(number int32 ) int32{
	
	return reverse(number, 0)
}

func reverse(number int32, result int32) int32 {
	var d int32
	var r int32
	
	if (number / 10 == 0 && number % 10 == 0) {
		return result
	}else{
		d = number % 10
		r = number / 10

		result = result * 10 + d
		return reverse(r, result)
	}
}

type Reverse struct {
	Reverse int32 `json:"reverse"`
}

type ReverseResponse struct {
	Reverse_result int32 `json:"reverse_result"`
}