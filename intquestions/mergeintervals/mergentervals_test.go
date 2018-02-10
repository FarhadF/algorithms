package main

import "testing"

func TestMergeIntervals(t *testing.T) {
	var input = []interval{
		interval{1, 5},
		interval{5, 8},
		interval{7, 9},
		interval{15, 20},
		interval{21, 23},
		interval{22, 25},
	}
	merged := merge(input)
	if merged[0].start != 1 || merged[0].end != 9 || merged[1].start != 15 || merged[1].end != 20 || merged[2].start != 21 || merged[2].end != 25 {
		t.Error("expected [{1 9} {15 20} {21 25}] but got ", merged)

	}
}
