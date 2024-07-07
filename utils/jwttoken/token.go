package jwttoken

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

var secretKey = os.Getenv("JWT_SECRET")

func EncodeJwt(payload jwt.MapClaims) (string, error) {
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET environment variable is not set")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DecodeJwt(tokenString string) (jwt.MapClaims, error) {
	if secretKey == "" {
		return nil, fmt.Errorf("JWT_SECRET environment variable is not set")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
