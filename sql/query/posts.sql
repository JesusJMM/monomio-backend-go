-- name: CreatePost :one
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
) RETURNING *;

-- name: UpdatePost :one
UPDATE posts
SET title=$1, description=$2, content=$3
WHERE id=$4
RETURNING *;

-- name: GetSinglePost :one
SELECT * FROM posts WHERE id=$1 LIMIT 1;

-- name: GetPostWithAuthor :one
SELECT 
  posts.*,
  users.name AS authorName,
  users.img_url AS authorImgURL
FROM posts
LEFT JOIN users ON users.id = posts.user_id
WHERE posts.id = $1;

-- name: GetSinglePostByAuthor :one
SELECT 
  posts.*,
  users.name AS authorName,
  users.img_url AS authorImgURL
FROM posts
LEFT JOIN users ON users.id = posts.user_id
WHERE users.id = $1;

-- name: GetPostsByAuthorPaginated :many
SELECT 
  posts.*,
  users.name AS authorName,
  users.img_url AS authorImgURL
FROM posts
LEFT JOIN users ON users.id = posts.user_id
WHERE users.name = $1
LIMIT $2
OFFSET $3;

-- name: GetPostByAuthorAndTitle :one
SELECT 
  posts.*,
  users.name AS authorName,
  users.img_url AS authorImgURL
FROM posts
LEFT JOIN users ON users.id = posts.user_id
WHERE users.name = $1 AND posts.title = $2
LIMIT 1;

-- name: GetPosts :many
SELECT * FROM posts
ORDER BY created_at;

-- name: GetPostsWithAuthor :many
SELECT 
  posts.*,
  users.name AS authorName,
  users.img_url AS authorImgURL
FROM posts
LEFT JOIN users ON users.id = posts.user_id;

-- name: GetPostsWithAuthorPaginated :many
SELECT 
  posts.*,
  users.name AS authorName,
  users.img_url AS authorImgURL
FROM posts
LEFT JOIN users ON users.id = posts.user_id
LIMIT $1
OFFSET $2;

-- name: DeletePost :exec
DELETE FROM posts WHERE id=$1 AND user_id=$2;
