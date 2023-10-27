package app

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"restful-api/helper"
	"time"
)

func NewDB() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DATABASE")

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/"+dbName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
