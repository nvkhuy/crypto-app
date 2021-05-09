package bullish

import (
	"github.com/adshao/go-binance/v2"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/enums"
	"log"
	"math"
	"strconv"
)

func typedKlines(ks []*binance.Kline) []*MK {
	var mks []*MK
	for _, k := range ks {
		mks = append(mks, toMK(k))
	}
	return mks
}

func reverseMks(ks []*MK) []*MK {
	for l, r := 0, len(ks)-1; l < r; l, r = l+1, r-1 {
		ks[l], ks[r] = ks[r], ks[l]
	}
	return ks[1:]
}
func toMK(k *binance.Kline) *MK {
	h, _ := strconv.ParseFloat(k.High, 64)
	l, _ := strconv.ParseFloat(k.Low, 64)
	c, _ := strconv.ParseFloat(k.Close, 64)
	o, _ := strconv.ParseFloat(k.Open, 64)
	return &MK{h, o, c, l}
}

func ABS(x float64) float64 {
	return math.Abs(x)
}
func logCountEq(c string, bs ...bool) {
	cnt := 0
	for _, b := range bs {
		if b == true {
			cnt += 1
		} else {
			break
		}
	}
	log.Printf("%s %s Case# %s Count=[%d/%d]", LogBullish, enums.Captain, c, cnt, len(bs))
}
