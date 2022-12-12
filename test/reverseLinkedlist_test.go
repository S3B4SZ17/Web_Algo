package test

import (
	"testing"

	"github.com/S3B4SZ17/Web_Algo/algorithms"
)

func TestReverse(t *testing.T) {
	list := &algorithms.List{}
	list.Add_seq(5)
	list.Add_seq(8)
	list.Add_seq(3)

	list.Reverse()

	got := list.PrintList()
	want := "5, 8, 3, 0"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Logf("\n\n============= Test passed ============= \ngot %q\nwanted %q\n\n", got, want)
	}
}
