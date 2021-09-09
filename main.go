package main

import (
	"fmt"
	"myungsworld/api/bithumb/Info"
	"myungsworld/database"
	"strconv"
	"time"
)


func main() {


	database.ConnectDB()

	done := make(chan bool, 100)

	go func() {

		var startPriceEveryTenMin float64
		var marketPrice float64
		var highPrice float64

		BTTMarketPrice := Info.CoinMarketCondition("BTT")
		startPriceEveryTenMin, _ = strconv.ParseFloat(BTTMarketPrice, 64)

		var price float64
		for true {
			time.Sleep(time.Second * 1)
			fmt.Println("BTT 매수타임 호시탐탐 검색중")

			BTTMarketPrice := Info.CoinMarketCondition("BTT")
			fmt.Println("현재가 BTT :", BTTMarketPrice)
			price, _ = strconv.ParseFloat(BTTMarketPrice, 64)
			fmt.Println("price: ", price)


			marketPrice = price
			if marketPrice > highPrice {
				highPrice = marketPrice
			}


		 	fmt.Println("시작가 :",startPriceEveryTenMin,"현재가 :",marketPrice,"고가 :",highPrice)

			// 10분안에 -7% 이상 떨어지면 50퍼 매도
			fluctateRate := ((marketPrice - startPriceEveryTenMin)/marketPrice)*100
			fmt.Println(fluctateRate)


			if BTTMarketPrice < "4" {
				done <- true

			}
		}
	}()

	//go func () {
	//	var priceArr [2]float64
	//	var price float64
	//
	//	for true {
	//		// 10분마다 초기화
	//		var tenMin int
	//
	//
	//		for true {
	//
	//			// 10분마다 초기화
	//			tenMin++
	//			//fmt.Println("tenMin : " ,tenMin)
	//			if tenMin == 600 {
	//				priceArr[0] = 0
	//				priceArr[1] = 0
	//				price = 0
	//				break
	//			}
	//
	//			time.Sleep(time.Second*1)
	//			fmt.Println("BTT 매수타임 호시탐탐 검색중")
	//
	//			BTTMarketPrice := Info.CoinMarketCondition("BTT")
	//			fmt.Println("현재가 BTT :", BTTMarketPrice)
	//			price , _ = strconv.ParseFloat(BTTMarketPrice,64)
	//			fmt.Println("price: ", price)
	//
	//			// 10분마다 고가를 기록
	//			priceArr[0] = price
	//			if priceArr[0] > priceArr[1] {
	//				priceArr[1] = priceArr[0]
	//			}
	//
	//			fmt.Println(priceArr)
	//
	//			if BTTMarketPrice < "4" {
	//				done <- true
	//
	//			}
	//		}
	//	}
	//
	//}()


	for i := 0 ; i < 3 ; i++ {
		<-done
		fmt.Println("아항")
	}


	//Info.GetBalance("ALL")
	//Info.GetMyTickerBalance("BTT","KRW")

	//Info.GetBalance("BTT")
	//Info.FetchAddress("BTT")
	//Execute.WithDrawKRW("302-0709-1079-11","1000")
	//Info.PendingOrder("BTT")
	//Execute.BuyHook("BTT","5000","5")
	//Execute.SellHook("BTT","500","6")
	//Execute.MarKetBuy("BTT","2")
	//Execute.MarketSell("BTT","1")
	//Info.CandleStick("BTT","24h")
	//Execute.MarketBuyHook("BTT","4","4","1500")
}
