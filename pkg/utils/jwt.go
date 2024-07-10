package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	mySigningKey = []byte("AllYourBase")
)

type Payload struct {
	UserId uint
	Email  string
}

func CreateToken(payload *Payload) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
		Issuer:    "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (bool, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return false, err
	}
	if !token.Valid {
		return false, err
	}

	return true, nil

}
