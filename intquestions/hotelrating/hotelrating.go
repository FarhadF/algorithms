package main

import (
	"fmt"
	"sort"
)

type hotel struct {
	id     int
	rating int
}

func main() {
	var n int
	if m, err := fmt.Scan(&n); err != nil || m != 1 {
		fmt.Println(m)
		panic(err)
	}
	fmt.Println(n)
	hotelsMap := make(map[int]int)
	for i := 0; i < n; i++ {
		input := make([]int, 2)
		for j := 0; j < 2; j++ {

			if z, err := fmt.Scan(&input[j]); err != nil || z != 1 {
				fmt.Println(z)
				panic(err)
			}
		}
		if v, ok := hotelsMap[input[0]]; ok {
			hotelsMap[input[0]] = (input[1] + v) / 2
		} else {
			hotelsMap[input[0]] = input[1]
		}
		fmt.Println(input)
	}
	fmt.Println(hotelsMap)
	hotels := make([]hotel, len(hotelsMap))
	i := 0
	for k, v := range hotelsMap {
		hotels[i] = hotel{k, v}
		i++
	}
	sort.Slice(hotels, func(i, j int) bool {
		if hotels[i].rating == hotels[j].rating {
			if hotels[i].id > hotels[i].id {
				return false
			} else {
				return true
			}
		} else {
			return hotels[i].rating > hotels[j].rating
		}
	})
	fmt.Println(hotels)
}
