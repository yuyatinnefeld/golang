package main

import (
	"fmt"
	"strconv"
)

func main() {

	t, f := true, false
	fmt.Printf("%T, %v \n", t, t)
	fmt.Printf("%T, %v \n", f, f)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var x int = 1
	xx := float64(x)
	fmt.Printf("%T, %v, %f \n", xx, xx, xx)

	var s string = "14"
	i, _ := strconv.Atoi(s) // str to int
	fmt.Printf("%T, %v \n", i, i)

	var ii int = 34
	st := strconv.Itoa(ii) // int to str
	fmt.Printf("%T, %v \n", st, st)

	// array for fixed value list
	var a [2]int
	a[0] = 100
	a[1] = 300

	fmt.Println(a)

	var b [2]int = [2]int{200, 200}
	fmt.Println(b)

	// slice for the flexible value list
	var sb []int = []int{100, 300}
	sb = append(sb, 301)
	sb = append(sb, 302)
	fmt.Println(sb)

	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)

	fmt.Println(n[2:4])
	fmt.Println(n[len(n)-1])

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)

	//make and cap of slice
	valSlice := make([]int, 3, 5)
	fmt.Printf("len=%d, cap=%d, value=%v \n", len(valSlice), cap(valSlice), valSlice)
	valSlice = append(valSlice, 1, 2, 34)
	fmt.Printf("len=%d, cap=%d, value=%v \n", len(valSlice), cap(valSlice), valSlice)

	//co := make([]int, 5)
	co := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		co = append(co, i)
		fmt.Println(co)
	}

	fmt.Println(co)
}
