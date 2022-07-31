package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"taurus-backend/constant"
)

var (
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
	db         *sql.DB
)

func init() {
	if constant.LocalStage == os.Getenv(constant.Stage) {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("get local env err:", err)
		}
	}
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	dbHandler, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatal("dbHandler handler open fail", err)
	}
	//设置数据库最大连接数
	dbHandler.SetConnMaxLifetime(100)
	dbHandler.SetMaxIdleConns(10)
	db = dbHandler
}

func GetDB() *sql.DB {
	return db
}
