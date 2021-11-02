package main

import (
	"myungsworld/database"
	Queue "myungsworld/queue"
)

// Ticker : 코인티커 , Cycle : 회복주기 , PercentCrashing : 폭락방지 한계 퍼센트
type Config struct {
	Ticker          string
	Cycle           int
	PercentCrashing float64
}

func main() {

	database.ConnectDB()

	loop := make(chan bool, 1)

	list := make([]Config, 0)
	list = append(list,
		Config{
			Ticker: "BTT", Cycle: 600, PercentCrashing: -3,
		}, Config{
			Ticker: "MIX", Cycle: 600, PercentCrashing: -3,
		})

	for _, config := range list {
		go Queue.Coin(config.Ticker, config.Cycle, config.PercentCrashing)
	}

	for i := 0; i < 1; i++ {
		<-loop
	}

}
