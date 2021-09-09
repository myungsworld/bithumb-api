package Execute

import (
	"encoding/json"
	"fmt"
	Middlewares "myungsworld/middlewares"
)

const (
	OrderCurrency = "order_currency="
	PaymentCurrency = "payment_currency=KRW"
)

type MarketBuyHooking struct {
	Status string `json:"status"`
	OrderId string `json:"order_id"`
	Message string `json:"message"`
}


func MarketBuyHook(ticker,watchPrice,price,units string) {
	const ENDPOINT = "/trade/stop_limit"

	params := fmt.Sprintf(
		PaymentCurrency+"&"+
		OrderCurrency+ticker+"&"+
		"watch_price="+watchPrice+"&"+
		"price="+price+"&"+
		"units="+units+"&"+
		"type=bid",
	)

	respData := Middlewares.Call(ENDPOINT,params)
	marketBuyHooking := MarketBuyHooking{}
	if err := json.Unmarshal(respData, &marketBuyHooking); err != nil {
		panic(err)
	}
	if marketBuyHooking.Status == "0000" {
		fmt.Printf("성공")
	} else {
		fmt.Println("-------자동 주문 매수 실패-------")
		fmt.Printf("Status Code : %s \n%s\n", marketBuyHooking.Status, marketBuyHooking.Message)
	}


}
