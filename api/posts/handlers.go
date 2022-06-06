package posts

import (
	"context"
	"database/sql"
	"errors"
	"log"
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

		postInput := postgres.CreatePostParams{
			Title:       payload.Title,
			Description: sql.NullString{String: payload.Description, Valid: true},
			Content:     sql.NullString{String: payload.Content, Valid: true},
			UserID:      int64(tokenClaims.UID),
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
		dbPost, err := h.db.GetSinglePost(context.Background(), int64(payload.ID))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				if payload.Title == "" || payload.Content == "" || payload.Description == "" {
					ctx.String(http.StatusBadRequest, "Resource don't exist, not enought data")
					return
				}

				newUser, err := h.db.CreatePost(context.Background(), postgres.CreatePostParams{
					Title:       payload.Title,
					Description: sql.NullString{String: payload.Description, Valid: true},
					Content:     sql.NullString{String: payload.Content, Valid: true},
				})
				if err != nil {
					ctx.String(http.StatusInternalServerError, "Error creating resource : %w", err)
					return
				}
				ctx.JSON(http.StatusCreated, apiDT.ResponseCreatePost{ID: int(newUser.ID)})
				return
			}
			ctx.String(http.StatusInternalServerError, "Error updating resource : %w", err)
			return
		}

		if dbPost.UserID != int64(tokenClaims.UID) {
			ctx.String(http.StatusForbidden, "You does not have permissions to modify the resource")
			return
		}

		postInput := postgres.UpdatePostParams{
			Title:       payload.Title,
			Description: sql.NullString{String: payload.Description, Valid: true},
			Content:     sql.NullString{String: payload.Content, Valid: true},
			ID:          int64(payload.ID),
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

func (h *PostsHandler) GetAllPosts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		posts, err := h.db.GetPostsWithAuthor(context.Background())
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		var out []apiDT.ResponseShortPost
		for _, p := range posts {
			out = append(out, apiDT.ResponseShortPost{
				ID:           int(p.ID),
				Title:        p.Title,
				Description:  p.Description.String,
				CreatedAt:    p.CreateAt,
				AuthorName:   p.Authorname.String,
				AuthorImgURL: p.Authorimgurl.String,
			})
		}
		ctx.JSON(http.StatusOK, out)
	}
}

func (h *PostsHandler) PostsPaginated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
    page := 1
    if ctx.Query("id") != "" {
      var err error
      page, err = strconv.Atoi(ctx.Query("id"))
      if err != nil {
        ctx.String(http.StatusBadRequest, "'id' url query param must be a number")
        return
      }
    }
		posts, err := h.db.GetPostsWithAuthorPaginated(context.Background(), postgres.GetPostsWithAuthorPaginatedParams{
			Limit:  10,
			Offset: int32(10 * (page - 1)),
		})
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Internal Server Error")
      log.Println(err)
			return
		}
		var out []apiDT.ResponseShortPost
		for _, p := range posts {
			out = append(out, apiDT.ResponseShortPost{
				ID:           int(p.ID),
				Title:        p.Title,
				Description:  p.Description.String,
				CreatedAt:    p.CreateAt,
				AuthorName:   p.Authorname.String,
				AuthorImgURL: p.Authorimgurl.String,
			})
		}
		ctx.JSON(http.StatusOK, out)
	}
}

func (h *PostsHandler) PostByUserAndTitle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userName := ctx.Param("user")
		postTitle := ctx.Param("title")
		post, err := h.db.GetPostByAuthorAndTitle(context.Background(), postgres.GetPostByAuthorAndTitleParams{
			Name:  userName,
			Title: postTitle,
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
			ID:           int(post.ID),
			Title:        post.Title,
			Description:  post.Description.String,
			Content:      post.Content.String,
			CreatedAt:    post.CreateAt,
			AuthorName:   post.Authorname.String,
			AuthorImgURL: post.Authorimgurl.String,
		}
		ctx.JSON(http.StatusOK, out)
	}
}

func (h *PostsHandler) PostByUserPaginated() gin.HandlerFunc {
  return func(c *gin.Context) {
    userName := c.Param("user")
    var page int = 1
    if c.Query("page") != "" {
      var err error
      page, err = strconv.Atoi(c.Query("page"))
      if err != nil {
        c.String(http.StatusBadRequest, "'page' query param must be a number")
        return
      }
    }
    posts, err := h.db.GetPostsByAuthorPaginated(context.Background(), postgres.GetPostsByAuthorPaginatedParams{
      Name: userName,
      Limit: 10,
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
		var out []apiDT.ResponseShortPost
		for _, p := range posts {
			out = append(out, apiDT.ResponseShortPost{
				ID:           int(p.ID),
				Title:        p.Title,
				Description:  p.Description.String,
				CreatedAt:    p.CreateAt,
				AuthorName:   p.Authorname.String,
				AuthorImgURL: p.Authorimgurl.String,
			})
		}
    c.JSON(http.StatusOK, out)
  }
}

func (h *PostsHandler) PostByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
      ctx.String(http.StatusBadRequest, "'id' url param must be a number")
      return
    }
    post, err := h.db.GetPostWithAuthor(context.Background(), int64(id))
    if err != nil {
      if errors.Is(err, sql.ErrNoRows){
        ctx.String(http.StatusNotFound, "Not found")
        return
      }
			ctx.String(http.StatusInternalServerError, "Internal Server Error")
			return
    }
    out := apiDT.ResponseCompletePost{
			ID:           int(post.ID),
			Title:        post.Title,
			Description:  post.Description.String,
			Content:      post.Content.String,
			CreatedAt:    post.CreateAt,
			AuthorName:   post.Authorname.String,
			AuthorImgURL: post.Authorimgurl.String,
    }
    ctx.JSON(http.StatusOK, out)
	}
}
