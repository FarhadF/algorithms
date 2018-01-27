package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	i, _ := strconv.Atoi(input)
	fmt.Println(fibonacciRecursive(i))
}

//Time complexity - O(2^n)
//Space complexity - O(n)
func fibonacciRecursive(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return fibonacciRecursive(i-1) + fibonacciRecursive(i-2)
	}
}
