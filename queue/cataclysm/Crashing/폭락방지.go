package Crashing

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

func BreakForCrashed(ticker string, startPriceEveryTenMin float64, marketPrice float64, seconds int, fluctateRate float64, cycle int, percentSecondCrashing float64) {

	fmt.Println(ticker, "Break For Crashed 진입")

	var wait = sync.WaitGroup{}

	// 대기
	wait.Add(1)

	balance := Info.GetMyTickerBalance(ticker)
	balance = balance / 2
	status, message, EA := Execute.MarketSell(ticker, balance)

	if status == "0000" {
		// database stuff
		transaction := models.Transaction{
			Type:        "폭락감지매도",
			Ticker:      ticker,
			Amount:      EA,
			TotalKRW:    balance * marketPrice,
			StartPrice:  startPriceEveryTenMin,
			MarketPrice: marketPrice,
			Fluctate:    fluctateRate,
			Seconds:     seconds,
			Content:     "1차 매도후 두번째 대기열 진입",

			CreatedAt: time.Now(),
		}

		if err := database.DB.Create(&transaction).Error; err != nil {
			panic(err)
		}

		// SNS stuff

		// 10분동안 더 떨어지면 더 팜
		newSeconds := 0
		for true {
			time.Sleep(time.Second * 1)
			newSeconds++
			if newSeconds >= cycle {
				break
			}

			tickerMarketPrice := Info.CoinMarketCondition(ticker)
			marketPrice, _ = strconv.ParseFloat(tickerMarketPrice, 64)
			fluctateRate2 := ((marketPrice - startPriceEveryTenMin) / marketPrice) * 100

			// 총 20분 모니터링 동안 -5퍼가 떨어지면 남은것의 절반을 더 매도
			if fluctateRate2 < percentSecondCrashing {
				balance2 := Info.GetMyTickerBalance(ticker)
				if balance2 <= 0 {
					break
				} else {
					balance2 = balance2 / 2
					status2, message2, EA2 := Execute.MarketSell(ticker, balance)

					if status2 == "0000" {
						// database stuff
						transaction2 := models.Transaction{
							Type:        "폭락감지2차매도",
							Ticker:      ticker,
							Amount:      EA2,
							TotalKRW:    balance * marketPrice,
							StartPrice:  startPriceEveryTenMin,
							MarketPrice: marketPrice,
							Fluctate:    fluctateRate2,
							Seconds:     newSeconds,
							Content:     fmt.Sprintf("%f퍼", percentSecondCrashing),
							CreatedAt:   time.Now(),
						}

						if err := database.DB.Create(&transaction2).Error; err != nil {
							panic(err)
						}

						break
					} else {
						info := models.Information{
							Ticker:    ticker,
							Status:    status2,
							Message:   message2,
							CreatedAt: time.Now(),
						}

						if err := database.DB.Create(&info).Error; err != nil {
							panic(err)
						}
						break
					}

				}

			}

		}
	} else {
		info := models.Information{
			Ticker:    ticker,
			Status:    status,
			Message:   message,
			CreatedAt: time.Now(),
		}

		if err := database.DB.Create(&info).Error; err != nil {
			panic(err)
		}

	}

	wait.Done()
}
