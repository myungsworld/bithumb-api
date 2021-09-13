package main

import (
	"myungsworld/database"
	"myungsworld/queue"
)

func main() {

	database.ConnectDB()

	loop := make(chan bool, 1)

	go Queue.Coin("BTT")
	//go Queue.Coin("ETH")
	//go Queue.Coin("BTT")

	for i := 0; i < 1; i++ {
		<-loop
	}

}
