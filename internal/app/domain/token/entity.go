package token

import (
	"time"
)

type Token struct {
	Username     string    `json:"username"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	Expire       time.Time `json:"expire"`
}
