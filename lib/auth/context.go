package auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthData struct {
	ID    primitive.ObjectID
	Email string
	Role  string
}

func GetAuthUserCtx(ctx *gin.Context) (*AuthData, error) {
	authDataStr := ctx.GetString("auth")

	var authData AuthData
	err := json.Unmarshal([]byte(authDataStr), &authData)
	if err != nil {
		return nil, err
	}

	return &authData, nil
}
