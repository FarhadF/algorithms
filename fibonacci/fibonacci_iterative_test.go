package main

import "testing"

func TestFibonacciIterative(t *testing.T) {
	i := fibonacciIterative(4)
	if i != 3 {
		t.Error("expected 4 but got:", i)
	}
	j := fibonacciIterative(0)
	if j != 0 {
		t.Error("expected 0 but got:", j)
	}
}