package main

import (
	"github.com/golang-work/adventure/bootstrap"
	"github.com/golang-work/adventure/support"
)

func main() {
	db, _ := support.DB.DB()
	defer db.Close()

	bootstrap.Run()
}
