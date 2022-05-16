package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Conn  *sql.DB
	url string
}

func New() *sql.DB {
	godotenv.Load(".env")

	db := DB{}
	db.url = os.Getenv("DB_NAME")
	db.Conn = connect(db.url)

	return db.Conn
}

func connect(dbName string) *sql.DB {
	db, err := sql.Open("sqlite3", "./"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
