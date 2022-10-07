package test

import (
	"testing"

	"github.com/S3B4SZ17/Web_Algo/algorithms"
)

func TestAddTwoNumbers(t *testing.T) {
	// First list def
	list := &algorithms.List{}
	list.Add(3) // 3
	list.Add(8) // 2
	list.Add(5) // 1

	// Second list def
	list2 := &algorithms.List{}
	list2.Add(5)
	list2.Add(4)
	list2.Add(2)

	// Summ the 2 lists together
	// 583 + 245 = 828
	algorithms.SumLists(list, list2)

	got := algorithms.SumLists(list, list2).PrintList()
	res := &algorithms.List{}
	res.Add(8)
	res.Add(2)
	res.Add(8)
	want := res.PrintList()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
