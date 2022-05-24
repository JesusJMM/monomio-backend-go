package auth

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/JesusJMM/monomio/api/apiDataTypes"
	"github.com/JesusJMM/monomio/postgres"
)

type AuthHandler struct {
  DB postgres.Queries
}

func (a *AuthHandler) Signup() gin.HandlerFunc{
  return func(ctx *gin.Context) {
    var payload apiDT.PayloadSignup
    if err := ctx.BindJSON(&payload); err != nil {
      log.Println(err)
      ctx.String(http.StatusBadRequest, "error: %w", err)
      return
    }
    hashedPassword, err := HashPassword(payload.Password)
    if err != nil {
      log.Println(err)
      ctx.String(http.StatusInternalServerError, "Internal Server Error")
      return
    }
    user := postgres.CreateUserParams{
      Name: payload.Name,
      Password: hashedPassword,
      ImgUrl: sql.NullString{ String: payload.ImgURL, Valid: payload.ImgURL != ""},
    }
    newUser, err := a.DB.CreateUser(context.Background(), user)
    if err != nil {
      if strings.Contains(err.Error(), "duplicate key value violates unique constraint"){
        ctx.String(http.StatusConflict, "User name already taken")
        return
      }
      log.Println(err)
      ctx.String(http.StatusInternalServerError, "Internal Server Error, error: %w", err)
      return
    }

    token, err := SignToken(int(newUser.ID))
    if err != nil {
      ctx.String(http.StatusCreated, "Internal Server Error, error: %w", err)
      log.Println(err)
      return
    }
  
    ctx.Header("Authorization", token)
    responseUser := apiDT.User{
      ID: int(newUser.ID),
      Name: newUser.Name,
      ImgURL: newUser.ImgUrl.String,
    }
    ctx.JSON(http.StatusCreated, apiDT.ResponseSignup{User: responseUser})
  }
}

func (a *AuthHandler) Login() gin.HandlerFunc{
  return func(ctx *gin.Context) {
    var payload apiDT.PayloadLogin
    if err := ctx.BindJSON(&payload); err != nil {
      ctx.String(http.StatusBadRequest, "Bad request, error: %w", err)
    }
    dbUser, err := a.DB.GetUserByName(context.Background(), payload.Name)
    if err != nil{
      if err == sql.ErrNoRows{
        ctx.String(http.StatusNotFound, "User does not exits")
        return
      }
      ctx.String(http.StatusInternalServerError, "Internal Server Error, error: %w", err)
      return
    }
    if !CheckPasswordHash(payload.Password, dbUser.Password){
      ctx.String(http.StatusForbidden, "Invalid password", err)
      return
    }
    token, err := SignToken(int(dbUser.ID))
    if err != nil {
      ctx.String(http.StatusCreated, "Internal Server Error, error: %w", err)
      log.Println(err)
      return
    }
    ctx.Header("Authorization", token)
    responseUser := apiDT.User{
      ID: int(dbUser.ID),
      Name: dbUser.Name,
      ImgURL: dbUser.ImgUrl.String,
    }
    ctx.JSON(http.StatusOK, apiDT.ResponseLogin{User: responseUser})
  }
}

