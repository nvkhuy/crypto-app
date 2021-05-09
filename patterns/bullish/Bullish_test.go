package bullish

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/enums"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/patterns"
	"testing"
)

var bull patterns.IPattern

func init() {
	bull = GetBullish()
}

func TestReverseMK(t *testing.T) {
	e := &binance.Kline{High: "6", Open: "1", Close: "1", Low: "1"}
	e1 := &binance.Kline{High: "5", Open: "1", Close: "1", Low: "1"}
	e2 := &binance.Kline{High: "4", Open: "1", Close: "1", Low: "1"}
	e3 := &binance.Kline{High: "3", Open: "1", Close: "1", Low: "1"}
	e4 := &binance.Kline{High: "2", Open: "1", Close: "1", Low: "1"}
	e5 := &binance.Kline{High: "1", Open: "1", Close: "1", Low: "1"}
	e6 := &binance.Kline{High: "0", Open: "1", Close: "1", Low: "1"}

	klines := []*binance.Kline{e, e1, e2, e3, e4, e5, e6}
	res := reverseMks(typedKlines(klines))
	if len(res) == 0 {
		t.Errorf("error empty")
	}
	if res[0].H != float64(1) {
		t.Errorf("not reverse yet")
	}
}
func TestBullish_IsInvertedHammer(t *testing.T) {
	e := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "0"}
	e1 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "0"}
	e2 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "0"}
	e3 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "0"}
	e4 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "0"}
	e5 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "0"}
	e6 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "0"}
	e7 := &binance.Kline{High: "0", Open: "1", Close: "1", Low: "1"}
	e8 := &binance.Kline{High: "0", Open: "1", Close: "1", Low: "1"}
	e9 := &binance.Kline{High: "12", Open: "1", Close: "1", Low: "8"}
	e10 := &binance.Kline{High: "12", Open: "1", Close: "1", Low: "7"}
	e11 := &binance.Kline{High: "12", Open: "1", Close: "1", Low: "6"}
	e12 := &binance.Kline{High: "12", Open: "1", Close: "1", Low: "5"}
	e13 := &binance.Kline{High: "12", Open: "1", Close: "1", Low: "4"}
	e14 := &binance.Kline{High: "10", Open: "4", Close: "3", Low: "2"}
	e15 := &binance.Kline{High: "0", Open: "1", Close: "1", Low: "1"}
	klines := []*binance.Kline{e, e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15}
	fmt.Println(bull.CaseMatching(klines))
	if bull.CaseMatching(klines) != enums.InvertedHammer {
		t.Errorf("wrong inverted hammer case")
	}
}
func TestBullish_IsDojiDragonfly(t *testing.T) {
	e := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "20"}
	e1 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "20"}
	e2 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "20"}
	e3 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "20"}
	e4 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "0"}
	e5 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "19"}
	e6 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "18"}
	e7 := &binance.Kline{High: "0", Open: "1", Close: "1", Low: "17"}
	e8 := &binance.Kline{High: "0", Open: "1", Close: "1", Low: "16"}
	e9 := &binance.Kline{High: "12", Open: "1", Close: "1", Low: "15"}
	e10 := &binance.Kline{High: "12", Open: "1", Close: "1", Low: "14"}
	e11 := &binance.Kline{High: "12", Open: "1", Close: "1", Low: "13"}
	e12 := &binance.Kline{High: "12", Open: "1", Close: "1", Low: "12"}
	e13 := &binance.Kline{High: "90", Open: "1", Close: "1", Low: "10"}
	e14 := &binance.Kline{High: "100", Open: "90", Close: "90", Low: "10"}
	e15 := &binance.Kline{High: "0", Open: "1", Close: "1", Low: "1"}
	klines := []*binance.Kline{e, e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15}
	fmt.Println(bull.CaseMatching(klines))
	if bull.CaseMatching(klines) != enums.DojiDragonfly {
		t.Errorf("wrong doji dragonfly case")
	}
}
func TestBullish_IsHarami(t *testing.T) {
	e := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "20"}
	e1 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "20"}
	e2 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "20"}
	e3 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "20"}
	e4 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "1"}
	e5 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "1"}
	e6 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "1"}
	e7 := &binance.Kline{High: "0", Open: "1", Close: "1", Low: "1"}
	e8 := &binance.Kline{High: "0", Open: "1", Close: "1", Low: "1"}
	e9 := &binance.Kline{High: "1", Open: "1", Close: "1", Low: "1"}
	e10 := &binance.Kline{High: "1", Open: "1", Close: "1", Low: "1"}
	e11 := &binance.Kline{High: "1", Open: "1", Close: "1", Low: "1"}
	e12 := &binance.Kline{High: "1", Open: "1", Close: "1", Low: "1"}
	e13 := &binance.Kline{High: "100", Open: "90", Close: "25", Low: "10"}
	e14 := &binance.Kline{High: "100", Open: "30", Close: "40", Low: "10"}
	e15 := &binance.Kline{High: "0", Open: "1", Close: "1", Low: "1"}
	klines := []*binance.Kline{e, e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15}
	fmt.Println(bull.CaseMatching(klines))
	if bull.CaseMatching(klines) != enums.Harami {
		t.Errorf("wrong harami case")
	}
}

func TestBullish_IsHommingPigeon(t *testing.T) {
	e := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "20"}
	e1 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "20"}
	e2 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "20"}
	e3 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "20"}
	e4 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "1"}
	e5 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "1"}
	e6 := &binance.Kline{High: "0", Open: "0", Close: "0", Low: "1"}
	e7 := &binance.Kline{High: "0", Open: "1", Close: "1", Low: "1"}
	e8 := &binance.Kline{High: "0", Open: "1", Close: "1", Low: "1"}
	e9 := &binance.Kline{High: "1", Open: "1", Close: "1", Low: "1"}
	e10 := &binance.Kline{High: "1", Open: "1", Close: "1", Low: "1"}
	e11 := &binance.Kline{High: "1", Open: "1", Close: "1", Low: "1"}
	e12 := &binance.Kline{High: "1", Open: "1", Close: "1", Low: "1"}
	e13 := &binance.Kline{High: "100", Open: "90", Close: "20", Low: "10"}
	e14 := &binance.Kline{High: "80", Open: "76", Close: "22", Low: "21"}
	e15 := &binance.Kline{High: "0", Open: "1", Close: "1", Low: "1"}
	klines := []*binance.Kline{e, e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15}
	fmt.Println(bull.CaseMatching(klines))
	if bull.CaseMatching(klines) != enums.HommingPigeon {
		t.Errorf("wrong homming pigeon case")
	}
}
