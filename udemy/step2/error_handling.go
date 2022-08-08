package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err_1 := os.Open("./if.go")
	if err_1 != nil {
		log.Fatal("Error")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err_2 := file.Read(data)
	if err_2 != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	if err_2 = os.Chdir("test"); err_2 != nil {
		log.Fatalln("Error by os.Chdir")
	} else {
		fmt.Println("ok")
	}

}
