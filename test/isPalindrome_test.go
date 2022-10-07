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

func TestGetWords(t *testing.T) {
	palidrome_array := []string{"abcd", "dcba", "lls", "s", "sssll"}
	got, _ := algorithms.GetWords(palidrome_array)

	want := [][2]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}}
	if len(*got) != len(want) {
		t.Errorf("got %d, wanted %d", got, want)
	}
	for i, v := range *got {
		if v != want[i] {
			t.Errorf("got %d, wanted %d", got, want)
		}
	}
	t.Logf("\n\n============= Test passed ============= \ngot %q\nwanted %q\n\n", got, want)

}
