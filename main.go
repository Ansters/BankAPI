package main

import (
	"bankapi/userapi"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("postgres://vffqgrsb:ArlFUNrd_FC0qIC-zRFhJf0FDEqa6CUS@elmer.db.elephantsql.com:5432/vffqgrsb"))
	if err != nil {
		log.Fatal(err)
	}

	userapi.StartServer(":"+os.Getenv("PORT"), db)
}
