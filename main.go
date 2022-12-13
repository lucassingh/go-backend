package main

import (
	"log"

	"github.com/lucassingh/go-backend/db"
	"github.com/lucassingh/go-backend/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Not connect to DB")
		return
	}

	handlers.HandlersCon()
}
