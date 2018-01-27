package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	i, _ := strconv.Atoi(input)
	fmt.Println(fibonacciIterative(i))
}

//Time complexity - O(n)
//Space complexity - O(1)
func fibonacciIterative(f int) int {
	if f == 0 {
		return 0
	} else {
		previousFib := 0
		fib := 1
		for i := 2; i <= f; i++ {
			tmp := fib
			fib += previousFib
			previousFib = tmp
		}
		return fib
	}
}
