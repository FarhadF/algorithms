package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 3
	fmt.Println(binarySearchIterative(slice, target))
}


func binarySearchIterative(slice []int, target int) int {
	low := 0
	high := len(slice)
	for {
		fmt.Println("slice:", slice)
		fmt.Println("high:", high)
		fmt.Println("low", low)
		middle := (low + high) / 2
		if low > high || middle >= len(slice) {
			fmt.Println("result: -1")
			return -1
		}

		fmt.Println("middle:", middle)
		if slice[middle] == target {
			fmt.Println("result:", middle)
			return middle
		}
		if slice[middle] > target {
			high = middle - 1
			continue
		}
		if slice[middle] < target {
			low = middle + 1
			continue
		}
	}
}