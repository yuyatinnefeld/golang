package main

import (
	"fmt"

	"myProject/mylib"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 1000}
	fmt.Println(mylib.Average(s))

}
