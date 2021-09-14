package main

import (
	"myungsworld/database"
	"myungsworld/queue"
)

func main() {

	database.ConnectDB()

	loop := make(chan bool, 1)

	tickers := []string{
		"XEC", "FIT", "DAC", "AMO", "TMTG",
		"CON", "MIX", "BTT", "EGG", "EM",
		"HIBS", "TEMCO", "EL", "OBSR", "XPR",
		"XPR", "WIKEN", "BASIC", "GOM2", "MBL",
		"FLETA", "QTCON", "TRV", "CKB", "AWO",
	}

	for _, ticker := range tickers {
		go Queue.Coin(ticker)
	}

	for i := 0; i < 1; i++ {
		<-loop
	}

}
