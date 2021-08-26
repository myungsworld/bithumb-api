package API

import (
	"encoding/json"
	"fmt"
	"myungsworld/middlewares"
)

type Detail struct {
	TotalKRW     float64 `json:"total_krw,string"`
	TotalBTT     float64 `json:"total_btt,string"`
	InUseBTT     float64 `json:"in_use_btt,string"`
	AvailableBTT float64 `json:"available_btt,string"`
	TotalBNT     float64 `json:"total_bnt,string"`
	InUseKRW     float64 `json:"in_use_krw,string"`

	AvailableKRW float64 `json:"available_krw,string"`
	XCoinLastBTT float64 `json:"xcoin_last_btt,string"`
}

type Balance struct {
	Status string `json:"status"`
	Data   Detail `json:"data"`
}

func GetBalance(ticker string) {
	const ENDPOINT = "/info/balance"
	const PARAMS = "currency="

	params := PARAMS + ticker

	respData := Middlewares.Call(ENDPOINT, params)

	fmt.Println(string(respData))

	balanceInfo := Balance{}
	if err := json.Unmarshal(respData, &balanceInfo); err != nil {
		panic(err)
	}

	fmt.Printf("- 상태 코드 : %s\n", balanceInfo.Status)
	fmt.Printf("- 총 원화 : %.f원\n", balanceInfo.Data.TotalKRW)
	fmt.Printf("- 뱅코르 : %.4f\n", balanceInfo.Data.TotalBNT)
	fmt.Printf("- 비트토렌트 : %.8f\n", balanceInfo.Data.TotalBTT)
	fmt.Printf("- 주문가능 원화 : %.f원\n", balanceInfo.Data.AvailableKRW)
	fmt.Printf("- 주문에 묶여있는 원화 : %.f원\n", balanceInfo.Data.InUseKRW)
}
