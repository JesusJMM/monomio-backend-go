package posts

import (
	"context"
	"database/sql"
	"net/http"

	apiDT "github.com/JesusJMM/monomio/api/apiDataTypes"
	"github.com/JesusJMM/monomio/api/auth"
	"github.com/JesusJMM/monomio/postgres"
	"github.com/gin-gonic/gin"
)

type PostsHandler struct {
  db *postgres.Queries
}

func New(db *postgres.Queries) PostsHandler{
  return PostsHandler{
    db: db,
  }
}

func (h PostsHandler) Create() gin.HandlerFunc {
  return func(ctx *gin.Context) {
    tokenClaims, err := auth.GetTokenClaimsFromContext(ctx)
    if err != nil {
      ctx.String(http.StatusForbidden, "error: Invalid Token")
      return
    }

    var payload apiDT.PayloadPost
    if err := ctx.BindJSON(&payload); err != nil {
      ctx.String(http.StatusBadRequest, "error: %w", err)
      return
    }

    postInput := postgres.CreatePostParams{
      Title: payload.Title,
      Description: sql.NullString{String: payload.Description, Valid: true},
      Content: sql.NullString{String: payload.Content, Valid: true},
      UserID: int64(tokenClaims.UID),
    }
    newUser, err := h.db.CreatePost(context.Background(), postInput)
    if err != nil{
      ctx.String(http.StatusInternalServerError, "Error creating resource : %w", err)
      return
    }
    ctx.JSON(http.StatusCreated, apiDT.ResponseCreatePost{ID: int(newUser.ID)})
  }
}
