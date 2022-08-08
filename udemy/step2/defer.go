package main

import (
	"fmt"
	"os"
)

func foo() {
	fmt.Println("later foo()")
}

func boo() {
	defer fmt.Println("later boo()")
	fmt.Println("boo()")
}

func main() {
	boo()
	defer foo()
	fmt.Println("first")

	file, _ := os.Open("./if.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Print(string(data))

}
