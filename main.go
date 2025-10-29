package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "test"
)

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Print("Successfully connected.")

	var count = 0
	for {
		_, err := db.Exec("INSERT INTO messages(message) VALUES('Hello, World!')")
		if err != nil {
			log.Printf("failed executing query %d: %v", count, err)
		}

		time.Sleep(10 * time.Second)
	}
}
