package main

import (
	"github.com/adikm/golang-bloggers/api"
	"github.com/adikm/golang-bloggers/db"
)

func main() {
	db.Connect()
	api.InitServer()

}
