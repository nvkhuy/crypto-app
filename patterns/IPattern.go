package patterns

import (
	"github.com/adshao/go-binance/v2"
)

type IPattern interface {
	CaseMatching(klines []*binance.Kline) string
}
