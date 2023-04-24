package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]int)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	fmt.Println(map1)

	var keys []int
	for key := range map1 {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Printf("map1[%v]=%v\n", key, map1[key])
	}
}
