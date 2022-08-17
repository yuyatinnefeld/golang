package main

import "fmt"

func do(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	case bool:
		fmt.Println(v)
	default:
		fmt.Println("I dont know the type")
	}

}

func main() {
	var i interface{} = 10
	do(i)
	do("hallo")
	do(true)
}
