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

	klines, err := client.NewKlinesService().Limit(3).Symbol("BTCUSDT").
		Interval("1d").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, k := range klines {
		println(k.Open)
	}

	// for _, kline := range klines {
	// 	klineArrayRaw := fmt.Sprintf("%+v\n", kline)
	// 	formattedKlineArray := strings.ReplaceAll(
	// 		strings.ReplaceAll(
	// 			strings.ReplaceAll(klineArrayRaw, "\n", ","),
	// 			"&", ""),
	// 		" ", ",")
	// 	reformattedklineArray := strings.ReplaceAll(
	// 		strings.ReplaceAll(
	// 			strings.ReplaceAll(
	// 				strings.ReplaceAll(formattedKlineArray, ",", `", "`),
	// 				":", `":"`),
	// 			"}", `"}`),
	// 		"{", `{"`)
	// 	klineArray := reformattedklineArray[:len(reformattedklineArray)-3]
	// 	// klineArrayJson, err := json.Marshal(klineArray)
	// 	// if err != nil {
	// 	// 	fmt.Println(err)
	// 	// 	return
	// 	// }
	// 	// var klineArrayJsonString Kline
	// 	// unmarshaledKlineArray := json.Unmarshal([]byte(klineArrayJson), &klineArrayJsonString)

	// 	fmt.Println(klineArray)

	// }
}

// Kline define kline info
type Kline struct {
	OpenTime                 string `json:"OpenTime"`
	Open                     string `json:"open"`
	High                     string `json:"High"`
	Low                      string `json:"Low"`
	Close                    string `json:"Close"`
	Volume                   string `json:"Volume"`
	CloseTime                string `json:"CloseTime"`
	QuoteAssetVolume         string `json:"QuoteAssetVolume"`
	TradeNum                 string `json:"TradeNum"`
	TakerBuyBaseAssetVolume  string `json:"TakerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"TakerBuyQuoteAssetVolume"`
}
