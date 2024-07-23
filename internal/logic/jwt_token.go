package logic

import (
	"fmt"
	"github/lambda-microservice/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (l *LogicImpl) createToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": user.ID,
			"exp":    time.Now().Add(24 * time.Hour).Unix(),
		})
	tokenStr, err := token.SignedString(l.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (l *LogicImpl) verifyToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return l.SecretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
