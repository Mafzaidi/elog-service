package token

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mafzaidi/elog/pkg/authorizer"
)

type JWTGen struct {
	Secret string
	Claims *authorizer.Claims
}

func Generate(t *JWTGen) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		t.Claims,
	)
	tokenString, err := token.SignedString([]byte(t.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Validate(cookie, secret string) (*authorizer.Claims, error) {
	token, err := jwt.ParseWithClaims(cookie, &authorizer.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	claims, ok := token.Claims.(*authorizer.Claims)
	if !ok {
		return claims, errors.New("token invalid")
	}

	if err != nil || !token.Valid {
		return claims, err
	}
	return claims, nil
}
