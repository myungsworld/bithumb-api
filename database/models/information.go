package models

import "time"

type Information struct {
	Ticker    string    `json:"ticker"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
