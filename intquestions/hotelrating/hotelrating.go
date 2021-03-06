package main

import (
	"fmt"
	"sort"
	"time"
)

type hotel struct {
	id     int
	rating int
}

/*sample input number of rows then each row hotel id and rating with space delimiter if  hotels with the same rating we need to print the one with larger hotel id:
	4
	1000 8
	2000 8
	2000 10
	1000 9
sample output:
	2000
	1000
*/

func main() {
	start := time.Now()
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
			if hotels[i].id < hotels[j].id {
				return false
			} else {
				return true
			}
		} else {
			return hotels[i].rating > hotels[j].rating
		}
	})
	fmt.Println(hotels)
	for _, h := range hotels {
		fmt.Println(h.id)

	}
	fmt.Println(time.Since(start))
}
