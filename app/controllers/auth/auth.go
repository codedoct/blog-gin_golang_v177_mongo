package auth

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"

	"blog-gin_golang_v177_mongo/domain/auth"
	"blog-gin_golang_v177_mongo/domain/auth/model"
	"blog-gin_golang_v177_mongo/domain/auth/repository"
	authLib "blog-gin_golang_v177_mongo/lib/auth"
	"blog-gin_golang_v177_mongo/lib/env"
	"blog-gin_golang_v177_mongo/lib/response"
)

type authController struct {
	service auth.AuthServiceInterface
}

func AuthController(db *mongo.Client) *authController {
	return &authController{service: auth.AuthService(repository.AuthRepository(db.Database(env.String("DATABASE_NAME", "base_code")).Collection("users")))}
}

func (c *authController) Login(ctx *gin.Context) {
	var req model.ReqBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	res, statusCode, err := c.service.SignIn(req)
	if err != nil {
		response.Error(ctx, statusCode, err.Error())
		return
	}
	response.Json(ctx, statusCode, res)
}

func (c *authController) Logout(ctx *gin.Context) {
	authUser, err := authLib.GetAuthUserCtx(ctx)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	status, err := c.service.Logout(authUser.Email)
	if err != nil {
		response.Error(ctx, status, err.Error())
		return
	}
	response.Json(ctx, status, nil)
}
