package captain

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/bwmarrin/discordgo"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/enums"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/global"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	ctx        = context.Background()
	LIMIT_DOWN = float64(30)
	LIMIT_UP   = float64(70)
)

func setLimitDown(limit float64, s *discordgo.Session, m *discordgo.MessageCreate) {
	LIMIT_DOWN = limit
	s.ChannelMessageSend(m.ChannelID, enums.Captain+enums.FingerRightPoint+fmt.Sprintf(" set limit down to %.f ", LIMIT_DOWN)+enums.Verified)
}

func setLimitUP(limit float64, s *discordgo.Session, m *discordgo.MessageCreate) {
	LIMIT_UP = limit
	s.ChannelMessageSend(m.ChannelID, enums.Captain+enums.FingerRightPoint+fmt.Sprintf(" set limit up to %.f ", LIMIT_UP)+enums.Verified)
}

func sendSymbolInfo(symbol string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if symbol == "" {
		return
	}
	symbol = strings.ToUpper(symbol)
	prices, err := global.Api.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, enums.Captain+"captain cannot get symbol info "+enums.CrossMarkRed)
	} else if len(prices) == 0 {
		log.Println("empty prices listing")
		return
	} else {
		s.ChannelMessageSend(m.ChannelID, enums.Stock+" "+prices[0].Symbol+" current price "+enums.FingerRightPoint+" "+prices[0].Price)
	}
}
func autoBook(n int, symbol string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if symbol == "" {
		return
	}
	symbol = strings.ToUpper(symbol)
	s.ChannelMessageSend(m.ChannelID, enums.Captain+"i'm seeking for the symbol "+symbol+" potentials "+enums.OkHand)
	serviceCycle.symbol = symbol
	serviceCycle.ticker = time.NewTicker(3600 * time.Second)
	serviceCycle.isVisualize = time.NewTicker(3600 * time.Second)
	serviceCycle.isTrade = time.NewTicker(10 * time.Second)
	serviceCycle.stop = make(chan bool)
	// TEST RSI
	serviceCycle.base = float64(1000)
	serviceCycle.quote = float64(0)
	go traceSymbol(n, s, m)
}
func stopBook() {
	serviceCycle.stop <- true
}
func traceSymbol(n int, s *discordgo.Session, m *discordgo.MessageCreate) {
	for {
		klines, err := global.Api.NewKlinesService().Symbol(serviceCycle.symbol).Interval("1m").Limit(n).Do(context.Background())
		select {
		case <-serviceCycle.ticker.C:
			if err != nil {
				log.Println(enums.Captain + "something wrong in tracing symbol process" + enums.Rejected)
			} else {
				if caseName := bull.CaseMatching(klines); caseName != "" {
					s.ChannelMessageSend(m.ChannelID, enums.Stock+enums.Captain+"good to buy "+strings.Repeat(enums.RocketEmoji, 3)+"\n"+
						enums.Stock+"CASE#"+enums.FingerRightPoint+caseName)
				}
			}
		case <-serviceCycle.isVisualize.C:
			s.ChannelMessageSend(m.ChannelID, enums.Stock+enums.Captain+enums.FingerRightPoint+visualizeTrend(klines))
			//s.ChannelMessageSend(m.ChannelID, enums.Stock+enums.Captain+enums.FingerRightPoint+showRSI(15, calcEMA))
		case <-serviceCycle.isTrade.C:
			kLines14, err := global.Api.NewKlinesService().Symbol(serviceCycle.symbol).Interval("1m").Limit(15).Do(context.Background())
			if err != nil {
				return
			}
			price, _ := strconv.ParseFloat(kLines14[14].Close, 64)

			rsi := calcEMA(kLines14)
			rsiNum := fmt.Sprintf("%d", int(rsi))

			global.Rdb.HIncrBy(ctx, "counter", rsiNum, 1)

			if rsi <= LIMIT_DOWN && serviceCycle.base > 0 {
				s.ChannelMessageSend(m.ChannelID, enums.Stock+enums.Captain+"GOOD TO BUY "+strings.Repeat(enums.RocketEmoji, 3)+"\n"+
					enums.Stock+"RSI#"+enums.FingerRightPoint+fmt.Sprintf("[%.f]", rsi))
				serviceCycle.quote = serviceCycle.base / price
				serviceCycle.base = 0
			}
			if rsi >= LIMIT_UP && serviceCycle.quote > 0 {
				s.ChannelMessageSend(m.ChannelID, enums.Stock+enums.Captain+"GOOD TO SELL "+strings.Repeat(enums.RocketEmoji, 3)+"\n"+
					enums.Stock+"RSI#"+enums.FingerRightPoint+fmt.Sprintf("[%.f]", rsi))
				serviceCycle.base = serviceCycle.quote * price
				serviceCycle.quote = 0
			}
			s.ChannelMessageSend(m.ChannelID, enums.Stock+enums.Captain+enums.FingerRightPoint+
				fmt.Sprintf("Base: %.2f - Quote: %.2f - RSI: %.2f - Price: %.2f", serviceCycle.base, serviceCycle.quote, rsi, price))
			s.ChannelMessageSend(m.ChannelID, enums.Stock+strings.Repeat(enums.RocketEmoji, 3)+
				fmt.Sprintf("Current Profit: %.2f", profit(serviceCycle.base, serviceCycle.quote, price)))
		case <-serviceCycle.stop:
			s.ChannelMessageSend(m.ChannelID, enums.Captain+"i'm stop looking for symbol "+serviceCycle.symbol+" potential"+enums.OkHand)
			return
		}
	}
}
func profit(base, quote, price float64) float64 {
	if base != 0 {
		return base - 1000
	}
	return quote*price - 1000
}
func visualizeTrend(klines []*binance.Kline) string {
	trendEmoji := ""
	for _, k := range klines {
		_, _, cc, co := toFloat(k)
		if cc < co {
			trendEmoji += enums.CrossMarkRed
		} else {
			trendEmoji += enums.Verified
		}
	}
	return trendEmoji
}

func showRSI(n int, calcRSI rsiFormula) (res string) {
	kLines, err := global.Api.NewKlinesService().Symbol(serviceCycle.symbol).Interval("1m").Limit(2*n + 1).Do(context.Background())
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		res += fmt.Sprintf("[%.f]-", calcRSI(kLines[i:i+n]))
	}
	return
}
func calcEMA(klines []*binance.Kline) (rsi float64) {
	var avgU, avgD, prv, n, u, d, alpha float64
	n = float64(len(klines)) - 1
	alpha = 1 / n
	_, _, cc, _ := toFloat(klines[0])
	prv = cc

	for _, k := range klines[1:] {
		_, _, cc, _ = toFloat(k)
		u, d = 0, 0
		if prv < cc {
			u = cc - prv
		}
		if prv > cc {
			d = prv - cc
		}
		avgU = alpha*u + (1-alpha)*avgU
		avgD = alpha*d + (1-alpha)*avgD
		prv = cc
	}

	rsi = 100 - 100/(1+(avgU/avgD))
	return
}
func calcSMA(klines []*binance.Kline) (rsi float64) {
	var avgU, avgD, prv, n float64
	n = float64(len(klines)) - 1

	for _, k := range klines[1:] {
		_, _, cc, _ := toFloat(k)
		if prv <= cc {
			avgU += cc
		} else {
			avgD += cc
		}
		prv = cc
	}
	avgU = avgU / n
	avgD = avgD / n
	rsi = 100 - 100/(1+(avgU/avgD))
	return
}
func toFloat(k *binance.Kline) (float64, float64, float64, float64) {
	h, _ := strconv.ParseFloat(k.High, 64)
	l, _ := strconv.ParseFloat(k.Low, 64)
	cc, _ := strconv.ParseFloat(k.Close, 64)
	co, _ := strconv.ParseFloat(k.Open, 64)
	return h, l, cc, co
}
