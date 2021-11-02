package Queue

import (
	"myungsworld/queue/cataclysm"
)

func Coin(ticker string, cycle int, percentCrashing float64) {

	// cycle 간격으로 폭등감지 or 폭락감지 후 매수 or 매도
	go cataclysm.Start(ticker,cycle, percentCrashing)

}
