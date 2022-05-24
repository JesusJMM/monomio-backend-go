package api

import (
	"github.com/JesusJMM/monomio/api/auth"
	"github.com/JesusJMM/monomio/postgres"
	"github.com/gin-gonic/gin"
)

func NewHandler(db postgres.Queries) *gin.Engine {
  r := gin.Default()

  api := r.Group("/api")

  authHandler := auth.AuthHandler{DB: db}
  api.POST("/auth/signup", authHandler.Signup())
  api.POST("/auth/login", authHandler.Login())
  return r
}
