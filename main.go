package main

import (
	mydb "github.com/klymenok/go-playground/db"
	"github.com/klymenok/go-playground/router"
	"net/http"
)

func main() {
	mydb.Init()
	err := http.ListenAndServe(":3333", router.Init())
	if err != nil {
		return
	}
}
