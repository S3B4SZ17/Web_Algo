package test

import (
	"testing"

	"github.com/S3B4SZ17/Web_Algo/algorithms"
)

func TestLongestSubstring(t *testing.T) {

	got := algorithms.LengthOfLongestSubstring("anviaj")
	want := 5

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	} else {
		t.Logf("\n\n============= Test passed ============= \ngot %d\nwanted %d\n\n", got, want)
	}
}
