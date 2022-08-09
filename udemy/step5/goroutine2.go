package main

import (
	"fmt"
	"sync"
	"time"
)

func gogoroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

	wg.Done()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go gogoroutine("world", &wg)
	normal("hello")
	wg.Wait()
}
