package auth

import (
	"fmt"
	"github/free-order-be/config"
	"github/free-order-be/internal/domain"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": user.ID,
			"exp":    time.Now().Add(24 * time.Hour).Unix(),
		})
	tokenStr, err := token.SignedString(config.Values.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func VerifyToken(tokenStr string, claim *jwt.MapClaims) error {
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(t *jwt.Token) (interface{}, error) {
		return config.Values.SecretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
