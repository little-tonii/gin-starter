package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId       int64
	TokenVersion int64
}

func GenerateAccessToken(secretKey string, claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      claims.UserId,
		"version": claims.TokenVersion,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func GenerateRefreshToken(secretKey string, claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      claims.UserId,
		"version": claims.TokenVersion,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(secretKey string, token string) (*Claims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Có lỗi xảy ra")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("Token không hợp lệ")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return nil, errors.New("Token không hợp lệ")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("Token không hợp lệ")
	}
	userId := claims["id"].(float64)
	tokenVersion := claims["version"].(float64)
	return &Claims{
		UserId:       int64(userId),
		TokenVersion: int64(tokenVersion),
	}, nil
}
