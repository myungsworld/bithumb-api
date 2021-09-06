package Execute

import (
	"encoding/json"
	"fmt"
	Middlewares "myungsworld/middlewares"
)

type Buying struct {
	Status  string `json:"status"`
	OrderId string `json:"order_id"`
	Message string `json:"message"`
}


func Buy(ticker string, units string, price string) {
	const ENDPOINT = "/trade/place"
	const PARAMS = "order_currency=주문통화&payment_currency=결제통화&units=주문수량(최대 50억 최소 500원)&price=거래가&type=거래유형(bid매수)"

	params := fmt.Sprintf("order_currency=%s&payment_currency=KRW&units=%s&price=%s&type=bid", ticker, units, price)
	respData := Middlewares.Call(ENDPOINT, params)
	buying := Buying{}
	if err := json.Unmarshal(respData, &buying); err != nil {
		panic(err)
	}
	if buying.Status == "0000" {
		fmt.Println("----------대충 체결된 연결선-----------")
		fmt.Printf("매수예약 체결\n")
		fmt.Printf("%s 토큰 %s원에 %s원 매수예약\n",ticker,price,units	)
	} else {
		fmt.Println("----------대충 실패한 연결선------------")
		fmt.Printf("매수예약 실패\n")
		fmt.Printf("StatusCode : %s\n", buying.Status)
		fmt.Printf("Message : %s\n", buying.Message)
	}

}
