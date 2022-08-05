package main

import "fmt"

func main() {
	m1 := map[string]int{"apple": 100, "banana": 200}
	m2 := map[int]string{1: "hallo", 2: "world"}

	fmt.Println(m1)
	fmt.Println(m1["apple"])
	m1["banana"] = 10000
	fmt.Println(m1)

	fmt.Println(m2)
	v, ok := m1["apple"]
	v1, ok1 := m1["nothing"]
	fmt.Println(v, ok)
	fmt.Println(v1, ok1)
}
