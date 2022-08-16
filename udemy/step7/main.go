package main

import (
	"fmt"

	"github.com/antonholmquist/jason"
)

func main() {
	v, _ := jason.NewObjectFromBytes([]byte(`{"Name": "fetaro"}`))
	fmt.Println(v)
}
