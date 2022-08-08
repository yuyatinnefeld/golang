package main

import "fmt"

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	fmt.Println(n)

	fmt.Println(n)
	fmt.Println(&n)

	var p *int = &n // pointing
	fmt.Println(p)
	fmt.Println(*p)

	var m int = 10000
	one(&m)
	fmt.Println(m)

}
