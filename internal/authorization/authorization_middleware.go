package authorization

import (
	"banner_service/internal/logger"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	Role string
}

func getRole(secretKey, tokenString string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(secretKey), nil
		})

	if err != nil {
		return "", fmt.Errorf("cannot pars: %v", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("token is not valid: %v", err)
	}

	return claims.Role, nil
}

// TO DO разграничить роли
func (a *Authorization) AuthorizationMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("token")
			if token == "" {
				logger.Error("Header do not contain a token")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			role, err := getRole(a.secretKey, token)
			if err != nil {
				logger.Error("token does not pass validation")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			r.Header.Set("role", role)

			next.ServeHTTP(w, r)
		})
	}
}
