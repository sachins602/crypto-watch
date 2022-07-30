package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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

	klines, err := client.NewKlinesService().Limit(10).Symbol("BTCUSDT").
		Interval("1d").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, kline := range klines {
		klineArrayRaw := fmt.Sprintf("%+v\n", kline)
		formattedKlineArray := strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(klineArrayRaw, "\n", ","),
				"&", ""),
			" ", ", ")
		klineArray := removeLastRune(formattedKlineArray)
		klineArrayJson, err := json.Marshal(klineArray)
		if err != nil {
			fmt.Println(err)
			return
		}
		var klineArrayJsonString Kline
		unmarshaledKlineArray := json.Unmarshal([]byte(klineArrayJson), &klineArrayJsonString)

		fmt.Println(unmarshaledKlineArray)

	}
}

// Kline define kline info
type Kline struct {
	OpenTime                 int64  `json:"openTime"`
	Open                     string `json:"open"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	Close                    string `json:"close"`
	Volume                   string `json:"volume"`
	CloseTime                int64  `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	TradeNum                 int64  `json:"tradeNum"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
}

func removeLastRune(s string) string {
	r := []rune(s)
	return string(r[:len(r)-1])
}
