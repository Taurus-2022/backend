package db

import (
	"database/sql"

	"log"
	"taurus-backend/constant"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func Init(e *constant.Env) (*sql.DB, error) {
	dbHandler, err := sql.Open("mysql", e.GetDBConnectString())
	if err != nil {
		log.Println("dbHandler handler open fail", err)
		return nil, err
	}
	dbHandler.SetConnMaxLifetime(100)
	dbHandler.SetMaxIdleConns(50)
	dbHandler.SetMaxOpenConns(100)
	db = dbHandler
	return db, nil
}

func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("db is nil")
	}
	return db
}
