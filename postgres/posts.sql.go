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
  content
)
VALUES (
  $1,
  $2,
  $3,
  $4
) RETURNING id, user_id, create_at, title, description, content
`

type CreatePostParams struct {
	UserID      int64
	Title       string
	Description sql.NullString
	Content     sql.NullString
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.UserID,
		arg.Title,
		arg.Description,
		arg.Content,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreateAt,
		&i.Title,
		&i.Description,
		&i.Content,
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

const getPostByAuthorAndTitle = `-- name: GetPostByAuthorAndTitle :one
SELECT 
  posts.id, posts.user_id, posts.create_at, posts.title, posts.description, posts.content,
  users.name AS authorName,
  users.img_url AS authorImgURL
FROM posts
LEFT JOIN users ON users.id = posts.user_id
WHERE users.name = $1 AND posts.title = $2
LIMIT 1
`

type GetPostByAuthorAndTitleParams struct {
	Name  string
	Title string
}

type GetPostByAuthorAndTitleRow struct {
	ID           int64
	UserID       int64
	CreateAt     time.Time
	Title        string
	Description  sql.NullString
	Content      sql.NullString
	Authorname   sql.NullString
	Authorimgurl sql.NullString
}

func (q *Queries) GetPostByAuthorAndTitle(ctx context.Context, arg GetPostByAuthorAndTitleParams) (GetPostByAuthorAndTitleRow, error) {
	row := q.db.QueryRowContext(ctx, getPostByAuthorAndTitle, arg.Name, arg.Title)
	var i GetPostByAuthorAndTitleRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreateAt,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.Authorname,
		&i.Authorimgurl,
	)
	return i, err
}

const getPostWithAuthor = `-- name: GetPostWithAuthor :one
SELECT 
  posts.id, posts.user_id, posts.create_at, posts.title, posts.description, posts.content,
  users.name AS authorName,
  users.img_url AS authorImgURL
FROM posts
LEFT JOIN users ON users.id = posts.user_id
WHERE posts.id = $1
`

type GetPostWithAuthorRow struct {
	ID           int64
	UserID       int64
	CreateAt     time.Time
	Title        string
	Description  sql.NullString
	Content      sql.NullString
	Authorname   sql.NullString
	Authorimgurl sql.NullString
}

func (q *Queries) GetPostWithAuthor(ctx context.Context, id int64) (GetPostWithAuthorRow, error) {
	row := q.db.QueryRowContext(ctx, getPostWithAuthor, id)
	var i GetPostWithAuthorRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreateAt,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.Authorname,
		&i.Authorimgurl,
	)
	return i, err
}

const getPosts = `-- name: GetPosts :many
SELECT id, user_id, create_at, title, description, content FROM posts
ORDER BY created_at
`

func (q *Queries) GetPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CreateAt,
			&i.Title,
			&i.Description,
			&i.Content,
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

const getPostsByAuthor = `-- name: GetPostsByAuthor :one
SELECT 
  posts.id, posts.user_id, posts.create_at, posts.title, posts.description, posts.content,
  users.name AS authorName,
  users.img_url AS authorImgURL
FROM posts
LEFT JOIN users ON users.id = posts.user_id
WHERE users.name = $1
`

type GetPostsByAuthorRow struct {
	ID           int64
	UserID       int64
	CreateAt     time.Time
	Title        string
	Description  sql.NullString
	Content      sql.NullString
	Authorname   sql.NullString
	Authorimgurl sql.NullString
}

func (q *Queries) GetPostsByAuthor(ctx context.Context, name string) (GetPostsByAuthorRow, error) {
	row := q.db.QueryRowContext(ctx, getPostsByAuthor, name)
	var i GetPostsByAuthorRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreateAt,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.Authorname,
		&i.Authorimgurl,
	)
	return i, err
}

const getPostsByAuthorPaginated = `-- name: GetPostsByAuthorPaginated :many
SELECT 
  posts.id, posts.user_id, posts.create_at, posts.title, posts.description, posts.content,
  users.name AS authorName,
  users.img_url AS authorImgURL
FROM posts
LEFT JOIN users ON users.id = posts.user_id
WHERE users.name = $1
LIMIT $2
OFFSET $3
`

type GetPostsByAuthorPaginatedParams struct {
	Name   string
	Limit  int32
	Offset int32
}

type GetPostsByAuthorPaginatedRow struct {
	ID           int64
	UserID       int64
	CreateAt     time.Time
	Title        string
	Description  sql.NullString
	Content      sql.NullString
	Authorname   sql.NullString
	Authorimgurl sql.NullString
}

func (q *Queries) GetPostsByAuthorPaginated(ctx context.Context, arg GetPostsByAuthorPaginatedParams) ([]GetPostsByAuthorPaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByAuthorPaginated, arg.Name, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostsByAuthorPaginatedRow
	for rows.Next() {
		var i GetPostsByAuthorPaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CreateAt,
			&i.Title,
			&i.Description,
			&i.Content,
			&i.Authorname,
			&i.Authorimgurl,
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

const getPostsWithAuthor = `-- name: GetPostsWithAuthor :many
SELECT 
  posts.id, posts.user_id, posts.create_at, posts.title, posts.description, posts.content,
  users.name AS authorName,
  users.img_url AS authorImgURL
FROM posts
LEFT JOIN users ON users.id = posts.user_id
`

type GetPostsWithAuthorRow struct {
	ID           int64
	UserID       int64
	CreateAt     time.Time
	Title        string
	Description  sql.NullString
	Content      sql.NullString
	Authorname   sql.NullString
	Authorimgurl sql.NullString
}

func (q *Queries) GetPostsWithAuthor(ctx context.Context) ([]GetPostsWithAuthorRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsWithAuthor)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostsWithAuthorRow
	for rows.Next() {
		var i GetPostsWithAuthorRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CreateAt,
			&i.Title,
			&i.Description,
			&i.Content,
			&i.Authorname,
			&i.Authorimgurl,
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

const getPostsWithAuthorPaginated = `-- name: GetPostsWithAuthorPaginated :many
SELECT 
  posts.id, posts.user_id, posts.create_at, posts.title, posts.description, posts.content,
  users.name AS authorName,
  users.img_url AS authorImgURL
FROM posts
LEFT JOIN users ON users.id = posts.user_id
LIMIT $1
OFFSET $2
`

type GetPostsWithAuthorPaginatedParams struct {
	Limit  int32
	Offset int32
}

type GetPostsWithAuthorPaginatedRow struct {
	ID           int64
	UserID       int64
	CreateAt     time.Time
	Title        string
	Description  sql.NullString
	Content      sql.NullString
	Authorname   sql.NullString
	Authorimgurl sql.NullString
}

func (q *Queries) GetPostsWithAuthorPaginated(ctx context.Context, arg GetPostsWithAuthorPaginatedParams) ([]GetPostsWithAuthorPaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsWithAuthorPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostsWithAuthorPaginatedRow
	for rows.Next() {
		var i GetPostsWithAuthorPaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CreateAt,
			&i.Title,
			&i.Description,
			&i.Content,
			&i.Authorname,
			&i.Authorimgurl,
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

const getSinglePost = `-- name: GetSinglePost :one
SELECT id, user_id, create_at, title, description, content FROM posts WHERE id=$1 LIMIT 1
`

func (q *Queries) GetSinglePost(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, getSinglePost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreateAt,
		&i.Title,
		&i.Description,
		&i.Content,
	)
	return i, err
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET title=$1, description=$2, content=$3
WHERE id=$4
RETURNING id, user_id, create_at, title, description, content
`

type UpdatePostParams struct {
	Title       string
	Description sql.NullString
	Content     sql.NullString
	ID          int64
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost,
		arg.Title,
		arg.Description,
		arg.Content,
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
	)
	return i, err
}
