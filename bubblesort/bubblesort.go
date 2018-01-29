package main

import "fmt"

func main() {
	//initial array to sort
	slice := []int{7, 4, 5, 2, 10, 0, 12, 15, 3}
	n := len(slice)
	fmt.Println(bubbleSort(slice, n))
}

//O(n^2)
func bubbleSort(slice []int, length int) []int {
	var temp int
	// after each iteration through whole slice the max value should be the last element, so we dont need to check it
	// again, the first loop ensures this.
	for j := 0; j < length-1; j++ {
		for i := 0; i < length-j-1; i++ {
			if slice[i] > slice[i+1] {
				temp = slice[i]
				slice[i] = slice[i+1]
				slice[i+1] = temp
			}
		}
	}
	return slice
}
