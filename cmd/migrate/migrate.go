package main

import (
	"github.com/seinyan/gorest/internal/database"
	"github.com/seinyan/gorest/internal/migrations"
	"log"
	"time"
)

func main() {

	db, err := database.NewDBConn()
	if err != nil {
		log.Fatal(err)
	}

	migrations.Migrate(db)

	time.Sleep(time.Second * 2)
}