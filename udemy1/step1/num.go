package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)

	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("%v, %T \n", u8, u8)
	/*
		long comment out
		blbalblalbal
	*/
	fmt.Println("1+1= ", 1+1)
	fmt.Println("10%3= ", 10%3)
	fmt.Println("10/2= ", 10/2)

	//gofmt num.go
	//gofmt -w num.go

	var s string = "Hello World"
	fmt.Println(strings.Replace(s, "H", "WW", 1))
	fmt.Println(strings.Contains(s, "Hello"))

}
