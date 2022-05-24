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

var SECRET_KEY string

func init(){
  SECRET_KEY = os.Getenv("TOKEN_SECRET_KEY")
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
  ss, err := token.SigningString()
  return ss, err
}

func ParseToken(tokenString string) (*jwt.Token, TokenClaims, error) {
  token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
    if t.Method != jwt.SigningMethodHS256 {
      return nil, fmt.Errorf("Invalid siging method")
    }
    return SECRET_KEY, nil
  })
  if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
    return token, *claims, nil
  }else{
    return nil, TokenClaims{}, err
  }
}
