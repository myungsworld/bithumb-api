package Execute

import (
	"fmt"
	Middlewares "myungsworld/middlewares"
)

func WithDrawKRW(account string, price string) {
	const ENDPOINT = "/trade/krw_withdrawal"
	const PARAMS = "bank=011_농협은행&account=계좌번호&price=출금금액"

	params := fmt.Sprintf("bank=011_농협은행&account=%s&price=%s", account, price)
	respData := Middlewares.Call(ENDPOINT,params)

	fmt.Println(string(respData))
}
