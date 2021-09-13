package Soaring

import (
	"fmt"
	"myungsworld/api/bithumb/Info"
	Execute "myungsworld/api/bithumb/transaction"
	"myungsworld/database"
	"myungsworld/database/models"
	"strconv"
	"sync"
	"time"
)

func SellingTiming(ticker string, startPrice float64, each float64) {
	var wait = sync.WaitGroup{}
	wait.Add(1)

	seconds := 0
	firstSell := false
	SecondSell := false
	LastSell := false

	for true {

		seconds++

		if seconds == 1200 {
			break
		}

		tickerMarketPrice := Info.CoinMarketCondition(ticker)
		marketPrice, _ := strconv.ParseFloat(tickerMarketPrice, 64)
		fluctateRate := ((marketPrice - startPrice) / marketPrice) * 100

		fmt.Println("폭등감지 두번째 대기열 진입 새로 카운트 시작")
		fmt.Println("시작가 :", startPrice, "현재가 :", marketPrice)
		fmt.Println("변동률:", fluctateRate, " ", seconds, "초")

		if fluctateRate > 10 && firstSell == false {
			amount := each / 5
			status, message, EA := Execute.MarketSell(ticker, amount)

			if status == "0000" {

				totalKRW := amount * marketPrice

				transaction := models.Transaction{
					Type:        "폭등감지10%매도",
					Ticker:      ticker,
					Amount:      EA,
					TotalKRW:    totalKRW,
					StartPrice:  startPrice,
					MarketPrice: marketPrice,
					Fluctate:    fluctateRate,
					Seconds:     seconds,
				}
				if err := database.DB.Create(&transaction).Error; err != nil {
					panic(err)
				}

				fmt.Println("10퍼 매도 성공")

			} else {
				info := models.Information{
					Ticker:    ticker,
					Status:    status,
					Message:   message,
					Content:   "폭등감지 10퍼 매도 실패",
					CreatedAt: time.Now(),
				}

				if err := database.DB.Create(&info).Error; err != nil {
					panic(err)
				}

				fmt.Println("10퍼 매도 실패")
			}

			firstSell = true

		}

		if fluctateRate > 20 && SecondSell == false {
			amount := each / 5
			amount *= 2
			status, message, EA := Execute.MarketSell(ticker, amount)

			if status == "0000" {

				totalKRW := amount * marketPrice

				transaction := models.Transaction{
					Type:        "폭등감지20%매도",
					Ticker:      ticker,
					Amount:      EA,
					TotalKRW:    totalKRW,
					StartPrice:  startPrice,
					MarketPrice: marketPrice,
					Fluctate:    fluctateRate,
					Seconds:     seconds,
				}
				if err := database.DB.Create(&transaction).Error; err != nil {
					panic(err)
				}

				fmt.Println("두번째 대기열 20퍼 매도 성공")
			} else {
				info := models.Information{
					Ticker:    ticker,
					Status:    status,
					Message:   message,
					Content:   "폭등감지 20퍼 매도 실패",
					CreatedAt: time.Now(),
				}

				if err := database.DB.Create(&info).Error; err != nil {
					panic(err)
				}

				fmt.Println("20퍼 매도 실패")

			}

			SecondSell = true

		}

		if fluctateRate > 30 && LastSell == false {
			balance := Info.GetMyTickerBalance(ticker)
			status, message, EA := Execute.MarketSell(ticker, balance)

			if status == "0000" {
				totalKRW := balance * marketPrice

				transaction := models.Transaction{
					Type:        "폭등감지30%매도",
					Ticker:      ticker,
					Amount:      EA,
					TotalKRW:    totalKRW,
					StartPrice:  startPrice,
					MarketPrice: marketPrice,
					Fluctate:    fluctateRate,
					Seconds:     seconds,
				}
				if err := database.DB.Create(&transaction).Error; err != nil {
					panic(err)
				}
			} else {
				info := models.Information{
					Ticker:    ticker,
					Status:    status,
					Message:   message,
					Content:   "폭등감지 30퍼 매도 실패",
					CreatedAt: time.Now(),
				}

				if err := database.DB.Create(&info).Error; err != nil {
					panic(err)
				}

				fmt.Println("30퍼 매도 실패")

			}

			LastSell = true
		}

		// 30퍼까지 다 팔았으면 대기열 해제
		if firstSell == true && SecondSell == true && LastSell == true {
			break
		}

		time.Sleep(time.Second * 1)

	}

	wait.Done()

}
