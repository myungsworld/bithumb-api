package Execute

import (
	"encoding/json"
	"fmt"
	Middlewares "myungsworld/middlewares"
)

type Selling struct {
	Status  string `json:"status"`
	OrderId string `json:"order_id"`
	Message string `json:"message"`
}

func SellHook(ticker string , units string, price string) {
	const ENDPOINT = "/trade/place"
	params := fmt.Sprintf("order_currency=%s&payment_currency=KRW&units=%s&price=%s&type=ask",ticker,units,price)

	respData := Middlewares.Call(ENDPOINT,params)
	fmt.Println(string(respData))
	selling := Selling{}
	if err := json.Unmarshal(respData, &selling); err != nil {
		panic(err)
	}

	if selling.Status == "0000" {
		fmt.Println("----------매도 예약 API 성공-----------")
		fmt.Println("매도예약 체결")
		fmt.Printf("%s 토큰 %s원에 %s원 매도예약\n",ticker,price,units	)
	} else {
		fmt.Println("----------매도 예약 API 실패------------")
		fmt.Printf("매도예약 실패\n")
		fmt.Printf("StatusCode : %s\n", selling.Status)
		fmt.Printf("Message : %s\n", selling.Message)
	}

}