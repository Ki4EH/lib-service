package catalog

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

const signingKey = "342ij432"

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
	Role   int `json: "role"`
}

func ParseToken(accessToken string) (int, int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, claims.Role, nil
}
