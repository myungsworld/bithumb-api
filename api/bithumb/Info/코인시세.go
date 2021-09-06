package Info

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Coin struct {
	Status string `json:"status"`
	Data   struct {
		OpeningPrice     string `json:"opening_price"`
		ClosingPrice     string `json:"closing_price"`
		MinPrice         string `json:"min_price"`
		MaxPrice         string `json:"max_price"`
		UnitsTraded      string `json:"units_traded"`
		AccTradeValue    string `json:"acc_trade_value"`
		PrevClosingPrice string `json:"prev_closing_price"`
		UnitsTraded24H   string `json:"units_traded_24H"`
		Fluctate24H      string `json:"fluctate_24H"`
		FluctateRate24H  string `json:"fluctate_rate_24H"`
		Date             string `json:"date"`
	}
}

func CoinMarketCondition(ticker string) {

	reqURL := fmt.Sprintf("https://api.bithumb.com/public/ticker/%s_KRW",ticker)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("string(bytes) : ", string(bytes))
	var coin Coin
	if err := json.Unmarshal(bytes, &coin); err != nil {
		panic(err)
	}
	fmt.Printf("Ticker : %s\n", ticker)
	fmt.Printf("당일 시작가 : %s\n", coin.Data.OpeningPrice)
	fmt.Printf("현재가 : %s\n", coin.Data.ClosingPrice)
	fmt.Printf("당일 최고점 : %s\n", coin.Data.MaxPrice)
	fmt.Printf("당일 최저점 : %s\n", coin.Data.MinPrice)
	fmt.Printf("최근 24시간 변동가 : %s\n",coin.Data.Fluctate24H)
	fmt.Printf("최근 24시간 변동 퍼센트 : %s\n",coin.Data.FluctateRate24H)

}
