package main

import (
	"flag"
	"log"
	"os"
	"rest/libraries/config"
	"rest/libraries/database"
	"rest/schema"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if _, ok := os.LookupEnv("APP_ENV"); !ok {
		config.Setup(".env")
	}

	db, err := database.OpenDB()
	if err != nil {
		log.Fatalf("error: connecting to db: %s", err)
	}
	defer db.Close()

	// Handle cli command
	flag.Parse()

	if flag.Arg(0) == "migrate" {
		if err := schema.Migrate(db); err != nil {
			log.Println("error applying migrations", err)
			os.Exit(1)
		}
		log.Println("Migrations complete")
		return
	} else if flag.Arg(0) == "seed" {
		if err := schema.Seed(db); err != nil {
			log.Println("error applying seed", err)
			os.Exit(1)
		}
		log.Println("Seed complete")
		return
	}
}
