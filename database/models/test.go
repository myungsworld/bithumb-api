package models

type Test struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Coin string `json:"coin"`
}
