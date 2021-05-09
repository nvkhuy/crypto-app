package bearish

import (
	"github.com/adshao/go-binance/v2"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/enums"
	"log"
	"math"
)

func (b *bearish) CaseMatching(klines []*binance.Kline) string {
	ks := typedKlines(klines)
	ks = reverseMks(ks)
	log.Print("________________")
	if b.isAbandonedBaby(ks) {
		return enums.AbandonedBaby
	}
	if b.isAdvanceBlock(ks) {
		return enums.AdvanceBlock
	}
	if b.isBeltHold(ks) {
		return enums.BeltHold
	}
	if b.isBreakaway(ks) {
		return enums.Breakaway
	}
	if b.isDarkCloudCover(ks) {
		return enums.DarkCloudCover
	}
	if b.isDeliberation(ks) {
		return enums.Deliberation
	}
	if b.isDownsideGapThreeMethods(ks) {
		return enums.DownsideGapThreeMethods
	}
	if b.isDownsideTasukiGap(ks) {
		return enums.DownsideTasukiGap
	}
	if b.isDojiStar(ks) {
		return enums.DojiStar
	}
	return ""
}

func (b *bearish) isAbandonedBaby(ks []*MK) bool {
	if len(ks) < 3 {
		return false
	}
	O, C, L, H := ks[0].O, ks[0].C, ks[0].L, ks[0].H
	O1, C1 := ks[1].O, ks[1].C
	O2, C2 := ks[2].O, ks[2].C
	H1, L1 := ks[1].H, ks[1].L
	H2, L2 := ks[2].H, ks[2].L
	eq1 := ABS(C2-O2) > 0.5*(H2-L2)
	eq2 := C2 > O2
	eq3 := ABS(C1-O1) <= 0.5*(H1-L1)
	eq4 := (C1+O1)/2-L1 >= 0.4*(H1-L1)
	eq5 := (C1+O1)/2-L1 <= 0.6*(H1-L1)
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
func (b *bearish) isAdvanceBlock(ks []*MK) bool {
	if len(ks) < 21 {
		return false
	}
	O, C, L, H := ks[0].O, ks[0].C, ks[0].L, ks[0].H
	O1, C1 := ks[1].O, ks[1].C
	O2, C2 := ks[2].O, ks[2].C
	H1, L1 := ks[1].H, ks[1].L
	H2, L2 := ks[2].H, ks[2].L
	AVGH21, AVGL21 := float64(0), float64(0)
	for i := 0; i < 21; i += 1 {
		n := float64(i)
		AVGH21 = AVGH21*n/(n+1) + ks[i].H/(n+1)
		AVGL21 = AVGL21*n/(n+1) + ks[i].L/(n+1)
	}
	eq1 := H-L > AVGH21-AVGL21
	eq2 := ABS(C1-O1) > .5*(H1-L1)
	eq3 := ABS(C2-O2) > .5*(H2-L2)
	eq4 := C > C1
	eq5 := C1 > C2
	eq6 := O1 > O2
	eq7 := O1 < C2
	eq8 := O > O1
	eq9 := O < C1
	eq10 := H-L < .8*(H1-L1)
	eq11 := H1-L1 < .8*(H2-L2)
	eq12 := H-C > O-L
	eq13 := H1-C1 > O1-L1
	logCountEq(enums.AdvanceBlock, eq1, eq2, eq3, eq4, eq5, eq6, eq7, eq8, eq9, eq10, eq11, eq12, eq13)
	if eq1 && eq2 && eq3 && eq4 && eq5 && eq6 && eq7 && eq8 && eq9 && eq10 && eq11 && eq12 && eq13 {
		return true
	}
	return false
}
func (b *bearish) isBeltHold(ks []*MK) bool {
	if len(ks) < 11 {
		return false
	}
	O, C, L, H := ks[0].O, ks[0].C, ks[0].L, ks[0].H
	C1 := ks[1].C
	C2 := ks[2].C
	H1, L1 := ks[1].H, ks[1].L
	C3 := ks[3].C
	AVGH10, AVGL10, MAXO10 := float64(0), float64(0), float64(0)
	for i := 0; i < 10; i += 1 {
		n := float64(i)
		AVGH10 = AVGH10*n/(n+1) + ks[i].H/(n+1)
		AVGL10 = AVGL10*n/(n+1) + ks[i].L/(n+1)
		MAXO10 = math.Max(MAXO10, ks[i].O)
	}
	eq1 := O == MAXO10
	eq2 := O > H1
	eq3 := O-C >= .7*(H-L)
	eq4 := H-L >= 1.2*(AVGH10-AVGL10)
	eq5 := H-O <= .01*(H-L)
	eq6 := C >= H1-.5*(H1-L1)
	eq7 := H1 > L1
	eq8 := H > L
	eq9 := C1 > C2
	eq10 := C2 < C3
	logCountEq(enums.BeltHold, eq1, eq2, eq3, eq4, eq5, eq6, eq7, eq8, eq9, eq10)
	if eq1 && eq2 && eq3 && eq4 && eq5 && eq6 && eq7 && eq8 && eq9 && eq10 {
		return true
	}
	return false
}
func (b *bearish) isBreakaway(ks []*MK) bool {
	if len(ks) < 5 {
		return false
	}
	O, C, L, H := ks[0].O, ks[0].C, ks[0].L, ks[0].H
	C1 := ks[1].C
	C2 := ks[2].C
	O3, C3 := ks[3].O, ks[3].C
	O4, C4 := ks[4].O, ks[4].C
	L3 := ks[3].L
	H4, L4 := ks[4].H, ks[4].L

	eq1 := ABS(C4-O4) > .5*(H4-L4)
	eq2 := C4 > O4
	eq3 := C3 > O3
	eq4 := L3 > H4
	eq5 := C2 > C3
	eq6 := C1 > C2
	eq7 := C < O
	eq8 := L < H4
	eq9 := H > L3
	logCountEq(enums.Breakaway, eq1, eq2, eq3, eq4, eq5, eq6, eq7, eq8, eq9)
	if eq1 && eq2 && eq3 && eq4 && eq5 && eq6 && eq7 && eq8 && eq9 {
		return true
	}
	return false
}
func (b *bearish) isDarkCloudCover(ks []*MK) bool {
	if len(ks) < 11 {
		return false
	}
	O, C := ks[0].O, ks[0].C
	O1, C1 := ks[1].O, ks[1].C
	H1, L1 := ks[1].H, ks[1].L
	AVGH10_1, AVGL10_1, MAXO10_1 := float64(0), float64(0), float64(0)
	for i := 1; i < 11; i += 1 {
		n := float64(i)
		AVGH10_1 = AVGH10_1*n/(n+1) + ks[i].H/(n+1)
		AVGL10_1 = AVGL10_1*n/(n+1) + ks[i].L/(n+1)
		MAXO10_1 = math.Max(MAXO10_1, ks[i].O)
	}

	eq1 := C1-O1 >= .7*(H1-L1)
	eq2 := H1-L1 >= AVGH10_1-AVGL10_1
	eq3 := O > C1
	eq4 := C < C1-.5*(C1-O1)
	eq5 := C > O1
	logCountEq(enums.DarkCloudCover, eq1, eq2, eq3, eq4, eq5)
	if eq1 && eq2 && eq3 && eq4 && eq5 {
		return true
	}
	return false
}

func (b *bearish) isDeliberation(ks []*MK) bool {
	if len(ks) < 11 {
		return false
	}
	O, C := ks[0].O, ks[0].C
	O1, C1 := ks[1].O, ks[1].C
	O2, C2 := ks[2].O, ks[2].C
	H, L := ks[0].H, ks[0].L
	H1, L1 := ks[1].H, ks[1].L
	H2, L2 := ks[2].H, ks[2].L

	eq1 := ABS(C2-O2) > .5*(H2-L2)
	eq2 := ABS(C1-O1) > .5*(H1-L1)
	eq3 := C1 > C2
	eq4 := C2 > O2
	eq5 := C1 > O1
	eq6 := O > H1
	eq7 := (C+O)/2-L > .4*(H-L)
	eq8 := (C+O)/2-L < .6*(H-L)
	eq9 := ABS(C-O) < .6*(H-L)
	logCountEq(enums.Deliberation, eq1, eq2, eq3, eq4, eq5, eq6, eq7, eq8, eq9)
	if eq1 && eq2 && eq3 && eq4 && eq5 && eq6 && eq7 && eq8 && eq9 {
		return true
	}
	return false
}
func (b *bearish) isDownsideGapThreeMethods(ks []*MK) bool {
	if len(ks) < 11 {
		return false
	}
	O, C := ks[0].O, ks[0].C
	O1, C1 := ks[1].O, ks[1].C
	O2, C2 := ks[2].O, ks[2].C
	H, L := ks[0].H, ks[0].L
	H1, L1 := ks[1].H, ks[1].L
	H2, L2 := ks[2].H, ks[2].L

	eq1 := ABS(C2-O2) > .5*(H2-L2)
	eq2 := ABS(C1-O1) > .5*(H1-L1)
	eq3 := C2 < O2
	eq4 := C1 < O1
	eq5 := H1 < L2
	eq6 := L < H1
	eq7 := H > L2
	eq8 := C > O
	logCountEq(enums.DownsideGapThreeMethods, eq1, eq2, eq3, eq4, eq5, eq6, eq7, eq8)
	if eq1 && eq2 && eq3 && eq4 && eq5 && eq6 && eq7 && eq8 {
		return true
	}
	return false
}
func (b *bearish) isDownsideTasukiGap(ks []*MK) bool {
	if len(ks) < 11 {
		return false
	}
	O, C := ks[0].O, ks[0].C
	O1, C1 := ks[1].O, ks[1].C
	O2, C2 := ks[2].O, ks[2].C
	H1 := ks[1].H
	L2 := ks[2].L

	eq1 := C2 < O2
	eq2 := C1 < O1
	eq3 := H1 < L2
	eq4 := O > C1
	eq5 := O < O1
	eq6 := C > H1
	eq7 := C < L2
	logCountEq(enums.DownsideTasukiGap, eq1, eq2, eq3, eq4, eq5, eq6, eq7)
	if eq1 && eq2 && eq3 && eq4 && eq5 && eq6 && eq7 {
		return true
	}
	return false
}
func (b *bearish) isDojiStar(ks []*MK) bool {
	if len(ks) < 11 {
		return false
	}
	O, C, L, H := ks[0].O, ks[0].C, ks[0].L, ks[0].H
	O1, C1, L1 := ks[1].O, ks[1].C, ks[1].L
	H1 := ks[1].H
	AVGH21, AVGL21 := float64(0), float64(0)
	for i := 0; i < 21; i += 1 {
		n := float64(i)
		AVGH21 = AVGH21*n/(n+1) + ks[i].H/(n+1)
		AVGL21 = AVGL21*n/(n+1) + ks[i].L/(n+1)
	}
	eq1 := ABS(C1-O1) > .5*(H1-L1)
	eq2 := O > C1
	eq3 := ABS(C-O) < .05*(H-L)
	eq4 := H-L < .2*(AVGH21-AVGL21)
	logCountEq(enums.DojiStar, eq1, eq2, eq3, eq4)
	if eq1 && eq2 && eq3 && eq4 {
		return true
	}
	return false
}
func (b *bearish) isDojiGravestone(ks []*MK) bool {
	if len(ks) < 11 {
		return false
	}
	O, C, L, H := ks[0].O, ks[0].C, ks[0].L, ks[0].H
	AVGH10, AVGL10 := float64(0), float64(0)
	MAXH10 := float64(0)
	for i := 0; i < 10; i += 1 {
		n := float64(i)
		AVGH10 = AVGH10*n/(n+1) + ks[i].H/(n+1)
		AVGL10 = AVGL10*n/(n+1) + ks[i].L/(n+1)
		MAXH10 = math.Max(MAXH10, ks[i].H)
	}
	eq1 := ABS(O-C) <= .01*(H-L)
	eq2 := (H - C) >= .95*(H-L)
	eq3 := H > L
	eq4 := H == MAXH10
	eq5 := (H - L) >= (AVGH10 - AVGL10)
	logCountEq(enums.DojiGravestone, eq1, eq2, eq3, eq4, eq5)
	if eq1 && eq2 && eq3 && eq4 && eq5 {
		return true
	}
	return false
}
