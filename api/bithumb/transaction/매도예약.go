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

func Sell(ticker string , units string, price string) {
	const ENDPOINT = "/trade/place"
	params := fmt.Sprintf("order_currency=%s&payment_currency=KRW&units=%s&price=%s&type=ask",ticker,units,price)

	respData := Middlewares.Call(ENDPOINT,params)
	fmt.Println(string(respData))
	selling := Selling{}
	if err := json.Unmarshal(respData, &selling); err != nil {
		panic(err)
	}
	fmt.Println(selling)
}