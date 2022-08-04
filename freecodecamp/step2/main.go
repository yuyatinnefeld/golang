package main

import (
	"fmt"
)

const constantVal = iota

func main() {
	n := 1 == 1
	m := 2 == 5
	fmt.Printf("%v, %T \n", n, n)
	fmt.Printf("%v, %T \n", m, m)

	var a uint16 = 42
	fmt.Printf("%v, %T \n", a, a)

	x := 10 // 1010
	y := 3  // 0011

	fmt.Println(x & y)  // 0010 = 2
	fmt.Println(x | y)  // 1011 = 11
	fmt.Println(x ^ y)  // 1001 = 9
	fmt.Println(x &^ y) // 1000 = 8

	b := 16
	fmt.Println(b << 4) // 2^4 * 2^4 = 2^8 = 256
	fmt.Println(b >> 3) // 2^4 / 2^3 = 2^1 = 2

	const myConst int = 42
	fmt.Printf("%v, %T \n", myConst, myConst)
}
