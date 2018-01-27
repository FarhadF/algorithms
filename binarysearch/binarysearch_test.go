package main

import "testing"

func TestBinarySearch(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarySearch(slice, 8)
	if result != 8 {
		t.Error("expecting 8 but got: ", result)
	}
	result = binarySearch(slice, 2)
	if result != 2 {
		t.Error("expecting 2 but got: ", result)
	}
	result = binarySearch(slice, 212)
	if result != -1 {
		t.Error("expecting -1 but got: ", result)
	}
}
