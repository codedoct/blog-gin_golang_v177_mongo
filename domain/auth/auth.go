package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"

	"blog-gin_golang_v177_mongo/domain/auth/model"
	"blog-gin_golang_v177_mongo/domain/auth/repository"
	"blog-gin_golang_v177_mongo/lib/constant"
	"blog-gin_golang_v177_mongo/lib/encrypt"
)

type AuthServiceInterface interface {
	SignIn(reqBody model.ReqBody) (*model.ResBody, int, error)
}

type authService struct {
	Repository repository.AuthRepositoryInterface
}

func AuthService(repository repository.AuthRepositoryInterface) AuthServiceInterface {
	return &authService{
		Repository: repository,
	}
}

func (s *authService) SignIn(reqBody model.ReqBody) (*model.ResBody, int, error) {
	user, status, err := s.Repository.FindOneByEmail(reqBody.Email)
	if err != nil {
		return nil, status, err
	}

	encryptedPassword := user["encrypted_password"].(string)
	if err = encrypt.CompareHashAndPassword(&encryptedPassword, &reqBody.Password); err != nil {
		return nil, http.StatusBadRequest, errors.New(constant.PasswordIsIncorrect)
	}
	claims := model.Claims{
		Email: user["email"].(string),
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		},
	}
	ss, err := encrypt.NewWithClaims(&claims)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err = s.Repository.UpdateToken(user["email"].(string), ss); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	var resBody model.ResBody
	resBody.Token = ss
	return &resBody, http.StatusOK, nil
}
