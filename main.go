package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	apiData, err := getApiData()
	if err != nil {
		log.Fatal("Error getting api data", err)
	}
	fmt.Println(apiData)

	coinHistoryData, err := getCoinHistoryData()
	if err != nil {
		log.Fatal("Error getting coin history data", err)
	}
	fmt.Println(coinHistoryData)

}

func getApiData() ([]CryptoModel, error) {
	apiLink := os.Getenv("API_LINK")
	apiResp, err := http.Get(apiLink)
	if err != nil {
		return nil, err
	}

	defer apiResp.Body.Close()

	body, err := ioutil.ReadAll(apiResp.Body)
	if err != nil {
		return nil, err
	}

	var cryptoDetails []CryptoModel

	err = json.Unmarshal([]byte(body), &cryptoDetails)
	if err != nil {
		return nil, err
	}

	return cryptoDetails, nil
}

func getCoinHistoryData() (CoinHistoryModel, error) {
	coinBaseLink := os.Getenv("COIN_BASE_LINK")
	coinEndTimeLink := os.Getenv("COIN_ENDTIME")
	coinLimitLink := os.Getenv("COIN_LIMIT")
	coinSymbolLink := os.Getenv("COIN_SYMBOL")
	coinIntervalLink := os.Getenv("COIN_INTERVAL") //1m 3m 5m 15m 30m 1h 2h 4h 8h 12h 1d 1w 1M

	coinEndTime := fmt.Sprintf(strconv.FormatInt(time.Now().Round(time.Minute).Unix(), 10) + "000")
	coinLimit := "1000"
	coinSymbol := "BTCUSDT"
	coinInterval := "1m"

	coinHistoryLink := fmt.Sprintf(
		coinBaseLink +
			coinEndTimeLink + coinEndTime +
			coinLimitLink + coinLimit +
			coinSymbolLink + coinSymbol +
			coinIntervalLink + coinInterval)

	fmt.Println(coinHistoryLink)

	coinHistoryResp, err := http.Get(coinHistoryLink)
	if err != nil {
		return nil, err
	}

	defer coinHistoryResp.Body.Close()

	body, err := ioutil.ReadAll(coinHistoryResp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(body)

	var coinHistoryDetails CoinHistoryModel

	err = json.Unmarshal([]byte(body), &coinHistoryDetails)
	if err != nil {
		return nil, err
	}

	return coinHistoryDetails, nil
}

type CryptoModel struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int    `json:"firstId"`
	LastID             int    `json:"lastId"`
	Count              int    `json:"count"`
}

// [
//   [
//     1499040000000,      // Open time
//     "0.01634790",       // Open
//     "0.80000000",       // High
//     "0.01575800",       // Low
//     "0.01577100",       // Close
//     "148976.11427815",  // Volume
//     1499644799999,      // Close time
//     "2434.19055334",    // Quote asset volume
//     308,                // Number of trades
//     "1756.87402397",    // Taker buy base asset volume
//     "28.46694368",      // Taker buy quote asset volume
//     "17928899.62484339" // Ignore.
//   ]
// ]

// [
//     [
//         1659046380000,
//         "23862.89000000",
//         "23893.03000000",
//         "23862.89000000",
//         "23873.90000000",
//         "146.18315000",
//         1659046439999,
//         "3491227.90198050",
//         4314,
//         "80.55943000",
//         "1924030.29215680",
//         "0"
//     ]
// ]

type CoinHistoryModel [][]struct {
	OpenTime       int64
	Open           string
	High           string
	Low            string
	Close          string
	Volume         string
	CloseTime      int64
	QuoteVolume    string
	NumberOfTrades int
	TakerBuyBase   string
	TakerBuyQuote  string
	Ignore         string
}

// 	 [
//     [
//         1659046380000,
//         "23862.89000000",
//         "23893.03000000",
//         "23862.89000000",
//         "23873.90000000",
//         "146.18315000",
//         1659046439999,
//         "3491227.90198050",
//         4314,
//         "80.55943000",
//         "1924030.29215680",
//         "0"
//     ],
//     [
//         1659046380000,
//         "23862.89000000",
//         "23893.03000000",
//         "23862.89000000",
//         "23873.90000000",
//         "146.18315000",
//         1659046439999,
//         "3491227.90198050",
//         4314,
//         "80.55943000",
//         "1924030.29215680",
//         "0"
//     ],
// ]
