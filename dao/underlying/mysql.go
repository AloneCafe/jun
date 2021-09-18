package underlying

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"jun/utils/conf"
	"log"
)

var (
	db *sqlx.DB
)

func GetDB() *sqlx.DB {
	return db
}

func init() {
	database, err := sqlx.Open("mysql", conf.GetGlobalConfig().Db.DSN)
	if err != nil {
		log.Println("Failed to initialize MySQL driver", err)
		return
	}
	db = database
	if cnt := conf.GetGlobalConfig().Db.MaxOpenConn; cnt != 0 {
		db.SetMaxOpenConns(cnt)
	}
	if cnt := conf.GetGlobalConfig().Db.MaxIdleConn; cnt != 0 {
		db.SetMaxIdleConns(cnt)
	}

	//defer db.Close()
}
