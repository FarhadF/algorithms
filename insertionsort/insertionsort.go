package main

import "fmt"

func main() {
	slice := []int{7, 4, 5, 2, 10, 0, 12, 15, 3}
	n := len(slice)
	fmt.Println(insertionSort(slice,n))
}

func insertionSort(slice []int, n int) []int{
	for i:=1; i<n ; i++ {
		for j := i; j>0 && slice[j] < slice[j - 1]; j--{
			temp := slice[j - 1]
			slice[j - 1] = slice[j]
			slice[j] = temp


		}

	}
	return slice
}
