package main

import (
	"fmt"
	"log"

	"mymodule/mypackage"
	"mymodule/utils"
)

func main() {
	utils.LoggingSettings(mypackage.Config.LogFile)
	log.Println("test")
	fmt.Println(mypackage.Config.ApiKey)
	fmt.Println(mypackage.Config.ApiSecret)

}
