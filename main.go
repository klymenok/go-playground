package main

import (
	"fmt"
	mydb "github.com/klymenok/go-playground/db"
	_ "github.com/klymenok/go-playground/docs"
	"github.com/klymenok/go-playground/router"
	"net/http"
)

// colors for console output
var ResetColor = "\033[0m"
var GreenColor = "\033[32m"

// @title           Task app API
// @version         1.0
// @description     This is a documentation for Task app
// @termsOfService  To be implemented

// @contact.name   Oleksii Klymenok
// @contact.url    https://github.com/klymenok
// @contact.email  klymenok.a@gmail.com

// @license.name  MIT
// @license.url   some url here

// @host      localhost:3333
// @BasePath  /

// @securityDefinitions.basic  NoAuth

func main() {
	fmt.Println("Database initialization...")
	mydb.Init()
	fmt.Println(GreenColor + "Database initialized" + ResetColor)
	fmt.Println("Starting web server on port 3333")
	fmt.Println(GreenColor + "Server started and ready for connections" + ResetColor)

	err := http.ListenAndServe(":3333", router.Init())
	if err != nil {
		return
	}
}
