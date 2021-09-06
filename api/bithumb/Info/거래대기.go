package Info

import (
	"encoding/json"
	"fmt"
	Middlewares "myungsworld/middlewares"
)

type Order struct {
	OrderCurrency   string `json:"order_currency"`
	PaymentCurrency string `json:"payment_currency"`
	OrderId         string `json:"order_id"`
	OrderDate       string `json:"order_date"`
	Type            string `json:"type"`
	Unit            string `json:"units"`
	UnitRemaining   string `json:"units_remaining"`
	Price           string `json:"price"`
}

type Orders struct {
	Status string  `json:"status"`
	Data   []Order `json:"data"`
}

func PendingOrder(ticker string) {
	const ENDPOINT = "/info/orders"
	const PARAMS = ""

	params := fmt.Sprintf("order_currency=%s&payment_currency=KRW", ticker)

	respData := Middlewares.Call(ENDPOINT, params)
	//fmt.Println(string(respData))

	orders := Orders{}
	if err := json.Unmarshal(respData, &orders); err != nil {
		panic(err)
	}

	fmt.Println(orders)

}
