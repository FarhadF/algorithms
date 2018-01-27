package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 9
	result := binarySearch(slice, target)
	fmt.Println("result:", result)
}

func binarySearch(slice []int, target int) int {
	l := len(slice)
	//fmt.Println("slice:", slice)
	//fmt.Println("length", l)
	//fmt.Println("length/2:", l/2)

	if l == 0 {
		return -1
	} else {
		middle := l / 2
		//fmt.Println("middle element:", slice[middle])
		if slice[middle] == target {
			return target
		} else if slice[middle] <= target {
			return binarySearch(slice[middle+1:], target)
		} else {
			return binarySearch(slice[:middle], target)
		}
	}
}
