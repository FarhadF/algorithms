package main

import "testing"

func TestFibonacciRecursive(t *testing.T) {
	i := fibonacciRecursive(4)
	if i != 3 {
		t.Error("expected 4 but got:", i)
	}
	j := fibonacciRecursive(0)
	if j != 0 {
		t.Error("expected 0 but got:", j)
	}
}
