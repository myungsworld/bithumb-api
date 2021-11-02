package Soaring

import (
	"fmt"
	Execute "myungsworld/api/bithumb/transaction"
	"myungsworld/database"
	"myungsworld/database/models"
	"sync"
	"time"
)

func BreakForSoared(ticker string, krw, marketPrice, startPrice, fluctate float64, seconds int, percentFirstSell, percentSecondSell, percentLastSell float64, soaringCycle int) {
	var wait = sync.WaitGroup{}
	wait.Add(1)

	each := krw / marketPrice
	EA := fmt.Sprintf("%.4f", each)

	status, message := Execute.MarKetBuy(ticker, EA)

	if status == "0000" {

		transaction := models.Transaction{
			Type:        "폭등감지매수",
			Ticker:      ticker,
			Amount:      EA,
			TotalKRW:    krw,
			StartPrice:  startPrice,
			MarketPrice: marketPrice,
			Fluctate:    fluctate,
			Seconds:     seconds,
			Content:     "매수후 두번째 대기열 진입",
			CreatedAt:   time.Now(),
		}

		if err := database.DB.Create(&transaction).Error; err != nil {
			panic(err)
		}

		// 매수 이후 두번째 대기열
		SellingTiming(ticker, marketPrice, each, percentFirstSell, percentSecondSell, percentLastSell, soaringCycle)

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
