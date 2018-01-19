package main

import "testing"

func TestFibonacciRecursiveMemoization(t *testing.T) {
	i := fibonacciRecursiveMemo(4)
	if i != 3 {
		t.Error("expected 4 but got:", i)
	}
	j := fibonacciRecursiveMemo(0)
	if j != 0 {
		t.Error("expected 0 but got:", j)
	}
}