package main

import (
	Execute "myungsworld/api/bithumb/transaction"
	"myungsworld/database"
)

func main() {

	database.ConnectDB()

	//Info.CoinMarketCondition("BTT")
	//Info.GetMyTickerBalance("BTT","KRW")
	//Info.GetBalance("ALL")
	//Info.GetBalance("BTT")
	//Info.FetchAddress("BTT")
	//Execute.WithDrawKRW("302-0709-1079-11","1000")
	//Info.PendingOrder("BTT")
	//Execute.BuyHook("BTT","5000","5")
	//Execute.SellHook("BTT","500","6")
	//Execute.MarKetBuy("BTT","2")
	//Execute.MarketSell("BTT","1")
	//Info.CandleStick("BTT","24h")
	Execute.MarketBuyHook("BTT","4","4","1500")
}
