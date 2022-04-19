package main

import (
	mydb "github.com/klymenok/go-playground/db"
	_ "github.com/klymenok/go-playground/docs"
	"github.com/klymenok/go-playground/router"
	"net/http"
)

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
	mydb.Init()
	err := http.ListenAndServe(":3333", router.Init())
	if err != nil {
		return
	}
}
