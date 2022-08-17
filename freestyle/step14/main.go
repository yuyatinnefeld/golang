package main

import (
	"fmt"
)

func Sum(x int, y int) int {
    return x + y
}

func main() {
	var a, b int = 5, 10
    var total int = Sum(a, b)
	var name = "Sum(x, y)"
	fmt.Printf("%s is executed\n", name)
	fmt.Printf("%d + %d = %d\n", a, b, total)
}