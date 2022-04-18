package model

type ReqBody struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required,gte=6"`
}
