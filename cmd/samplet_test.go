package main

import "testing"

func TestAdd(t *testing.T) {
	var x, y, result = 2, 2, 4

	realResult := Multiple(x, y)

	if realResult != result {
		t.Errorf("Multiple err: %d != %d", result, realResult)
	}
}
