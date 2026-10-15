package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("super-secret-key")

func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtKey)
}
