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

	bi := []byte{72, 73}
	fmt.Println(bi)
	fmt.Println(string(bi))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

	add, mul := add(10, 20)
	fmt.Println(add, mul)

	fu := func(x int) {
		fmt.Println("inner func", x)
	}

	fu(233)

	foo(10, 30)
	foo(10, 20, 30)
}

func add(x int, y int) (add int, mul int) {
	add = x + y
	mul = x * y
	return add, mul
}

func foo(params ...int) {
	fmt.Println(len(params), params)

	for _, param := range params {
		fmt.Println(param)
	}
}
