package main

import (
	"fmt"
	"os/user"
	"time"
)

const Pi = 3.14
const (
	Username = "myName"
	PassWord = "myPassWord"
)

func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	bazz()
	fmt.Println("konichiwa world")
	fmt.Println(time.Now())
	fmt.Println(user.Current())

	fmt.Println(Pi, Username, PassWord)

}
