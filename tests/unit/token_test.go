package unit

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"rest_api/internal/app/adapters/db/mongo/repo"
	"rest_api/internal/app/domain/token"
	mongoClient "rest_api/pkg/client/mongo"
	"testing"
)

func setupDatabase() *mongo.Database {
	engine, err := mongoClient.GetMongoEngine()
	if err != nil {
		panic(err)
	}
	return engine.Database("TestUsers")
}

func TestGenerateToken(t *testing.T) {
	db := setupDatabase()
	repo := repo.NewTokenRepository(db.Collection("Users"))
	service := token.NewTokenService(repo, []byte("yaebyalibaby"))
	fmt.Println(service.GenerateToken("testusesrname"))
	//serivce := TokenService()
}
