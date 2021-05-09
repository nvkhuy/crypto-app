package bot

import (
	"github.com/friendsofgo/errors"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/models"
	"log"
	"strconv"
)

const SupervisorLog = "[Supervisor]"

type Supervisor struct {
	*BotDiscord
	Bots []*BotDiscord
}

func (s *Supervisor) AddBot(b *models.Bot) {
	if b.Token.String == "" {
		log.Println(SupervisorLog + "bot has empty token")
		return
	}
	chatBot := &BotDiscord{
		Bot: b,
	}
	s.Bots = append(s.Bots, chatBot)
}
func (s *Supervisor) GetBot(ID int) (*BotDiscord, error) {
	if len(s.Bots) == 0 {
		return nil, errors.New("captain bots is empty")
	}
	for _, b := range s.Bots {
		if b.ID == ID {
			return b, nil
		}
	}
	id := strconv.Itoa(ID)
	log.Println(SupervisorLog + "cannot get bot with id " + id)
	err := errors.New("cannot get bot with id " + id)
	return nil, err
}
