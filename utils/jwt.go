package utils

import (
	"fmt"
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

func VerifyAccessToken(tokenString string) (uuid.UUID, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpted signing method: %v", token.Header["alg"])

		}
		return []byte(config.ConfigData.AccessTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return uuid.Nil, fmt.Errorf("could not parse claims")
	}

	userIdStr, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("user_id not found in token")
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user_id format: %v", err)
	}

	return userId, nil

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

func VerifyRefreshToken(tokenString string) (uuid.UUID, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpted signing method: %v", token.Header["alg"])

		}
		return []byte(config.ConfigData.RefreshTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return uuid.Nil, fmt.Errorf("could not parse claims")
	}

	userIdStr, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("user_id not found in token")
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user_id format: %v", err)
	}

	return userId, nil

}
