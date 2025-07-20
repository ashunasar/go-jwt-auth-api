package utils

import (
	"time"

	"github.com/ashunasar/go-jwt-auth-api/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func SignAccessToken(userId uuid.UUID) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(1 * time.Hour).Unix(), // Token expires in 1 hours
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.ConfigData.AccessTokenSecret))

	return signedToken, err

}

func SignRefreshToken(userId uuid.UUID) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(60 * 24 * time.Hour).Unix(), // Token expires in 60 days
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.ConfigData.RefreshTokenSecret))

	return signedToken, err

}
