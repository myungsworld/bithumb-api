package Middlewares

import (
	"os"
)

func FetchBithumbKey() (apiKey , secretKey string) {
	BithumbApiKey := os.Getenv("BithumbApiKey")
	BithumbApiSecretKey := os.Getenv("BithumbApiSecretKey")
	return BithumbApiKey,BithumbApiSecretKey
}