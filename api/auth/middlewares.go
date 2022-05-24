package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context){ 
  authHeader := c.GetHeader("authorization")
  if authHeader == ""{
    c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("error: No token is provided"))
    return
  }
  rawToken := strings.Fields(authHeader)
  if len(rawToken) < 2 {
    c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("error: Bad token format"))
    return
  }
  token, claims, err := ParseToken(rawToken[1])
  if err != nil || !token.Valid {
    c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("error: Token expirated"))
    return
  }
  c.Set("tokenClaims", claims)
}
