package main

import (
	"fmt"
	"time"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	polkadotStock := "DOT-USD"
	currentTime := time.Now()
	today := currentTime.Format("2006-01-02")
	polkadot, _ := quote.NewQuoteFromYahoo(polkadotStock, "2022-08-01", today, quote.Daily, true)

	fmt.Println("### CSV DATA ###")
	fmt.Print(polkadot.CSV())

	fmt.Println("### RSI ###")
	rsi2 := talib.Rsi(polkadot.Close, 2)
	fmt.Println(rsi2)

	fmt.Println("### MOVING AVG ###")
	mva := talib.Ema(polkadot.Close, 14)
	fmt.Println(mva)

	//https://finance.yahoo.com/chart/DOT-USD
}
