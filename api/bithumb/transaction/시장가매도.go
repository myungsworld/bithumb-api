package Execute

import (
	"encoding/json"
	"fmt"
	Middlewares "myungsworld/middlewares"
)

type MarketSelling struct {
	Status  string `json:"status"`
	OrderId string `json:"order_id"`
	Message string `json:"message"`
}

//EA 는 코인의 물량임
func MarketSell(ticker string, EA float64) (string, string, string) {
	const ENDPOINT = "/trade/market_sell"
	const PARAMS = "order_currency=주문통화&payment_currency=KRW&units=코인갯수"

	// 소수점 4자리까지만 지원됨 고로 BTC ,ETH 등 덩치가 큰건 지금 알고리즘과 맞지않음
	//EA = math.Floor(EA)
	each := fmt.Sprintf("%.4f", EA)

	params := fmt.Sprintf("order_currency=%s&payment_currency=KRW&units=%s", ticker, each)
	respData := Middlewares.Call(ENDPOINT, params)

	marketSelling := MarketSelling{}
	if err := json.Unmarshal(respData, &marketSelling); err != nil {
		panic(err)
	}

	if marketSelling.Status == "0000" {
		fmt.Printf("%s 코인 %s개 시장가로 매도 체결!\n", ticker, each)
	} else {
		fmt.Println("-------시장가 매도 실패-------")
		fmt.Printf("Status Code : %s \n%s\n", marketSelling.Status, marketSelling.Message)
	}

	return marketSelling.Status, marketSelling.Message, each
}
