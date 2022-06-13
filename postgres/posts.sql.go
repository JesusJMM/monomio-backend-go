// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: posts.sql

package postgres

import (
	"context"
	"database/sql"
	"time"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
  user_id,
  title,
  description,
  content,
  slug,
  feed_img,
  article_img
)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
) RETURNING id, user_id, create_at, title, description, content, feed_img, article_img, slug, published, updated_at
`

type CreatePostParams struct {
	UserID      int64
	Title       string
	Description sql.NullString
	Content     sql.NullString
	Slug        string
	FeedImg     sql.NullString
	ArticleImg  sql.NullString
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.UserID,
		arg.Title,
		arg.Description,
		arg.Content,
		arg.Slug,
		arg.FeedImg,
		arg.ArticleImg,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreateAt,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.FeedImg,
		&i.ArticleImg,
		&i.Slug,
		&i.Published,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts WHERE id=$1 AND user_id=$2
`

type DeletePostParams struct {
	ID     int64
	UserID int64
}

func (q *Queries) DeletePost(ctx context.Context, arg DeletePostParams) error {
	_, err := q.db.ExecContext(ctx, deletePost, arg.ID, arg.UserID)
	return err
}

const getAllPosts = `-- name: GetAllPosts :many
SELECT 
  posts.id, posts.user_id, posts.create_at, posts.title, posts.description, posts.content, posts.feed_img, posts.article_img, posts.slug, posts.published, posts.updated_at, 
  users.name user_name,
  users.img_url user_img_url
FROM posts
LEFT JOIN users 
ON posts.user_id = users.id
`

type GetAllPostsRow struct {
	ID          int64
	UserID      int64
	CreateAt    time.Time
	Title       string
	Description sql.NullString
	Content     sql.NullString
	FeedImg     sql.NullString
	ArticleImg  sql.NullString
	Slug        string
	Published   sql.NullBool
	UpdatedAt   time.Time
	UserName    sql.NullString
	UserImgUrl  sql.NullString
}

func (q *Queries) GetAllPosts(ctx context.Context) ([]GetAllPostsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllPostsRow
	for rows.Next() {
		var i GetAllPostsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CreateAt,
			&i.Title,
			&i.Description,
			&i.Content,
			&i.FeedImg,
			&i.ArticleImg,
			&i.Slug,
			&i.Published,
			&i.UpdatedAt,
			&i.UserName,
			&i.UserImgUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsPaginated = `-- name: GetPostsPaginated :many
SELECT 
  posts.id, posts.user_id, posts.create_at, posts.title, posts.description, posts.content, posts.feed_img, posts.article_img, posts.slug, posts.published, posts.updated_at, 
  users.name user_name,
  users.img_url user_img_url
FROM posts
LEFT JOIN users 
ON posts.user_id = users.id
ORDER BY posts.create_at DESC
LIMIT $1
OFFSET $2
`

type GetPostsPaginatedParams struct {
	Limit  int32
	Offset int32
}

type GetPostsPaginatedRow struct {
	ID          int64
	UserID      int64
	CreateAt    time.Time
	Title       string
	Description sql.NullString
	Content     sql.NullString
	FeedImg     sql.NullString
	ArticleImg  sql.NullString
	Slug        string
	Published   sql.NullBool
	UpdatedAt   time.Time
	UserName    sql.NullString
	UserImgUrl  sql.NullString
}

func (q *Queries) GetPostsPaginated(ctx context.Context, arg GetPostsPaginatedParams) ([]GetPostsPaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostsPaginatedRow
	for rows.Next() {
		var i GetPostsPaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CreateAt,
			&i.Title,
			&i.Description,
			&i.Content,
			&i.FeedImg,
			&i.ArticleImg,
			&i.Slug,
			&i.Published,
			&i.UpdatedAt,
			&i.UserName,
			&i.UserImgUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const postBySlugAndUser = `-- name: PostBySlugAndUser :one
SELECT 
  posts.id, posts.user_id, posts.create_at, posts.title, posts.description, posts.content, posts.feed_img, posts.article_img, posts.slug, posts.published, posts.updated_at,
  users.name user_name,
  users.img_url user_img_url
FROM posts
LEFT JOIN users 
ON posts.user_id = users.id
WHERE posts.slug = $1 AND users.id = $2
ORDER BY posts.create_at DESC
LIMIT 1
`

type PostBySlugAndUserParams struct {
	Slug string
	ID   int64
}

type PostBySlugAndUserRow struct {
	ID          int64
	UserID      int64
	CreateAt    time.Time
	Title       string
	Description sql.NullString
	Content     sql.NullString
	FeedImg     sql.NullString
	ArticleImg  sql.NullString
	Slug        string
	Published   sql.NullBool
	UpdatedAt   time.Time
	UserName    sql.NullString
	UserImgUrl  sql.NullString
}

func (q *Queries) PostBySlugAndUser(ctx context.Context, arg PostBySlugAndUserParams) (PostBySlugAndUserRow, error) {
	row := q.db.QueryRowContext(ctx, postBySlugAndUser, arg.Slug, arg.ID)
	var i PostBySlugAndUserRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreateAt,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.FeedImg,
		&i.ArticleImg,
		&i.Slug,
		&i.Published,
		&i.UpdatedAt,
		&i.UserName,
		&i.UserImgUrl,
	)
	return i, err
}

const publishPost = `-- name: PublishPost :exec
UPDATE posts SET published = TRUE WHERE posts.id = $1
`

func (q *Queries) PublishPost(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, publishPost, id)
	return err
}

const unPublishPost = `-- name: UnPublishPost :exec
UPDATE posts SET published = FALSE WHERE posts.id = $1
`

func (q *Queries) UnPublishPost(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, unPublishPost, id)
	return err
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts SET 
  title=$1,
  description=$2,
  content=$3,
  feed_img=$4, 
  article_img=$5,
  slug=$6,
  updated_at=CURRENT_TIMESTAMP
WHERE id=$7
RETURNING id, user_id, create_at, title, description, content, feed_img, article_img, slug, published, updated_at
`

type UpdatePostParams struct {
	Title       string
	Description sql.NullString
	Content     sql.NullString
	FeedImg     sql.NullString
	ArticleImg  sql.NullString
	Slug        string
	ID          int64
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost,
		arg.Title,
		arg.Description,
		arg.Content,
		arg.FeedImg,
		arg.ArticleImg,
		arg.Slug,
		arg.ID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreateAt,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.FeedImg,
		&i.ArticleImg,
		&i.Slug,
		&i.Published,
		&i.UpdatedAt,
	)
	return i, err
}
