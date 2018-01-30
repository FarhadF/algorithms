package main

import "testing"

func TestBubbleSort(t *testing.T) {
	slice := []int{1, 300, 12, 14, 5, 0, 230}
	l := len(slice)
	result := bubbleSort(slice, l)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l-1; j++ {
			if result[i] > result[j] {
				t.Error("expected ", result[i], " to be smaller than ", result[j])
				break
			}
		}
	}
}
