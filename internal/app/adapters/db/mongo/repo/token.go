package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rest_api/internal/app/adapters/db/mongo/models"
	"rest_api/internal/app/domain/token"
	"time"
)

func NewTokenRepository(collection *mongo.Collection) *TokenRepositoryImpl {
	return &TokenRepositoryImpl{
		collection: collection,
	}
}

type TokenRepositoryImpl struct {
	collection *mongo.Collection
}

func (repo *TokenRepositoryImpl) CreateToken(username string, accessToken string, refreshToken string, expire time.Time) error {
	tokenModel := models.Token{
		Username:     username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expire:       time.Now().Add(time.Hour * 2),
	}
	_, err := repo.collection.InsertOne(context.TODO(), tokenModel)
	return err
}

func (repo *TokenRepositoryImpl) FindTokenByUsername(username string) (token.Token, error) {
	var tokenModel token.Token
	err := repo.collection.FindOne(context.TODO(), bson.M{"userName": username}).Decode(&tokenModel)
	if err != nil {
		return tokenModel, err
	}
	return tokenModel, nil
}
func (repo *TokenRepositoryImpl) DeleteToken(tokenID primitive.ObjectID) error {
	_, err := repo.collection.DeleteOne(context.TODO(), bson.M{"_id": tokenID})
	return err
}
