package CoinGecko

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Coin struct {
	Id     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func GetCoin() []Coin {
	req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/coins/list", nil)
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
	var dat []Coin
	if err := json.Unmarshal(bytes, &dat); err != nil {
		panic(err)
	}

	return dat
}
