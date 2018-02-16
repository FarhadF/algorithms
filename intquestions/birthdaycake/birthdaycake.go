package main

import (
	"bufio"
	"os"
	"strconv"
	//"strings"
	"fmt"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	/*maxCapacity := 10000000000
	buf := make([]byte, 8000)
	scanner.Buffer(buf, maxCapacity)*/
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	//fmt.Println(n)
	if err != nil {
		panic(err)
	}
	strSlice := make([]string, 0)
    //bufio cant scan more than 4096 bytes from input so try to call split function on space (scanner.Split(bufio.ScanWords))
    // and loop append
	for i := 0; i < n; i++ {
		scanner.Scan()
		strSlice = append(strSlice, scanner.Text())
	}
	if err := scanner.Err(); err != nil {

		fmt.Fprintln(os.Stderr, "reading input:", err)

	}
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		//fmt.Println(i)
		slice[i], err = strconv.Atoi(strSlice[i])
		if err != nil {
			panic(err)
		}
		//fmt.Println(slice)
	}
	sort.Slice(slice, func(i, j int) bool {
		if slice[i] > slice[j] {
			return true
		} else {
			return false
		}
	})

	//fmt.Println(slice)
	var c int
	for i := 0; i < n-1; i++ {
		c++
		if slice[i] > slice[i+1] {
			break
		}
		if i == n-2 {
			c++
		}
	}
	fmt.Println(c)

}
