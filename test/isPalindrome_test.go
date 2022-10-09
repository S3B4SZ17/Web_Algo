package test

import (
	"testing"

	"github.com/S3B4SZ17/Web_Algo/algorithms"
)

func TestIsPalindrome(t *testing.T) {

	got := algorithms.IsPalindrom("Oir a dario")
	// Oir a dario
	// Arriba la birra"
	// Roma ni se conoce sin oro ni se conoce sin amor
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	} else {
		t.Logf("\n\n============= Test passed ============= \ngot %t\nwanted %t\n\n", got, want)
	}
}

func TestPalindromePairs(t *testing.T) {
	palidrome_array := []string{"abcd", "dcba", "lls", "s", "sssll"}
	got := algorithms.PalindromePairs(palidrome_array)

	want := [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}}
	if len(got) != len(want) {
		t.Errorf("got %d, wanted %d", got, want)
	}
	for i := 0; i < len(got); i++ {
		for j := 0; j < len(got[i]); j++ {
			if got[i][j] != want[i][j] {
				t.Errorf("got[%d][%d] = %d; want[%d][%d] = %d; \n", i, j, got[i][j], i, j, want[i][j])
			}
		}
	}

}
