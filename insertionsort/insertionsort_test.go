package main

import "testing"

func TestSelectionSort(t *testing.T) {
	slice := []int{1, 300, 12, 16, 230, 0, 4, 30, 80}
	result := insertionSort(slice, len(slice))
	l := len(slice)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l-1; j++ {
			if result[i] > result[j] {
				t.Error("expected ", result[i], "to be less than", result[j])
			}
		}
	}
}
