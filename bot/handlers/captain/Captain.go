package captain

import (
	"github.com/bwmarrin/discordgo"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/enums"
	"log"
	"strconv"
	"strings"
)

func CaptainHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.State.User.ID == m.Author.ID {
		return
	}
	if m.Content == "hey" {
		s.ChannelMessageSend(m.ChannelID, enums.Captain+"captain here")
	}
	requestHandler(s, m)
}
func requestHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	requests := strings.Split(m.Content, " ")
	if len(requests) == 0 {
		return
	}
	if strings.Contains(requests[0], "pr") { // price
		if len(requests) == 1 {
			return
		}
		sendSymbolInfo(requests[1], s, m)
	}
	if strings.Contains(requests[0], "or") { // ord
		if len(requests) == 1 {
			goto STOP
		}
		n := 15
		if len(requests) > 2 {
			_n, err := strconv.Atoi(requests[2]) // number of kline
			if err != nil {
				log.Println()
			}
			n = _n
		}
		autoBook(n, requests[1], s, m)
	}
STOP:
	if strings.Contains(requests[0], "st") {
		stopBook()
	}
	if strings.Contains(requests[0], "lm") {
		if len(requests) < 3 {
			goto SKIP
		}
		if strings.Contains(requests[1], "u") {
			limit, err := strconv.ParseFloat(requests[2], 32)
			if err != nil {
				goto SKIP
			}
			setLimitUP(limit, s, m)
			log.Printf("set limit up to %.f", limit)
		}
		if strings.Contains(requests[1], "d") {
			limit, err := strconv.ParseFloat(requests[2], 32)
			if err != nil {
				goto SKIP
			}
			setLimitDown(limit, s, m)
			log.Printf("set limit down to %.f", limit)
		}
	}
SKIP:
}
