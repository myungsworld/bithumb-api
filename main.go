package main

import (
	"myungsworld/database"
	"myungsworld/queue"
)

const (
	BTT = "BTT"

)

func main() {

	database.ConnectDB()

	loop := make(chan bool, 1)

	go Queue.Coin("BTT")
	go Queue.Coin("LTC")
	go Queue.Coin("ADA")
	go Queue.Coin("ASM")
	go Queue.Coin("XEC")
	go Queue.Coin("XLM")

	for i := 0; i < 1; i++ {
		<-loop
	}

}
