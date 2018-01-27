package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 1
	low := 0
	high := len(slice)
	for {
		fmt.Println("slice:", slice)
		fmt.Println("high:", high)
		fmt.Println("low", low)
		if low >= high {
			fmt.Println("result: -1")
			break
		}
		middle := (low + high) / 2
		fmt.Println("middle:", middle)
		if slice[middle] == target {
			fmt.Println("result:", middle)
			break
		}
		if slice[middle] >= target {
			high = middle - 1
			continue
		}
		if slice[middle] <= target {
			low = middle + 1
			continue
		}

	}
}
