package cmd

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

func GetConnect(cfg *models.Config) *sqlx.DB {
	dbCfg := cfg.DataBase
	dataSourceName := dbCfg.Username + ":" + dbCfg.Password + "@tcp(" + dbCfg.Host + ":" + dbCfg.Port + ")/" + dbCfg.DBName

	conn, err := sqlx.Connect("mysql", dataSourceName)
	conn.SetConnMaxLifetime(time.Minute * 6)
	if err != nil {
		panic(err)
	}
	//defer func(conn *sqlx.DB) {
	//	err := conn.Close()
	//	if err != nil {
	//
	//	}
	//}(conn)

	return conn
}
