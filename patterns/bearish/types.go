package bearish

import (
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/patterns"
)

const LogBearish = "[BEARISH]"

type MK struct {
	H float64
	O float64
	C float64
	L float64
}

type bearish struct{}

func GetBearish() patterns.IPattern {
	b := &bearish{}
	return b
}
