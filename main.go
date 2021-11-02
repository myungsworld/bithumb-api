package main

import (
	"myungsworld/database"
	Queue "myungsworld/queue"
	"time"
)

// Ticker : 코인티커 BasicCycle : 기본 대기 시간
// PercentCrashing : 50퍼 매도 시점 마이너스율 PercentSecondCrashing : 매도 후 그 반절 매도 시점 CrashedCycle : 폭락감지 큐 대기 시간
// PercentSoaring : 50000원 매수 시점 플러스율 PercentFirstSell : 첫번째 매도 플러스율 PercentSecondSell : 두번째, PercentLastSell : 마지막 SoaringCycle : 폭등 감지 큐 대기 시간
type Config struct {
	Ticker     string
	BasicCycle int

	PercentCrashing       float64
	PercentSecondCrashing float64
	CrashedCycle          int

	PercentSoaring    float64
	PercentFirstSell  float64
	PercentSecondSell float64
	PercentLastSell   float64
	SoaringCycle      int
}

func main() {

	database.ConnectDB()
	loop := make(chan bool, 1)

	list := make([]Config, 0)
	list = setCoinDefaultConfig(list)

	for _, config := range list {
		time.Sleep(time.Millisecond*250)
		go Queue.Coin(
			config.Ticker, config.BasicCycle,
			config.PercentCrashing, config.PercentSecondCrashing, config.CrashedCycle,
			config.PercentSoaring, config.PercentFirstSell, config.PercentSecondSell, config.PercentLastSell, config.SoaringCycle,
		)
	}

	for i := 0; i < 1; i++ {
		<-loop
	}

}

func setCoinDefaultConfig(list []Config) []Config {


	// order by opening price desc
	list = append(list,
		Config{
			Ticker: "BNB", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		}, Config{
			Ticker: "KSM", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "COMP", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "AAVE", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "SOL", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "LTC", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "BSV", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "AXS", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "BTG", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "ETC", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "DOT", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "NMR", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "LUNA", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "ATOM", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "LINK", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "LPT", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "BAL", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "WAVES", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "XVS", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "UNI", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "REP", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "CAKE", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "QTUM", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "ALICE", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "OMG", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},Config{
			Ticker: "SUSHI", BasicCycle: 600,
			PercentCrashing: -3, PercentSecondCrashing: -5, CrashedCycle: 600,
			PercentSoaring: 3, PercentFirstSell: 10, PercentSecondSell: 20, PercentLastSell: 30, SoaringCycle: 1200,
		},
		)

	return list
}