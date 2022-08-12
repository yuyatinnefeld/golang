package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 100
	fmt.Println(len(ch))
	x := <-ch
	fmt.Println(x)
	ch <- 100
	fmt.Println(len(ch))

	close(ch)
	for c := range ch {
		fmt.Println(c)
	}
}
