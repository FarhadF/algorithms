package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"strings"
	"sort"
)

type interval struct {
	start int
	end   int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	//fmt.Println(n)
	intervals := make([]interval, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		interval := interval{}
		split := strings.Split(scanner.Text(), " ")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		interval.start, interval.end = start, end
		intervals = append(intervals, interval)
	}
	//fmt.Println(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].start < intervals[j].start {
			return true
		} else {
			return false
		}
	})
	fmt.Println(intervals)
	//fmt.Println(len(intervals))
	var merged []interval
	//var index int = 0
	for i, v := range intervals {
		//fmt.Println("i",i)
		if i== 0 {
			merged = append(merged,v)
		} else if i != len(intervals) {
    		if v.start <= merged[len(merged)-1].end {
				//fmt.Println("possible merge")
				merged[len(merged)-1].end = v.end
				//			index ++
			} else {
				merged = append(merged, v)
			}
		}
	}
	fmt.Println(merged)
}
