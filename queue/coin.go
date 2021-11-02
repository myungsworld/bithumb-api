package Queue

import (
	"myungsworld/queue/cataclysm"
)

func Coin(ticker string, basicCycle int,
	percentCrashing, percentSecondCrashing float64, crashedCycle int,
	percentSoaring, percentFirstSell, percentSecondSell, percentLastSell float64, soaringCycle int) {

	// cycle 간격으로 폭등감지 or 폭락감지 후 매수 or 매도
	go cataclysm.Start(
		ticker, basicCycle,
		percentCrashing, percentSecondCrashing, crashedCycle,
		percentSoaring, percentFirstSell, percentSecondSell, percentLastSell, soaringCycle)
}
