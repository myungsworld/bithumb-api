package Queue

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

func BreakForCrashed(ticker string, startPriceEveryTenMin float64, marketPrice float64) {

	var wait = sync.WaitGroup{}

	// 대기
	wait.Add(1)

	balance := Info.GetMyTickerBalance(ticker)
	if balance == 0 {
		fmt.Println(ticker, "잔고가 없음")
		time.Sleep(time.Second * 5)
	} else {
		balance = balance / 2
		EA := Execute.MarketSell(ticker, balance)

		// database stuff
		transaction := models.Transaction{
			Type:      "매도",
			Ticker:    ticker,
			Amount:    EA,
			TotalKRW:  balance * marketPrice,
			CreatedAt: time.Now(),
		}

		if err := database.DB.Create(&transaction).Error; err != nil {
			panic(err)
		}

		// SNS stuff

		// 10분동안 더 떨어지면 더 팜

		for true {
			time.Sleep(time.Second * 1)

			tickerMarketPrice := Info.CoinMarketCondition(ticker)
			marketPrice, _ = strconv.ParseFloat(tickerMarketPrice, 64)
			fluctateRate := ((marketPrice - startPriceEveryTenMin) / marketPrice) * 100

			// 총 20분 모니터링 동안 -5퍼가 떨어지면 남은것의 절반을 더 매도
			if fluctateRate < -5 {
				balance2 := Info.GetMyTickerBalance(ticker)
				if balance2 == 0 {
					fmt.Println(ticker, "잔고가 없음")
					time.Sleep(time.Second * 5)
				} else {
					balance2 = balance2 / 2
					EA2 := Execute.MarketSell(ticker, balance)

					// database stuff
					transaction2 := models.Transaction{
						Type:      "매도",
						Ticker:    ticker,
						Amount:    EA2,
						TotalKRW:  balance * marketPrice,
						CreatedAt: time.Now(),
					}

					if err := database.DB.Create(&transaction2).Error; err != nil {
						panic(err)
					}
				}

				break

			}

		}

		wait.Done()

	}
}