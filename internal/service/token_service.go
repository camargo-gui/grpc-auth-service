package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	SecretKey string
}

func NewTokenService(secretKey string) *TokenService {
	return &TokenService{
		SecretKey: secretKey,
	}
}

func (service *TokenService) GenerateToken(userID uint32, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(expiration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(service.SecretKey))
	if err != nil {
		return "", errors.New("failed to sign the token")
	}

	return signedToken, nil
}

func (service *TokenService) ValidateToken(tokenString string) (uint32, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(service.SecretKey), nil
	})
	if err != nil {
		return 0, errors.New("failed to parse the token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	userID := uint32(claims["sub"].(float64))
	return userID, nil
}
