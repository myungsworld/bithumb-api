package main

import (
	"fmt"
	"myungsworld/api/bithumb/Info"
	Execute "myungsworld/api/bithumb/transaction"
	"myungsworld/database"
	"strconv"
	"sync"
	"time"
)

var wait = sync.WaitGroup{}


func main() {

	database.ConnectDB()

	loop := make(chan bool, 1)

	var sellTransaction int
	var buyTransaction int
	var seconds int
	var startPriceEveryTenMin float64
	var marketPrice float64
	var highPrice float64

	go func() {

		for true {

			if buyTransaction == 1 {

				wait.Add(1)
				fmt.Println("과연연")
				time.Sleep(time.Second*10)
				fmt.Println("역시는역시")
				wait.Done()


			}

			BTTMarketPrice := Info.CoinMarketCondition("BTT")
			startPriceEveryTenMin, _ = strconv.ParseFloat(BTTMarketPrice, 64)

			for true {
				time.Sleep(time.Second * 1)

				// 10분 주기로 초기화
				seconds++
				if seconds == 600 {
					sellTransaction = 0
					buyTransaction = 0
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
				fmt.Println("변동률:", fluctateRate, " ", seconds, "초")

				// 10분안에 -3% 이상 떨어지면 50퍼 매도
				if fluctateRate < -3 && sellTransaction == 0 {
					balance := Info.GetMyTickerBalance("BTT")
					balance = balance / 2
					Execute.MarketSell("BTT", balance)
					sellTransaction++
				}

				// 10분안에 3% 이상 오를시 ~원 매수 + 변동률 보면서 대기
				if fluctateRate > 0.01 && buyTransaction == 0 {
					availableKRW := Info.GetBalance("ALL")
					fmt.Println("availableKRW : ", availableKRW)
					buyTransaction++
					break
				}

			}
		}
	}()

	for i := 0; i < 1; i++ {
		<-loop
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
