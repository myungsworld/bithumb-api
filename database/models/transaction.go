package models

import "time"

type Transaction struct {
	Type string `json:"type"`
	Ticker string `json:"ticker"`
	Amount string `json:"amount"`
	TotalKRW float64 `json:"total_krw"`

	CreatedAt time.Time `json:"created_at"`
}

