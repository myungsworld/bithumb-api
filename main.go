package main

import (
	"fmt"
	"myungsworld/api/bithumb/Info"
	"myungsworld/database"
	"strconv"
	"time"
)

func main() {

	database.ConnectDB()

	done := make(chan bool, 100)

	var transaction int
	var seconds int
	var startPriceEveryTenMin float64
	var marketPrice float64
	var highPrice float64

	go func() {

		for true {

			BTTMarketPrice := Info.CoinMarketCondition("BTT")
			startPriceEveryTenMin, _ = strconv.ParseFloat(BTTMarketPrice, 64)

			for true {
				time.Sleep(time.Second * 1)

				// 30분 주기로 초기화
				seconds++
				if seconds == 20 {
					transaction = 0
					seconds = 0
					startPriceEveryTenMin = 0
					marketPrice = 0
					highPrice = 0
					break
				}

				fmt.Println("BTT 매수타임 호시탐탐 검색중")

				BTTMarketPrice := Info.CoinMarketCondition("BTT")
				marketPrice, _ = strconv.ParseFloat(BTTMarketPrice, 64)
				fmt.Println("marketPrice: ", marketPrice)

				if marketPrice > highPrice {
					highPrice = marketPrice
				}

				fluctateRate := ((marketPrice - startPriceEveryTenMin) / marketPrice) * 100

				fmt.Println("시작가 :", startPriceEveryTenMin, "현재가 :", marketPrice, "고가 :", highPrice)
				fmt.Println("변동률:", fluctateRate)
				fmt.Println(seconds,"초")

				// 30분안에 -7% 이상 떨어지면 50퍼 매도 + 시작가 변경
				if transaction == 0 && fluctateRate < -7 {
					transaction++
					done <- true
				}

				if BTTMarketPrice < "4" {
					done <- true

				}
			}
		}
	}()

	for i := 0; i < 3; i++ {
		<-done
		fmt.Println("아항")
	}

	//Info.GetBalance("ALL")
	//Info.GetMyTickerBalance("BTT","KRW")

	//Info.GetBalance("BTT")
	//Info.FetchAddress("BTT")
	//Execute.WithDrawKRW("302-0709-1079-11","1000")
	//Info.PendingOrder("BTT")
	//Execute.BuyHook("BTT","5000","5")
	//Execute.SellHook("BTT","500","6")
	//Execute.MarKetBuy("BTT","2")
	//Execute.MarketSell("BTT","1")
	//Info.CandleStick("BTT","24h")
	//Execute.MarketBuyHook("BTT","4","4","1500")
}
