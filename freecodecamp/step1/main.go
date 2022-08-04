package main

import (
	"fmt"
	"strconv"
)

var I int = 20 //upper case out of func
var theURLRequest = "http://google.com"
var myName string = "hello"

var (
	userID     int     = 123
	userName   string  = "Tom"
	userAge    int     = 23
	userWeight float32 = 3.0
)

func main() {
	j := 50 //lower case in the func
	fmt.Println(j)
	// comment
	fmt.Printf("%v, %T \n", I, I)
	fmt.Printf("%v, %T \n", j, j)

	fmt.Println(myName)
	fmt.Println(userName)
	fmt.Println(userID)
	fmt.Println(userAge)

	var x int = 42
	var k float32
	k = float32(x) //convert int to float32
	fmt.Printf("%v, %T \n", k, k)
	var kk = strconv.Itoa(x)
	fmt.Printf("%v, %T \n", kk, kk)
}
