package main

import (
	"fmt"
	"log"
	"mymodule/bitflyer"
	"mymodule/mypackage"
	"mymodule/utils"
)

func main() {
	utils.LoggingSettings(mypackage.Config.LogFile)
	log.Println("test")
	//fmt.Println(mypackage.Config.ApiKey)
	//fmt.Println(mypackage.Config.ApiSecret)

	apiClient := bitflyer.New(mypackage.Config.ApiKey, mypackage.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())

	tickerChannel := make(chan bitflyer.Ticker)
	go apiClient.GetRealTimeTicker(mypackage.Config.ProductCode, tickerChannel)
	for ticker := range tickerChannel {
		fmt.Println(ticker)
	}

	//fmt.Println(models.DbConnection)

}
