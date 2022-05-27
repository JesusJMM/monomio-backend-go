package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	UID int `json:"uid"`
	jwt.RegisteredClaims
}

var SECRET_KEY []byte

func init() {
	SECRET_KEY = []byte(os.Getenv("TOKEN_SECRET_KEY"))
}

func SignToken(UID int) (string, error) {
	claims := TokenClaims{
		UID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SECRET_KEY)
	return ss, err
}

func ParseToken(tokenString string) (*jwt.Token, *TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
    if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Unspected signing method: %v", t.Header["alg"])
    }
		return SECRET_KEY, nil
	})
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return token, claims, nil
	} else {
		return nil, nil, err
	}
}
