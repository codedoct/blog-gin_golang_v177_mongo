package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"

	"blog-gin_golang_v177_mongo/lib/constant"
)

type AuthRepositoryInterface interface {
	FindOneByEmail(email string) (bson.M, int, error)
	UpdateToken(email, ss string) error
}

type authRepository struct {
	db *mongo.Collection
}

func AuthRepository(db *mongo.Collection) AuthRepositoryInterface {
	return &authRepository{db: db}
}

func (r *authRepository) FindOneByEmail(email string) (bson.M, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user bson.M
	err := r.db.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusBadRequest, errors.New(constant.UserNotFound)
		}
		return nil, http.StatusInternalServerError, err
	}
	return user, http.StatusOK, nil
}

func (r *authRepository) UpdateToken(email, ss string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.db.UpdateOne(ctx, bson.M{"email": email},
		bson.D{
			{"$set", bson.D{{"token", ss}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
