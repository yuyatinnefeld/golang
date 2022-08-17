package main

import "fmt"

func main() {
	f := 11

	ff := float64(f)

	fmt.Printf("%v, %T \n", ff, ff)

	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])

	m1 := map[string]int{"mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T, %v \n", m1, m1)
}
