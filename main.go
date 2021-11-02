package main

import (
	"myungsworld/database"
	Queue "myungsworld/queue"
)

// Ticker : 코인티커
// Cycle : 회복주기
// PercentCrashing : 50퍼 매도 시점 마이너스율
// PercentLastCrashing : 매도 후 그 반절 매도 시점
type Config struct {
	Ticker                string
	Cycle                 int
	PercentCrashing       float64
	PercentSecondCrashing float64
}

func main() {

	database.ConnectDB()

	loop := make(chan bool, 1)

	list := make([]Config, 0)
	list = append(list,
		Config{
			Ticker: "BTT", Cycle: 600, PercentCrashing: -3, PercentSecondCrashing: -5,
		}, Config{
			Ticker: "MIX", Cycle: 600, PercentCrashing: -3, PercentSecondCrashing: -5,
		})

	for _, config := range list {
		go Queue.Coin(
			config.Ticker,
			config.Cycle,
			config.PercentCrashing,
			config.PercentSecondCrashing)
	}

	for i := 0; i < 1; i++ {
		<-loop
	}

}
