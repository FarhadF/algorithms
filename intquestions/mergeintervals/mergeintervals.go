package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type interval struct {
	start int
	end   int
}

//first input will be an integer with number of intervals
//next inputs will be intervals consisting of 2 integers seperated by space

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	intervals := getIntervals(n, scanner)
	//sort on start ascending
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].start < intervals[j].start {
			return true
		} else {
			return false
		}
	})
	fmt.Println(intervals)
	merged := merge(intervals)
	fmt.Println(merged)
}

//get intervals
func getIntervals(n int, scanner *bufio.Scanner) []interval {
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
	return intervals
}

//merge intervals
func merge(intervals []interval) []interval {
	var merged []interval
	for i, v := range intervals {
		if i == 0 {
			merged = append(merged, v)
		} else if i != len(intervals) {
			if v.start <= merged[len(merged)-1].end {
				merged[len(merged)-1].end = v.end
			} else {
				merged = append(merged, v)
			}
		}
	}
	return merged
}
