package api

import (
	"github.com/gin-gonic/gin"

	"github.com/JesusJMM/monomio/api/auth"
	"github.com/JesusJMM/monomio/api/posts"
	"github.com/JesusJMM/monomio/api/users"
	"github.com/JesusJMM/monomio/postgres"
)

func NewHandler(db postgres.Queries) *gin.Engine {
  r := gin.Default()

  api := r.Group("/api")

  authH := auth.AuthHandler{DB: db}
  api.POST("/auth/signup", authH.Signup())
  api.POST("/auth/login", authH.Login())

  postH := posts.New(db)
  api.GET("/posts/feed",        postH.PostsPaginated())
  api.GET("/posts/all",         postH.GetAllPosts())
  api.GET("/post/:user/:slug",  postH.PostByUserAndSlug())
  api.GET("/post/:user",        postH.PostByUserPaginated())
  api.GET("/posts/dashboard", auth.AuthRequired,  postH.PostByUserPaginatedPrivate())
  api.POST("/post",   auth.AuthRequired,          postH.Create())
  api.PUT("/post",    auth.AuthRequired,          postH.Update())
  api.DELETE("/post", auth.AuthRequired,          postH.Delete())

  userH := users.New(db)
  {
    api.GET("/authors", userH.GetAuthors())
    api.GET("/author/:name", userH.GetAuthor())
  }

  return r
}
