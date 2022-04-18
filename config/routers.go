package config

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"blog-gin_golang_v177_mongo/app/controllers/root"
	"blog-gin_golang_v177_mongo/config/collection"
)

func Route(db *mongo.Client) {
	router := gin.Default()
	corsConfig(router)

	router.GET("/", root.Index)
	main := router.Group("v1")
	collection.MainRouter(db, main)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
