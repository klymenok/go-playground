package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	url string
}

func New() *DB {
	godotenv.Load(".env")
	dbName := os.Getenv("DB_NAME")
	return &DB{dbName}
}

func connect(dbName string) *sql.DB {
	db, err := sql.Open("sqlite3", "./"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (db *DB) Exec(sql string) (sql.Result, error) {
	conn := connect(db.url)
	defer conn.Close()

	return conn.Exec(sql)
}

func (db *DB) QueryRow(sql string) *sql.Row  {
	conn := connect(db.url)
	defer conn.Close()

	return conn.QueryRow(sql)
}
