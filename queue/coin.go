package Queue

import (
	"fmt"
	"myungsworld/api/bithumb/Info"
	"myungsworld/database"
	"myungsworld/database/models"
	"strconv"
	"time"
)

func Coin(ticker string) {

	var seconds int
	var startPriceEveryTenMin float64
	var marketPrice float64

	for true {
		// 초기화
		seconds = 0
		startPriceEveryTenMin = 0
		marketPrice = 0

		tickerMarketPrice := Info.CoinMarketCondition(ticker)
		startPriceEveryTenMin, _ = strconv.ParseFloat(tickerMarketPrice, 64)

		for true {
			time.Sleep(time.Second * 1)

			// 10분 주기로 초기화
			seconds++
			if seconds == 600 {
				seconds = 0
				startPriceEveryTenMin = 0
				marketPrice = 0

				break
			}

			fmt.Println(ticker, "매수타임 호시탐탐 검색중")

			tickerMarketPrice := Info.CoinMarketCondition(ticker)
			marketPrice, _ = strconv.ParseFloat(tickerMarketPrice, 64)
			fmt.Println("marketPrice: ", marketPrice)

			fluctateRate := ((marketPrice - startPriceEveryTenMin) / marketPrice) * 100

			fmt.Println("시작가 :", startPriceEveryTenMin, "현재가 :", marketPrice)
			fmt.Println("변동률:", fluctateRate, " ", seconds, "초")

			// 폭락 방지 함수
			// 10분안에 -3% 이상 떨어지면 50퍼 매도 + 대기열로 전환 + 메세지
			// 대기열 진입후 10분 더 지켜보다가 -5퍼 이상 떨어지면 남은 코인의 절반 더 매도
			if fluctateRate < -3 {

				//정보 수집
				info := models.Information{
					Ticker: ticker,
					Content: fmt.Sprintf(
						"%d초 -3퍼센트 하락 (시작가 : %.6f 현재가 : %.6f)",
						seconds,
						startPriceEveryTenMin,
						marketPrice,
					),
					CreatedAt: time.Now(),
				}
				if err := database.DB.Create(&info).Error; err != nil {
					panic(err)
				}

				// 매도할 수량이 없으면 break
				balance := Info.GetMyTickerBalance(ticker)
				if balance*marketPrice <= 5000 {
					break
				}

				// 매도 시작
				BreakForCrashed(ticker, startPriceEveryTenMin, marketPrice, seconds, fluctateRate)
				break

			}
			// 폭등 감지 함수
			// 10분안에 3% 이상 오를시 ~원 매수 + 대기열로 전환 + 메세지
			if fluctateRate > 3 {
				availableKRW := Info.GetBalance("ALL")
				fmt.Println("availableKRW : ", availableKRW)

				//// 잔고가 만원 이하일시
				//if availableKRW
			}

		}
	}
}
