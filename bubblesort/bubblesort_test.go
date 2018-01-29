package main

import "testing"

func TestBubbleSort(t *testing.T) {
	slice := []int{1, 300, 12, 14, 5, 0, 230}
	l := len(slice)
	result := bubbleSort(slice, l)
	for i := 0; i < l-1; i++ {
		if result[i] > result[i+1] {
			t.Error("expected ", result[i], " to be smaller than ", result[i+1])
			break
		}
	}
}
