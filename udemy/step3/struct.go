package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func main() {

	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 10000
	fmt.Println(v)

	v2 := Vertex{X: 9898}
	fmt.Println(v2)

	v3 := Vertex{222, 423, "hello"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4)

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)

	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)
}
