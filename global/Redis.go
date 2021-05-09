package global

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	rd "github.com/khanhhuynguyenvu/crypto-trading-heroku/enums/redis"
	"io/ioutil"
	"log"
	"os"
)

const REDIS = "[REDIS]"

var (
	Rdb *redis.Client
	ctx = context.Background()
)

type redisConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func SetRedis() {
	cfg := fetchConfig(rd.Live)
	url := cfg.Host + ":" + cfg.Port
	Rdb = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})
	pong := Rdb.Ping(ctx)
	if pong == nil {
		logRedis("cannot ping redis successfully")
		return
	}
	logRedis(fmt.Sprintf("connect redis host:%s port:%s successfully", cfg.Host, cfg.Port))
}
func fetchConfig(url string) (config *redisConfig) {
	if url == "" {
		log.Println()
		return
	}
	jsonFile, err := os.Open(url)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)
	if config == nil {
		logRedis("cannot fetch redis config")
		return
	}
	return
}
func logRedis(msg string) {
	log.Println(REDIS + msg)
}
