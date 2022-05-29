package users

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"

	apiDT "github.com/JesusJMM/monomio/api/apiDataTypes"
	"github.com/JesusJMM/monomio/postgres"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
  db postgres.Queries
}

func New(db postgres.Queries) *UserHandler {
  return &UserHandler{db: db}
}

func (h *UserHandler) GetAuthors() gin.HandlerFunc{
  return func(c *gin.Context) {
    users, err := h.db.GetUsersAndBio(context.Background())
    if err != nil {
      c.String(http.StatusInternalServerError, "Internal Server Error")
      log.Println(err)
      return
    }
    var out []apiDT.ResponseUser
    for _, user := range users {
      out = append(out, apiDT.ResponseUser{
        ID: int(user.ID),
        Name: user.Name,
        ImgURL: user.ImgUrl.String,
      })
    }
    c.JSON(http.StatusOK, out)
  }
}

func (h *UserHandler) GetAuthor() gin.HandlerFunc{
  return func(c *gin.Context) {
    name := c.Param("name")
    user, err := h.db.GetUserByName(context.Background(), name)
    if err != nil {
      if errors.Is(err, sql.ErrNoRows) {
        c.String(http.StatusNotFound, "User not found")
        return
      }
      c.String(http.StatusInternalServerError, "Internal Server Error")
      log.Println(err)
      return
    }
    out := apiDT.ResponseUser{
      ID: int(user.ID),
      Name: user.Name,
      ImgURL: user.ImgUrl.String,
    }
    c.JSON(http.StatusOK, out)
  }
}
