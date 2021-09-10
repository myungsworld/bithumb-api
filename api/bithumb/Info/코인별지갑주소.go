package Info

import (
	"encoding/json"
	"fmt"
	Middlewares "myungsworld/middlewares"
)

type Wallet struct {
	Status string  `json:"status"`
	Data   Address `json:"data"`
}

type Address struct {
	WalletAddress string `json:"wallet_address"`
	Currency string `json:"currency"`
}

func FetchAddress(ticker string) (address string) {
	const ENDPOINT = "/info/wallet_address"
	const PARAMS = "currency="
	params := PARAMS + ticker

	respData := Middlewares.Call(ENDPOINT, params)

	fmt.Println(string(respData))

	wallet := Wallet{}
	if err := json.Unmarshal(respData,&wallet); err != nil {
		panic(err)
	}

	return wallet.Data.WalletAddress
}
