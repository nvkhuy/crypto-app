package global

import "github.com/adshao/go-binance/v2"

var Api *binance.Client

func SetApi() {
	var (
		apiKey    = ""
		secretKey = ""
	)
	Api = binance.NewClient(apiKey, secretKey)
}
