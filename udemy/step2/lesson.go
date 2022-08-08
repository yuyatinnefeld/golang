package main

import (
	"fmt"
	"sort"
)

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	fmt.Println(l)

	sort.Ints(l)
	fmt.Println(l)

	a := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	fmt.Println(a)

	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	total := 0
	for _, v := range m {
		total += v
	}
	fmt.Println(total)

}
