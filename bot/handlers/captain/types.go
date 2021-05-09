package captain

import (
	"github.com/adshao/go-binance/v2"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/patterns"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/patterns/bearish"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/patterns/bullish"
	"sync"
	"time"
)

var serviceCycle Tracer
var arbitrageCycle Tracer

var doOnce sync.Once
var bull patterns.IPattern
var bear patterns.IPattern

const CaptainLog = "[CAPTAIN HANDLER]"

type rsiFormula func(kLines []*binance.Kline) float64

type Tracer struct {
	symbol      string
	ticker      *time.Ticker
	isVisualize *time.Ticker
	isTrade     *time.Ticker
	base        float64
	quote       float64
	stop        chan bool
}

func init() {
	doOnce.Do(func() {
		bull = bullish.GetBullish()
		bear = bearish.GetBearish()
	})
}
