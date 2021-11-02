package cataclysm

import (
	"fmt"
	"myungsworld/api/bithumb/Info"
	"myungsworld/database"
	"myungsworld/database/models"
	"myungsworld/queue/cataclysm/Crashing"
	"myungsworld/queue/cataclysm/Soaring"
	"strconv"
	"time"
)

func Start(ticker string,cycle int , percentCrashing float64) {

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

			// cycle초 주기로 초기화
			seconds++
			if seconds == cycle {
				seconds = 0
				startPriceEveryTenMin = 0
				marketPrice = 0

				break
			}

			tickerMarketPrice := Info.CoinMarketCondition(ticker)
			marketPrice, _ = strconv.ParseFloat(tickerMarketPrice, 64)

			fluctateRate := ((marketPrice - startPriceEveryTenMin) / marketPrice) * 100

			fmt.Println(ticker, "변동률", fluctateRate)

			// 폭락 방지 함수
			// cycle 초 안에 percentCrashing% 이상 떨어지면 50퍼 매도 + 두번째 대기열로 전환 + 메세지
			// 대기열 진입후 10분 더 지켜보다가 -5퍼 이상 떨어지면 남은 코인의 절반 더 매도
			if fluctateRate < percentCrashing {

				//정보 수집
				info := models.Information{
					Ticker: ticker,
					Content: fmt.Sprintf(
						"%d초 -3퍼센트 하락 (시작가 : %.8f 현재가 : %.8f)",
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
				Crashing.BreakForCrashed(ticker, startPriceEveryTenMin, marketPrice, seconds, fluctateRate)
				break

			}
			// 폭등 감지 함수
			// 10분안에 3% 이상 오를시 50000원 매수 + 두번째 대기열로 전환 + 메세지
			if fluctateRate > 3 {

				//정보 수집
				info := models.Information{
					Ticker: ticker,
					Content: fmt.Sprintf(
						"%d초 3퍼센트 상승 (시작가 : %.8f 현재가 : %.8f)",
						seconds,
						startPriceEveryTenMin,
						marketPrice,
					),
					CreatedAt: time.Now(),
				}
				if err := database.DB.Create(&info).Error; err != nil {
					panic(err)
				}

				// 매수할 돈이 부족할시
				availableKRW := Info.GetBalance(ticker)
				if availableKRW < 50000 {
					break
				}

				//매수시작(코인 티커, 한화 수량, 시장가, 시작가, 변동률, 걸린시간)
				Soaring.BreakForSoared(ticker, 50000, marketPrice, startPriceEveryTenMin, fluctateRate, seconds)
				break
			}

		}
	}

}
