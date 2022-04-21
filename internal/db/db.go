package mydb

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func connect(dbName string) *sql.DB {
	db, err := sql.Open("sqlite3", "./"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Connection() *sql.DB {
	godotenv.Load(".env")
	return connect(os.Getenv("DB_NAME"))
}
