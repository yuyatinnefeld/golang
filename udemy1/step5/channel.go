package main

import "fmt"

func gogoroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func gogoroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go gogoroutine1(s, c)
	go gogoroutine2(s, c)
	result := <-c
	fmt.Println(result)
	result2 := <-c
	fmt.Println(result2)

}
