package unit

import (
	"rest_api/pkg/client/mongo"
	"testing"
)

func TestMongoConnection(t *testing.T) {
	engine, err := mongo.GetMongoEngine()
	if err != nil {
		panic(err)
	}
	err = engine.Ping(nil, nil)
	if err != nil {
		panic(err)
	}
}
