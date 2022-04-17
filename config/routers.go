package config

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"blog-gin_golang_v177_mongo/app/controllers/root"
)

func Route(db *mongo.Client) {
	router := gin.Default()
	corsConfig(router)

	router.GET("/", root.Index)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
