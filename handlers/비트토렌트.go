package Handler

import (
	"fmt"
	"myungsworld/api/bithumb/Info"
	Execute "myungsworld/api/bithumb/transaction"
	"strconv"
	"sync"
	"time"
)
func BTT() {

	var wait = sync.WaitGroup{}
	var sellTransaction int
	var buyTransaction int
	var seconds int
	var startPriceEveryTenMin float64
	var marketPrice float64
	var highPrice float64

	for true {

		if buyTransaction == 1 {

			wait.Add(1)

			// transaction stuff


			// database stuff


			fmt.Println("이것첫번째 10초 짜리")
			time.Sleep(time.Second*10)
			fmt.Println("역시는역시")
			wait.Done()
			sellTransaction = 0
			buyTransaction = 0
			seconds = 0

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

			// 10분안에 -3% 이상 떨어지면 50퍼 매도 + 대기열로 전환 + 메세지
			if fluctateRate < -3 && sellTransaction == 0 {
				balance := Info.GetMyTickerBalance("BTT")
				balance = balance / 2
				Execute.MarketSell("BTT", balance)
				sellTransaction++

				// queue stuff

				// SNS stuff

				// database stuff


			}

			// 10분안에 3% 이상 오를시 ~원 매수 + 대기열로 전환 + 메세지
			if fluctateRate > 0.01 && buyTransaction == 0 {
				availableKRW := Info.GetBalance("ALL")
				fmt.Println("availableKRW : ", availableKRW)
				buyTransaction++

				// SNS stuff

				// database stuff
				break
			}

		}
	}
}

