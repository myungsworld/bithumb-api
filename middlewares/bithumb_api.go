package Middlewares

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const URL = "https://api.bithumb.com"

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

func Call(endPoint, params string) (data []byte) {
	apiKey, secretKey := FetchBithumbKey()
	eEndPoint := url.QueryEscape(endPoint)
	params += "&endpoint=" + eEndPoint

	hmacKey := secretKey
	nonceInt64 := microsectime()
	apiNonce := fmt.Sprint(nonceInt64)

	hmacData := endPoint + string(0) + params + string(0) + apiNonce
	hashHmacStr := hashHmac(hmacKey, hmacData)
	apiSign := hashHmacStr

	req, err := http.NewRequest("POST", URL+endPoint, bytes.NewBufferString(params))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Api-Key", apiKey)
	req.Header.Add("Api-Sign", apiSign)
	req.Header.Add("Api-Nonce", apiNonce)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	contentLengthStr := strconv.Itoa(len(params))
	req.Header.Add("Content-Length", contentLengthStr)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return respData

}
