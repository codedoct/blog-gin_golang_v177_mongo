package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"

	"blog-gin_golang_v177_mongo/domain/auth/model"
	"blog-gin_golang_v177_mongo/domain/auth/repository"
	"blog-gin_golang_v177_mongo/lib/constant"
	"blog-gin_golang_v177_mongo/lib/encrypt"
)

type AuthServiceInterface interface {
	SignIn(reqBody model.ReqBody) (*model.ResBody, int, error)
	CheckAuth(bearerToken string) (bson.M, error)
	Logout(email string) (int, error)
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

func (s *authService) CheckAuth(bearerToken string) (bson.M, error) {
	tokenRaw, claims, err := encrypt.Parse(bearerToken)
	if err != nil {

		return nil, err
	}

	email := claims["email"].(string)
	user, _, err := s.Repository.FindOneByEmail(email)
	if err != nil {
		return nil, err
	}

	if user["token"] != tokenRaw {
		return nil, errors.New(constant.UserHasSignedOut)
	}
	return user, nil
}

func (s *authService) Logout(email string) (int, error) {
	if err := s.Repository.DeleteToken(email); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
