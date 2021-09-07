package Execute

import (
	"encoding/json"
	"fmt"
	Middlewares "myungsworld/middlewares"
)

type MarketBuying struct {
	Status  string `json:"status"`
	OrderId string `json:"order_id"`
	Message string `json:"message"`
}

// EA : each
func MarKetBuy(ticker, EA string) {
	const ENDPOINT = "/trade/market_buy"
	const PARAMS = "order_currency=주문통화&payment_currency=KRW&units=주문금액"

	params := fmt.Sprintf("order_currency=%s&payment_currency=KRW&units=%s", ticker, EA)

	respData := Middlewares.Call(ENDPOINT, params)
	marketBuying := MarketBuying{}
	if err := json.Unmarshal(respData, &marketBuying); err != nil {
		panic(err)
	}
	if marketBuying.Status == "0000" {
		fmt.Printf("%s 코인 %s개 시장가로 매수 체결\n", ticker, EA)
	} else {
		fmt.Println("-------시장가 매수 실패-------")
		fmt.Printf("Status Code : %s \n%s\n", marketBuying.Status, marketBuying.Message)
	}
}
