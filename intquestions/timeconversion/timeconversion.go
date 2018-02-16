package main

import (
	"bufio"
	"os"
	"time"
	"fmt"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	t, err := time.Parse("3:04:05PM", str)
	if err != nil {
		panic(err)
	}
	//fmt.Println(t)
	fmt.Println(t.Format("15:04:05"))
}
