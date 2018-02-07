package db

import (
	"database/sql"
	"practice/go/encrypt/common"

	_ "github.com/go-sql-driver/mysql"
)

var MySQL *sql.DB

func InitMysql(config *common.MySQLConfig) {
	pool, err := sql.Open("mysql", config.DSN)
	if err != nil {
		panic(err)
	}

	if config.MaxIdle > 0 {
		pool.SetMaxIdleConns(config.MaxIdle)
	}
	if config.MaxOpen > 0 {
		pool.SetMaxOpenConns(config.MaxOpen)
	}

	MySQL = pool
}
