package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"time"
	"strings"
	"sort"
)

type hotel struct {
	id int
	rating int
}

func main() {
	start := time.Now()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	hotelsMap := make(map[int]int)
	for i:=0; i<n; i++ {
		scanner.Scan()
		hotelRating := strings.Split(scanner.Text()," ")
		var hotelRatingInt []int
		for _, v := range hotelRating {
			temp, _ := strconv.Atoi(v)
			hotelRatingInt = append(hotelRatingInt, temp)
		}
		if v, ok := hotelsMap[hotelRatingInt[0]]; ok {
			hotelsMap[hotelRatingInt[0]] = (hotelsMap[hotelRatingInt[1]] + v) / 2
		} else {
			hotelsMap[hotelRatingInt[0]] = hotelRatingInt[1]
		}
	}
	i := 0
	hotelsRatingSlice := make([]hotel, len(hotelsMap))
	fmt.Println(hotelsMap)
	for k,v := range hotelsMap{
		hotelsRatingSlice[i] = hotel{k, v}
		i++
	}
	sort.Slice(hotelsRatingSlice, func(i,j int) bool {
		if hotelsRatingSlice[i].rating == hotelsRatingSlice[j].rating {
			if hotelsRatingSlice[i].id < hotelsRatingSlice[j].id {
				return false
			} else {
				return true
			}
		} else {
			return hotelsRatingSlice[i].rating > hotelsRatingSlice[j].rating
		}
	})
	fmt.Println(hotelsRatingSlice)
	for _, h := range hotelsRatingSlice {
		fmt.Println(h.id)
	}
	fmt.Println(time.Since(start))
}
