package posts

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	apiDT "github.com/JesusJMM/monomio/api/apiDataTypes"
	"github.com/JesusJMM/monomio/api/auth"
	"github.com/JesusJMM/monomio/postgres"
	"github.com/gin-gonic/gin"
)

type PostsHandler struct {
	db postgres.Queries
}

func New(db postgres.Queries) PostsHandler {
	return PostsHandler{
		db: db,
	}
}

func (h PostsHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenClaims, err := auth.GetTokenClaimsFromContext(ctx)

		var payload apiDT.PayloadCreatePost
		if err := ctx.BindJSON(&payload); err != nil {
			ctx.String(http.StatusBadRequest, "error: %w", err)
			return
		}

		if _, err := h.db.PostBySlugAndUserID(context.Background(), postgres.PostBySlugAndUserIDParams{
			Slug: payload.Slug,
			ID:   int64(tokenClaims.UID),
		}); err == nil {
			if errors.Is(err, sql.ErrNoRows) {
				ctx.String(http.StatusConflict, "slug already used")
			} else {
				ctx.String(http.StatusInternalServerError, "internal server error")
			}
			return
		}

		postInput := postgres.CreatePostParams{
			Title:       payload.Title,
			Description: sql.NullString{String: payload.Description, Valid: true},
			Content:     sql.NullString{String: payload.Content, Valid: true},
			UserID:      int64(tokenClaims.UID),
			Slug:        payload.Slug,
			FeedImg:     sql.NullString{String: payload.FeedImg, Valid: payload.FeedImg != ""},
			ArticleImg:  sql.NullString{String: payload.ArticleImg, Valid: payload.ArticleImg != ""},
		}
		newUser, err := h.db.CreatePost(context.Background(), postInput)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error creating resource : %w", err)
			return
		}
		ctx.JSON(http.StatusCreated, apiDT.ResponseCreatePost{ID: int(newUser.ID)})
	}
}

func (h PostsHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenClaims, _ := auth.GetTokenClaimsFromContext(ctx)

		var payload apiDT.PayloadUpdatePost
		if err := ctx.BindJSON(&payload); err != nil {
			ctx.String(http.StatusBadRequest, "error: %w", err)
			return
		}

		// verify if post exist and if the user is the author of the post
		dbPost, err := h.db.PostByID(context.Background(), int64(payload.ID))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				ctx.String(http.StatusBadRequest, "Resource do not exist")
				return
			}
		}

		if dbPost.UserID != int64(tokenClaims.UID) {
			ctx.String(http.StatusForbidden, "You does not have permissions to modify the resource")
			return
		}

    fmt.Println(payload.Content == "")
		postInput := postgres.UpdatePostParams{
			Title:       InvTernaryfunc(payload.Title, "", dbPost.Title),
			Description: sql.NullString{String: InvTernaryfunc(payload.Description, "", dbPost.Description.String), Valid: true},
			Content:     sql.NullString{String: InvTernaryfunc(payload.Content, "", dbPost.Content.String), Valid: true},
			Slug:        InvTernaryfunc(payload.Slug, "", dbPost.Slug),
			ID:          int64(payload.ID),
			FeedImg:     sql.NullString{String: InvTernaryfunc(payload.FeedImg, "", dbPost.FeedImg.String), Valid: true},
			ArticleImg:  sql.NullString{String: InvTernaryfunc(payload.ArticleImg, "", dbPost.ArticleImg.String), Valid: true},
		}

		newUser, err := h.db.UpdatePost(context.Background(), postInput)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error creating resource : %w", err)
			return
		}
		ctx.JSON(http.StatusOK, apiDT.ResponseCreatePost{ID: int(newUser.ID)})
	}
}

func (h *PostsHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenClaims, _ := auth.GetTokenClaimsFromContext(ctx)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.String(http.StatusBadRequest, "'id' url param must be a number")
		}
		h.db.DeletePost(context.Background(), postgres.DeletePostParams{
			ID:     int64(id),
			UserID: int64(tokenClaims.UID),
		})
	}
}

// Inverse Ternary
// returns a if a is diferent to value
// else return def
func InvTernaryfunc[T comparable](a, value, def T) T {
	if a == value {
		return def
	}
	return a
}
