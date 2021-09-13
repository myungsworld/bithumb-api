package models

import "time"

type Information struct {
	Ticker  string `json:"ticker"`
	Content string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
}
