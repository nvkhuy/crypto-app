package bullish

import (
	"github.com/adshao/go-binance/v2"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/enums"
	"log"
	"math"
)

func (b *bullish) CaseMatching(klines []*binance.Kline) string {
	ks := typedKlines(klines)
	ks = reverseMks(ks)
	log.Print("________________")
	if b.isAbandonedBaby(ks) {
		return enums.AbandonedBaby
	}
	if b.isDojiDragonfly(ks) {
		return enums.DojiDragonfly
	}
	if b.isHarami(ks) {
		return enums.Harami
	}
	if b.isHommingPigeonz(ks) {
		return enums.HommingPigeon
	}
	if b.isInvertedHammer(ks) {
		return enums.InvertedHammer
	}
	if b.isMorningStar(ks) {
		return enums.MorningStar
	}
	if b.isPiercingLine(ks) {
		return enums.PiercingLine
	}
	if b.isMatchingLow(ks) {
		return enums.MatchingLow
	}
	return ""
}
func (b *bullish) isAbandonedBaby(ks []*MK) bool {
	if len(ks) < 3 {
		return false
	}
	var O, C, L, H, O1, C1, O2, C2, L1, L2, H1, H2 float64
	O, C, L, H = ks[0].O, ks[0].C, ks[0].L, ks[0].H
	O1, C1 = ks[1].O, ks[1].C
	O2, C2 = ks[2].O, ks[2].C
	H1, L1 = ks[1].H, ks[1].L
	H2, L2 = ks[2].H, ks[2].L
	eq1 := 2*ABS(C2-O2) > H2-L2
	eq2 := C2 > O2
	eq3 := 20*ABS(C1-O1) <= H1-L1
	eq4 := 5*((C1+O1)/2-L1) >= 2*(H1-L1)
	eq5 := 5*((C1+O1)/2-L1) <= 3*(H1-L1)
	eq6 := L1 > H2
	eq7 := C < O
	eq8 := H < L1
	eq9 := O > C2
	eq10 := L > O2 && C < L2
	logCountEq(enums.AbandonedBaby, eq1, eq2, eq3, eq4, eq5, eq6, eq7, eq8, eq9, eq10)
	if eq1 && eq2 && eq3 && eq4 && eq5 && eq6 && eq7 && eq8 && eq9 && eq10 {
		return true
	}
	return false
}
func (b *bullish) isInvertedHammer(ks []*MK) bool {
	if len(ks) < 10 {
		return false
	}
	var O, C, H, L, H1, L1, STOC1, AVGH10, AVGL10, MINL5 float64
	O, C, H, L = ks[0].O, ks[0].C, ks[0].H, ks[0].L
	H1, L1 = math.Max(H, ks[1].H), math.Min(L, ks[1].L)
	STOC1 = ((C - L1) / (H1 - L1)) * 100
	MINL5 = L
	for i := 0; i < 10; i += 1 {
		n := float64(i)
		AVGH10 = AVGH10*n/(n+1) + ks[i].H/(n+1)
		AVGL10 = AVGL10*n/(n+1) + ks[i].L/(n+1)
	}
	for i := 0; i < 5; i += 1 {
		MINL5 = math.Min(MINL5, ks[i].L)
	}
	eq1 := 5*ABS(O-C) <= H-L
	eq2 := 10*ABS(O-C) >= H-L
	eq3 := 2*(H-O) >= H-L
	eq4 := 2*(H-C) >= H-L
	eq5 := 2*(O-L) <= H-L || 20*(C-L) <= H-L
	eq6 := 5*(H-L) >= 4*(AVGH10-AVGH10)
	eq7 := 2*O <= H1+L1
	eq8 := STOC1 <= 50
	eq9 := L == MINL5
	eq10 := H > L
	logCountEq(enums.InvertedHammer, eq1, eq2, eq3, eq4, eq5, eq6, eq7, eq8, eq9, eq10)
	if eq1 && eq2 && eq3 && eq4 && eq5 && eq6 && eq7 && eq8 && eq9 && eq10 {
		return true
	}
	return false
}
func (b *bullish) isDojiDragonfly(ks []*MK) bool {
	if len(ks) < 10 {
		return false
	}
	var O, C, H, L, H1, L1, STOC1, AVGH10, AVGL10, MINL10 float64
	O, C, H, L = ks[0].O, ks[0].C, ks[0].H, ks[0].L
	H1, L1 = ks[1].H, ks[1].L
	STOC1 = ((C - L1) / (H1 - L1)) * 100
	MINL10 = ks[0].L
	for i := 0; i < 10; i += 1 {
		n := float64(i)
		AVGH10 = AVGH10*n/(n+1) + ks[i].H/(n+1)
		AVGL10 = AVGL10*n/(n+1) + ks[i].L/(n+1)
		MINL10 = math.Min(MINL10, ks[i].L)
	}
	eq1 := 50*ABS(O-C) <= H-L
	eq2 := STOC1 >= 70
	eq3 := H-L >= AVGH10-AVGL10
	eq4 := L == MINL10
	logCountEq(enums.DojiDragonfly, eq1, eq2, eq3, eq4)
	if eq1 && eq2 && eq3 && eq4 {
		return true
	}
	return false
}
func (b *bullish) isHarami(ks []*MK) bool {
	if len(ks) < 11 {
		return false
	}
	var O, C, O1, C1, H1, L1, AVGH10, AVGL10 float64
	O, C = ks[0].O, ks[0].C
	O1, C1 = ks[1].O, ks[1].C
	H1, L1 = ks[1].H, ks[1].L
	for i := 1; i < 11; i += 1 {
		n := float64(i)
		AVGH10 = AVGH10*n/(n+1) + ks[i].H/(n+1)
		AVGL10 = AVGL10*n/(n+1) + ks[i].L/(n+1)
	}
	eq1 := 10*ABS(O1-C1) >= 7*(H1-L1)
	eq2 := H1-L1 >= AVGH10-AVGL10
	eq3 := C > O
	eq4 := O > C1
	eq5 := O1 > C
	eq6 := 6*(O1-C1) >= 10*(C-O)
	logCountEq(enums.Harami, eq1, eq2, eq3, eq4, eq5, eq6)
	if eq1 && eq2 && eq3 && eq4 && eq5 && eq6 {
		return true
	}
	return false
}
func (b *bullish) isHommingPigeonz(ks []*MK) bool {
	if len(ks) < 3 {
		return false
	}
	var O, C, H, L, O1, C1, H1, L1 float64
	O, C, H, L = ks[0].O, ks[0].C, ks[0].H, ks[0].L
	O1, C1 = ks[1].O, ks[1].C
	H1, L1 = ks[1].H, ks[1].L
	eq1 := C1 < O1
	eq2 := 5*ABS(C-O) >= 3*(H1-L1)
	eq3 := 2*ABS(C1-O1) >= H1-L1
	eq4 := H < O1
	eq5 := L > C1
	eq6 := C < O
	eq7 := 6*(O1-C1) >= 10*(C-O)
	logCountEq(enums.HommingPigeon, eq1, eq2, eq3, eq4, eq5, eq6, eq7)
	if eq1 && eq2 && eq3 && eq4 && eq5 && eq6 && eq7 {
		return true
	}
	return false
}
func (b *bullish) isMorningStar(ks []*MK) bool {
	if len(ks) < 3 {
		return false
	}
	var O, C, O1, C1, O2, C2, H2, L2 float64
	O, C = ks[0].O, ks[0].C
	O1, C1 = ks[1].O, ks[1].C
	O2, C2, H2, L2 = ks[2].O, ks[2].C, ks[2].H, ks[2].L
	eq1 := O2 > C2
	eq2 := 5*(O2-C2) > 3*(H2-L2)
	eq3 := C2 > O1
	eq4 := 2*ABS(O1-C1) < ABS(O2-C2)
	eq5 := C > O
	eq6 := O > O1
	eq7 := O > C1
	logCountEq(enums.MorningStar, eq1, eq2, eq3, eq4, eq5, eq6, eq7)
	if eq1 && eq2 && eq3 && eq4 && eq5 && eq6 && eq7 {
		return true
	}
	return false
}
func (b *bullish) isPiercingLine(ks []*MK) bool {
	if len(ks) < 11 {
		return false
	}
	var O, C, O1, C1, H1, L1, AVGH10, AVGL10 float64
	O, C = ks[0].O, ks[0].C
	O1, C1 = ks[1].O, ks[1].C
	H1, L1 = ks[1].H, ks[1].L
	for i := 1; i < 11; i += 1 {
		n := float64(i)
		AVGH10 = AVGH10*n/(n+1) + ks[i].H/(n+1)
		AVGL10 = AVGL10*n/(n+1) + ks[i].L/(n+1)
	}

	eq1 := O1 > C1
	eq2 := H1-L1 >= AVGH10-AVGL10
	eq3 := O < C1
	eq4 := 2*C > C1+O1
	eq5 := C < O1
	logCountEq(enums.PiercingLine, eq1, eq2, eq3, eq4, eq5)
	if eq1 && eq2 && eq3 && eq4 && eq5 {
		return true
	}
	return false
}
func (b *bullish) isMatchingLow(ks []*MK) bool {
	if len(ks) < 11 {
		return false
	}
	var O, C, O1, C1, H1, L1 float64
	O, C = ks[0].O, ks[0].C
	O1, C1 = ks[1].O, ks[1].C
	H1, L1 = ks[1].H, ks[1].L

	eq1 := C1 < O1
	eq2 := 20*ABS(C1-O1) > H1-L1
	eq3 := C < O
	eq4 := 100*ABS(C/C1-1) < 1
	logCountEq(enums.MatchingLow, eq1, eq2, eq3, eq4)
	if eq1 && eq2 && eq3 && eq4 {
		return true
	}
	return false
}
