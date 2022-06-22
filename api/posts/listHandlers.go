package posts

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"

	apiDT "github.com/JesusJMM/monomio/api/apiDataTypes"
	"github.com/JesusJMM/monomio/api/apiutils"
	"github.com/JesusJMM/monomio/api/auth"
	"github.com/JesusJMM/monomio/postgres"
	"github.com/gin-gonic/gin"
)

func (h *PostsHandler) GetAllPosts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		posts, err := h.db.GetAllPosts(context.Background())
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
    out := []apiDT.ResponseCompletePost{}
		for _, p := range posts {
			out = append(out, apiDT.ResponseCompletePost{
				ID:           int(p.ID),
				Title:        p.Title,
				Description:  p.Description.String,
				CreatedAt:    p.CreateAt,
				AuthorName:   p.UserName.String,
				AuthorImgURL: p.UserImgUrl.String,
				Content:      p.Content.String,
				Slug:         p.Slug,
				ArticleImg:   p.ArticleImg.String,
				FeedImg:      p.FeedImg.String,
				UpdatedAt:    p.UpdatedAt,
				Published:    p.Published.Bool,
			})
		}
		ctx.JSON(http.StatusOK, out)
	}
}

func (h *PostsHandler) PostsPaginated() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, ok := apiutils.GetIntQueryParam(c, "page", 1)
		if !ok {
			return
		}
		posts, err := h.db.PostsPag(context.Background(), postgres.PostsPagParams{
			Limit:  10,
			Offset: int32(10 * (page - 1)),
		})
		if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			log.Println(err)
			return
		}
    out := []apiDT.ResponseShortPost{}
		for _, p := range posts {
			out = append(out, apiDT.ResponseShortPost{
				ID:           int(p.ID),
				Title:        p.Title,
				Description:  p.Description.String,
				CreatedAt:    p.CreateAt,
				Slug:         p.Slug,
				FeedImg:      p.FeedImg.String,
				AuthorName:   p.UserName.String,
				AuthorImgURL: p.UserImgUrl.String,
			})
		}
		c.JSON(http.StatusOK, out)
	}
}

func (h *PostsHandler) PostByUserAndSlug() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userName := ctx.Param("user")
		postSlug := ctx.Param("slug")
		p, err := h.db.PostBySlugAndUser(context.Background(), postgres.PostBySlugAndUserParams{
			Slug: postSlug,
			Name: userName,
		})
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				ctx.String(http.StatusNotFound, "Not found")
				return
			}
			ctx.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		out := apiDT.ResponseCompletePost{
			ID:           int(p.ID),
			Title:        p.Title,
			Description:  p.Description.String,
			CreatedAt:    p.CreateAt,
			AuthorName:   p.UserName.String,
			AuthorImgURL: p.UserImgUrl.String,
			Content:      p.Content.String,
			Slug:         p.Slug,
			ArticleImg:   p.ArticleImg.String,
			FeedImg:      p.FeedImg.String,
			UpdatedAt:    p.UpdatedAt,
			Published:    p.Published.Bool,
		}
		ctx.JSON(http.StatusOK, out)
	}
}

func (h *PostsHandler) PostByUserPaginated() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.Param("user")
		page, ok := apiutils.GetIntQueryParam(c, "page", 1)
		if !ok {
			return
		}
		posts, err := h.db.PostsPagByUser(context.Background(), postgres.PostsPagByUserParams{
			Name:   userName,
			Limit:  10,
			Offset: int32(10 * (page - 1)),
		})
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.String(http.StatusNotFound, "Not found")
				return
			}
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
    out := []apiDT.ResponseShortPost{}
		for _, p := range posts {
			out = append(out, apiDT.ResponseShortPost{
				ID:           int(p.ID),
				Title:        p.Title,
				Description:  p.Description.String,
				CreatedAt:    p.CreateAt,
				Slug:         p.Slug,
				FeedImg:      p.FeedImg.String,
				AuthorName:   p.UserName.String,
				AuthorImgURL: p.UserImgUrl.String,
			})
		}
		c.JSON(http.StatusOK, out)
	}
}

func (h *PostsHandler) PostByUserPaginatedPrivate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenClaims, _ := auth.GetTokenClaimsFromContext(c)
		page, ok := apiutils.GetIntQueryParam(c, "page", 1)
		if !ok {
			return
		}
		posts, err := h.db.PostsPagByUserPrivate(context.Background(), postgres.PostsPagByUserPrivateParams{
			UserID: int64(tokenClaims.UID),
			Limit:  10,
			Offset: int32(10 * (page - 1)),
		})
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.String(http.StatusNotFound, "Not found")
				return
			}
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
    out := []apiDT.ResponseShortPost {} 
		for _, p := range posts {
			out = append(out, apiDT.ResponseShortPost{
				ID:           int(p.ID),
				Title:        p.Title,
				Description:  p.Description.String,
				CreatedAt:    p.CreateAt,
				Slug:         p.Slug,
				FeedImg:      p.FeedImg.String,
				AuthorName:   p.UserName.String,
				AuthorImgURL: p.UserImgUrl.String,
			})
		}
		c.JSON(http.StatusOK, out)
	}
}
