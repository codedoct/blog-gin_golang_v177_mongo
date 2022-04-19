package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"

	"blog-gin_golang_v177_mongo/domain/auth"
	"blog-gin_golang_v177_mongo/domain/auth/repository"
	authLib "blog-gin_golang_v177_mongo/lib/auth"
	"blog-gin_golang_v177_mongo/lib/env"
	"blog-gin_golang_v177_mongo/lib/response"
)

func Auth(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authService := auth.AuthService(repository.AuthRepository(db.Database(env.String("DATABASE_NAME", "base_code")).Collection("users")))
		user, err := authService.CheckAuth(c.Request.Header.Get("Authorization"))
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		userStr, err := json.Marshal(&authLib.AuthData{
			ID:    user["_id"].(primitive.ObjectID),
			Email: user["email"].(string),
			Role:  user["role"].(string),
		})
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Set("auth", string(userStr))
	}
}
