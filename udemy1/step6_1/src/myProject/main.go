package main

import (
	"fmt"

	"myProject/mylib"
	"myProject/mylib/under"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 1000}
	fmt.Println(mylib.Average(s))
	fmt.Println(mylib.SaySomething())

	under.SubHello()

	p1 := mylib.Person{Name: "Tom", Age: 30}
	fmt.Println(p1)

	fmt.Println(mylib.Public)
}
