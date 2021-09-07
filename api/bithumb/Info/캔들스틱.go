package Info

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Time int `json:"time"`
	MarketPrice float64 `json:"market_price"`
}

type CandleStickResp struct {
	Status string `json:"status"`
	Data  []interface{}
}

func CandleStick(ticker, chartInterval string) {
	reqURL := fmt.Sprintf("https://api.bithumb.com/public/candlestick/%s_KRW/%s", ticker, chartInterval)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	var candleStick CandleStickResp
	if err := json.Unmarshal(bytes, &candleStick); err != nil {
		fmt.Println(err)
		panic(err)
	}

	FetchNum := 5

	//result := make([]Result,FetchNum)
	for i := len(candleStick.Data) ; i > len(candleStick.Data) - FetchNum ; i-- {
		fmt.Println(candleStick.Data[i-1])
	}

}
