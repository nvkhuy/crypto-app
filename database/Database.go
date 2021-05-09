package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
)

const DBLog = "[DB]"

type DBConfig struct {
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Host   string `json:"host"`
	Port   string `json:"port"`
	DBName string `json:"dbname"`
}

type Database struct {
	Config  *DBConfig
	Session *sql.DB
}

func (db *Database) fetchConfig(url string) {
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
	json.Unmarshal(byteValue, &db.Config)
}
func (db *Database) Open(configUrl string) {
	if configUrl != "" {
		db.fetchConfig(configUrl)
	}
	cf := db.Config
	url := cf.User + ":" + cf.Pass + "@tcp(" + cf.Host + ":" + cf.Port + ")/" + cf.DBName + "?parseTime=true"
	log.Println(DBLog+"url connections:", url)
	dtb, err := sql.Open("mysql", url)
	if err != nil {
		log.Println(DBLog + err.Error())
		return
	}
	if err = dtb.Ping(); err != nil {
		log.Println(DBLog + "cannot ping database")
		return
	} else {
		log.Println(DBLog + "ping database successfully")
	}
	if dtb != nil {
		log.Println(DBLog + cf.DBName + " connected")
		db.Session = dtb
	} else {
		err = errors.New(DBLog + "cannot connect database")
	}
}
func (db *Database) Close() {
	if db.Session != nil {
		log.Println(DBLog + db.Config.DBName + " closed")
		db.Session.Close()
	}
}
