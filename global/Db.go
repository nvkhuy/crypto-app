package global

import (
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/database"
	"github.com/khanhhuynguyenvu/crypto-trading-heroku/enums/db"
)

var Db *database.Database

func SetDb() {
	Db = &database.Database{}
	Db.Open(db.Live) // fetch config then open
}
