package main

import "fmt"

func main() {
	slice := []int{1, 300, 12, 16, 230, 0, 4, 30, 80}
	length := len(slice)
	fmt.Println(selectionSort(slice, length))
}
func selectionSort(slice []int, length int) []int {
	var min int
	var temp int
	for i := 0; i < length-1; i++ {
		min = i
		for j := i + 1; j < length; j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}
		temp = slice[i]
		slice[i] = slice[min]
		slice[min] = temp
	}
	return slice
}
