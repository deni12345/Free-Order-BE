package middleware

import (
	"context"
	"fmt"
	"github/free-order-be/internal/auth"

	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type claimKey struct{}

func GetClaimKey() claimKey {
	return claimKey{}
}

func AuthenWithJWToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claim := jwt.MapClaims{}
		tokenStr := r.Header.Get("Authorization")
		if len(tokenStr) == 0 {
			UnauthenticatedError(w, fmt.Errorf("token is required"))
			return
		}
		if strings.Contains(tokenStr, "Bearer") {
			tokenStr = strings.Split(tokenStr, "Bearer ")[1]
		}
		err := auth.VerifyToken(tokenStr, &claim)
		if err != nil {
			UnauthenticatedError(w, fmt.Errorf("token is not valid"))
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), claimKey{}, claim))
		next.ServeHTTP(w, r)
	})
}
