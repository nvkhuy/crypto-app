package global

import (
	"errors"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/bot"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/enums"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/models"
	"log"
	"net/http"
)

var Captain *bot.Supervisor

func SetCaptain() (int, error, string) {
	executor := Db.Session
	if executor == nil {
		err := errors.New("empty sql executor")
		return http.StatusInternalServerError, err, ""
	}
	captainInfo, err := models.Bots(models.BotWhere.ID.EQ(1)).One(executor)
	if err != nil {
		return http.StatusInternalServerError, err, ""
	}
	if captainInfo == nil {
		return http.StatusInternalServerError, errors.New(bot.SupervisorLog + "empty captain info"), ""
	}
	if Captain != nil {
		Captain.Session.ChannelMessageSend(enums.GeneralChannelId, enums.CaptainAlreadyHere)
		return http.StatusConflict, nil, enums.CaptainAlreadyHere
	}
	Captain = &bot.Supervisor{
		BotDiscord: &bot.BotDiscord{
			Bot: captainInfo,
		},
	}
	Captain.Open()
	Captain.Session.ChannelMessageSend(enums.GeneralChannelId, enums.CaptainReady)
	log.Println(bot.SupervisorLog + enums.CaptainReady)
	return http.StatusOK, nil, enums.CaptainReady
}
