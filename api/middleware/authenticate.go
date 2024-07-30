package middleware

import (
	"fmt"
	"net/http"
)

func AuthenWithJWToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if len(token) == 0 {
			UnauthenticatedError(w, fmt.Errorf("token is not valid"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
