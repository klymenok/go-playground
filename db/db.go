package mydb

import (
	"database/sql"
	"errors"
	"fmt"
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

func createDatabase(dbName string) {
	db := connect(dbName)
	defer db.Close()

	// Create DB Tables

	sqlUser := `create table user 
				(id integer not null primary key autoincrement, 
				first_name text, 
				last_name text);`
	sqlTask := `create table task 
				(id integer not null primary key autoincrement, 
				title text, description text, 
				created_by integer not null, 
				assignee integer not null, 
				completed boolean,
				FOREIGN KEY(created_by) REFERENCES user(id),
				FOREIGN KEY(assignee) REFERENCES user(id))
				`
	sqlComment := `create table comment 
				(id integer not null primary key autoincrement, 
				text text, 
				created_by integer, 
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY(created_by) REFERENCES user(id))
				`
	_, err := db.Exec(sqlUser)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(sqlTask)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(sqlComment)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database created")
}

func Init() {
	godotenv.Load(".env")
	if _, err := os.Stat(fmt.Sprintf("./%s", os.Getenv("DB_NAME"))); errors.Is(err, os.ErrNotExist) {
		createDatabase(os.Getenv("DB_NAME"))
	}

	db := connect(os.Getenv("DB_NAME"))
	db.Ping()
}

func Connection() {

}
