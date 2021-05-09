package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/bot"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/global"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const ginHandlerLog = "[GIN HANDLER]"

type GinHandler struct{}

func (g *GinHandler) Handling() {
	r := gin.Default()
	r.GET("/", g.hello)
	v1 := r.Group("/v1")
	{
		captain := v1.Group("/captain")
		{
			captain.GET("/start", g.captainStart)
		}
		botRouter := v1.Group("/bots/:ID")
		{
			botRouter.GET("/start", g.botStart)
			botRouter.GET("/stop", g.botStop)
		}
	}
	r.Run(g.getPort())
}

func (g *GinHandler) getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println(httpHandlerLog + "INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func (g *GinHandler) hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": ginHandlerLog + "start",
	})
}
func (g *GinHandler) captainStart(c *gin.Context) {
	code, err, successMsg := global.SetCaptain()
	if err != nil {
		c.JSON(code, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(code, gin.H{
		"message": successMsg,
	})
}
func (g *GinHandler) botStart(c *gin.Context) {
	var b *bot.BotDiscord
	ID := c.Param("ID")
	id, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot convert bot id",
		})
		return
	}
	b, err = global.Captain.GetBot(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	b.Open()
	b.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if s.State.User.ID == m.Author.ID {
			return
		}
		if strings.Contains(m.Content, "hey") {
			b.Session.ChannelMessageSend(m.ChannelID, "sup bruh? ")
		}
	})
	c.JSON(http.StatusOK, gin.H{
		"message": b.Name.String + " start",
	})
}
func (g *GinHandler) botStop(c *gin.Context) {
	var b *bot.BotDiscord
	ID := c.Param("ID")
	id, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot convert bot id",
		})
		return
	}
	b, err = global.Captain.GetBot(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	b.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": b.Name.String + " stop",
	})
}
