package models

import "time"

type Transaction struct {
	Type        string  `json:"type"`
	Ticker      string  `json:"ticker"`
	Amount      string  `json:"amount"`
	TotalKRW    float64 `json:"total_krw"`
	StartPrice  float64 `json:"start_price"`
	MarketPrice float64 `json:"market_price"`
	Fluctate    float64 `json:"fluctate"`
	Seconds     int     `json:"seconds"`
	Content     string  `json:"content"`

	CreatedAt time.Time `json:"created_at"`
}
