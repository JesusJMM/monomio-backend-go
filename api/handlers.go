package api

import (
	"github.com/JesusJMM/monomio/api/auth"
	"github.com/JesusJMM/monomio/api/posts"
	"github.com/JesusJMM/monomio/api/users"
	"github.com/JesusJMM/monomio/postgres"
	"github.com/gin-gonic/gin"
)

func NewHandler(db postgres.Queries) *gin.Engine {
  r := gin.Default()

  api := r.Group("/api")
  authorized := api.Group("")
  authorized.Use(auth.AuthRequired)

  authH := auth.AuthHandler{DB: db}
  api.POST("/auth/signup", authH.Signup())
  api.POST("/auth/login", authH.Login())

  postH := posts.New(db)
  {
    api.GET("/posts", postH.GetAllPosts())
    api.GET("/post/:user/:title", postH.PostByUserAndTitle())
    api.GET("/posts/:id", postH.PostByID())

    authorized.POST("/posts/", postH.Create())
    authorized.PUT("/posts/", postH.Update())
    authorized.DELETE("/posts/:id", postH.Delete())
  }
  userH := users.New(db)
  {
    api.GET("/authors", userH.GetAuthors())
    api.GET("/author/:name", userH.GetAuthor())
  }

  return r
}
