package main

import (
	_ "github.com/joho/godotenv/autoload"
	"log"

	"blog-gin_golang_v177_mongo/config"
	"blog-gin_golang_v177_mongo/db"
)

func main() {
	mongo := db.Mongodb()
	defer func() {
		if err := db.Client.Disconnect(db.Ctx); err != nil {
			log.Fatal(err)
			panic(err)
		}
	}()
	config.Route(mongo)
}
