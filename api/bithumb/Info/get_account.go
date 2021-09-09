package Info

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"myungsworld/middlewares"
	"os"
	"time"
)

func microsectime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func hashHmac(hmacKey string, hmacData string) (hashHmacStr string) {
	hmh := hmac.New(sha512.New, []byte(hmacKey))
	hmh.Write([]byte(hmacData))

	hexData := hex.EncodeToString(hmh.Sum(nil))
	hashHmacBytes := []byte(hexData)
	hmh.Reset()

	hashHmacStr = base64.StdEncoding.EncodeToString(hashHmacBytes)

	return hashHmacStr
}

type AccountRec struct {
	Created   int64   `json:"created,string"`
	AccountId string  `json:"account_id"`
	TradeFee  float64 `json:"trade_fee,string"`
	Balance   float64 `json:"balance,string"`
}

type Account struct {
	Status string     `json:"status"`
	Data   AccountRec `json:"data"`
}

func GetMyTickerBalance(orderCurrency string) {

	const ENDPOINT = "/info/account"
	const PARAMS = "order_currency=BTC&payment_currency=KRW"

	params := fmt.Sprintf("order_currency=%s&payment_currency=KRW",orderCurrency)

	respData := Middlewares.Call(ENDPOINT, params)
	accountJsonRecInfo := Account{}
	if err := json.Unmarshal(respData, &accountJsonRecInfo); err != nil {
		panic(err)
	}

	fmt.Printf("- Status Code: %s\n", accountJsonRecInfo.Status)
	fmt.Printf("- Created: %d\n", accountJsonRecInfo.Data.Created)
	fmt.Printf("- Account ID: %s\n", accountJsonRecInfo.Data.AccountId)
	fmt.Printf("- Trade Fee: %.4f\n", accountJsonRecInfo.Data.TradeFee)
	fmt.Printf("- Balance: %.8f\n", accountJsonRecInfo.Data.Balance)

	os.Exit(0)

}
