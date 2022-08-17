package main

import "fmt"

type MyInt int

func (i MyInt) Double() int {
	fmt.Printf("%T, %v \n", i, i)
	return int(i * 2)
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
