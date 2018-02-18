package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

//get N and space delimited array and print the reverse

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	//fmt.Println(n)
	scanner.Scan()
	strSlice := strings.Split(scanner.Text(), " ")
	reverseSlice := make([]string, 0)
	for i := n - 1; i >= 0; i-- {
		//fmt.Println(strSlice[i])
		reverseSlice = append(reverseSlice, strSlice[i])
	}
	//fmt.Println(reverseSlice, strSlice)
	fmt.Println(strings.Join(reverseSlice, " "))
}
