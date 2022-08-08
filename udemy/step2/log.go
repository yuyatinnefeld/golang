package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)

}

func main() {
	LoggingSettings("test.log")
	fmt.Println("hi")
	log.Println("logging!")
	log.Printf("%v, %v", "hi", "test")
	//log.Fatalln("error")
	//fmt.Println("does not execute")

	_, error := os.Open("not-exists-file")
	if error != nil {
		log.Fatalln("Exit", error)
	}

}
