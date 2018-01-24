package main

import "fmt"

func main() {
	target := "test"
	slice := []string{"wtf", "is", "this", "shitty", "test"}
	fmt.Println(linearSearch(slice, target))
}
//O(n)
func linearSearch(slice []string, target string) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
