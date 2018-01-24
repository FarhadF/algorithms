package main

import "testing"

func TestLinearSearch(t *testing.T) {

	target := "target"
	slice := []string{"this", "slice", "contains", "target", "not?"}
	result := linearSearch(slice, target)
	if result != 3 {
		t.Error("expected 3 but got ", result)
	}
}
