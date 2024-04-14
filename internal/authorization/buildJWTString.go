package authorization

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (a *Authorization) BuildJWTString(role string, tokenEXP time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenEXP)),
		},
		Role: role,
	})

	tokenString, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		return "", fmt.Errorf("cannot get token: %v", err)
	}

	return tokenString, nil
}
