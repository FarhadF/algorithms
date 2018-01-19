package main

import (
	"os"
	"strconv"
	"fmt"
)

func main() {
	input := os.Args[1]
	i, _ := strconv.Atoi(input)
	fmt.Println(fibonacciRecursiveMemo(i))
}

var fibMap = make(map[int]int)
//using memoization
func fibonacciRecursiveMemo(f int) int {
	if f == 0 {
		return 0
	} else if f == 1 {
		return 1
	} else if fibMap[f] != 0 {
		return fibMap[f]
	} else {
		fibMap[f] = fibonacciRecursiveMemo(f-1) + fibonacciRecursiveMemo(f-2)
		return fibMap[f]
	}
}
