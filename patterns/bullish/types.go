package bullish

import "github.com/khanhhuynguyenvu/crypto-trading-heroku/patterns"

const LogBullish = "[BULLISH]"

type MK struct {
	H float64
	O float64
	C float64
	L float64
}

type bullish struct{}

func GetBullish() patterns.IPattern {
	b := &bullish{}
	return b
}
