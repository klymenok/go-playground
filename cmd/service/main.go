package main

import (
	"fmt"
	"net/http"

	_ "github.com/klymenok/go-playground/docs"

	"github.com/klymenok/go-playground/internal/db"
	"github.com/klymenok/go-playground/internal/handlers"
	"github.com/klymenok/go-playground/internal/todo"
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
	fmt.Println("Starting web server on port 3333")
	fmt.Println(GreenColor + "Server started and ready for connections" + ResetColor)

	manager := todo.NewManager(db.New())

	err := http.ListenAndServe(":3333", handlers.Init(manager))
	if err != nil {
		return
	}
}
