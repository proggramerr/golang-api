package models

import (
	"time"
)

type Token struct {
	Username     string    `json:"userName" bson:"userName"`
	AccessToken  string    `json:"accessToken" bson:"accessToken"`
	RefreshToken string    `json:"refreshToken" bson:"refreshToken"`
	Expire       time.Time `json:"expire" bson:"expire"`
}
