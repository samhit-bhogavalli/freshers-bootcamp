package main

import (
	"strconv"
	"testing"
)

type isEvenTests struct {
	input    int
	required bool
}

/*func TestIsEven(t *testing.T) {
	ans := IsEven(4)
	if ans != true {
		t.Errorf("test failed required %t but given %t", ans, !ans)
	}
}*/

func TestIsEvenTableDriven(t *testing.T) {
	Tests := []isEvenTests{{5, false}, {4, true}}
	for i, test := range Tests {
		t.Run("test "+strconv.Itoa(i), func(t *testing.T) {
			if test.required != IsEven(test.input) {
				t.Error("test failed")
			}
		})
	}
}
