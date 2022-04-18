package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	ID                uint64
	Email             string
	EncryptedPassword string
	Token             string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Claims struct {
	Email string `json:"email"`
	*jwt.StandardClaims
}
