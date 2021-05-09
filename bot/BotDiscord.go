package bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/models"
	"log"
)

type BotDiscord struct {
	*models.Bot
	Session *discordgo.Session
}

var (
	SessionCreatedBefore        = errors.New(BotLog + "session Created Before")
	EmptyToken                  = errors.New(BotLog + "empty Bot Token")
	ErrorCreatingDiscordSession = errors.New(BotLog + "error Creating Discord Session")
)

const (
	BotLog = "[DISCORD BOT]"
)

func (b *BotDiscord) createSession() (err error) {
	if b.Session != nil {
		err = SessionCreatedBefore
		return
	}
	if b.Token.String == "" {
		err = EmptyToken
		return
	}
	dg, err := discordgo.New("Bot " + b.Token.String)
	if err != nil {
		log.Println(BotLog + " error " + err.Error())
		return
	}
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	b.Session = dg
	log.Println(BotLog + " session created")
	return nil
}
func (b *BotDiscord) Open() {
	var err error
	if b.Session == nil {
		err = b.createSession()
	} else {
		log.Println(BotLog + b.Name.String + "already opened")
		return
	}
	if err != nil {
		log.Fatal(ErrorCreatingDiscordSession)
		return
	}
	err = b.Session.Open()
	if err != nil {
		log.Println(BotLog+"bot with name"+b.Name.String+"has opening error connection,", err)
		return
	}
	log.Println(BotLog + "open discord successfully")
}
func (b *BotDiscord) Close() {
	if b.Session != nil {
		b.Session.Close()
		b.Session = nil
	}
}

func (b *BotDiscord) AddHandlerOnce(fh interface{}) {
	b.Session.AddHandler(fh)
}
func (b *BotDiscord) GetSession() interface{} {
	return b.Session
}
