package main

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

func main() {
	var (
		apiKey    = ""
		secretKey = ""
	)
	client := binance.NewClient(apiKey, secretKey)
	//futuresClient := binance.NewFuturesClient(apiKey, secretKey)   // USDT-M Futures
	//deliveryClient := binance.NewDeliveryClient(apiKey, secretKey) // Coin-M Futures

	klines, err := client.NewKlinesService().Limit(5).Symbol("BTCUSDT").
		Interval("1d").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, kline := range klines {
		fmt.Print(kline)
	}
}
