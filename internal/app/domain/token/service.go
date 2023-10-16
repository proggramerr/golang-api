package token

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TokenRepository interface {
	CreateToken(username string, accessToken string, refreshToken string, expire time.Time) error
	FindTokenByUsername(username string) (Token, error)
	DeleteToken(tokenID primitive.ObjectID) error
}

type TokenService struct {
	tokenRepo TokenRepository
	secretKey []byte
}

func NewTokenService(tokenRepo TokenRepository, secretKey []byte) *TokenService {
	return &TokenService{tokenRepo: tokenRepo, secretKey: secretKey}
}

func (service *TokenService) GenerateToken(username string) (string, string, error) {
	// Генерация Access Token
	accessToken, err := service.generateJWT(username, time.Hour*2) // Настройте срок действия по вашим потребностям

	if err != nil {
		return "", "", err
	}

	// Генерация Refresh Token
	refreshToken, err := service.generateJWT(username, time.Hour*24*7) // Настройте срок действия по вашим потребностям

	if err != nil {
		return "", "", err
	}

	// Сохранение Refresh Token в базе данных
	err = service.tokenRepo.CreateToken(username, accessToken, refreshToken, time.Now().Add(time.Hour*2))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (service *TokenService) generateJWT(username string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(service.secretKey)
}
