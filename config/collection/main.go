package collection

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"blog-gin_golang_v177_mongo/app/controllers/auth"
	"blog-gin_golang_v177_mongo/config/middleware"
)

func MainRouter(db *mongo.Client, main *gin.RouterGroup) {
	authCtrl := auth.AuthController(db)
	auth := main.Group("auth")
	{
		auth.POST("/login", authCtrl.Login)
		auth.POST("/logout", middleware.Auth(db), authCtrl.Logout)
	}
}
