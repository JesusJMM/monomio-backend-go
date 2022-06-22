-- name: CreatePost :one
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
) RETURNING *;

-- name: UpdatePost :one
UPDATE posts SET 
  title=$1,
  description=$2,
  content=$3,
  feed_img=$4, 
  article_img=$5,
  slug=$6,
  updated_at=CURRENT_TIMESTAMP
WHERE id=$7
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts WHERE id=$1 AND user_id=$2;

-- name: GetAllPosts :many
SELECT 
  posts.*, 
  users.name user_name,
  users.img_url user_img_url
FROM posts
LEFT JOIN users 
ON posts.user_id = users.id
WHERE posts.published = true;

-- name: PostsPag :many
SELECT 
  posts.*, 
  users.name user_name,
  users.img_url user_img_url
FROM posts
LEFT JOIN users 
ON posts.user_id = users.id
WHERE posts.published = true
ORDER BY posts.create_at DESC
LIMIT $1
OFFSET $2;

-- name: PostsPagByUser :many
SELECT 
  posts.*, 
  users.name user_name,
  users.img_url user_img_url
FROM posts
LEFT JOIN users 
ON posts.user_id = users.id
WHERE posts.published = true AND users.name = $1
ORDER BY posts.create_at DESC
LIMIT $2
OFFSET $3;

-- name: PostsPagByUserPrivate :many
SELECT 
  posts.*, 
  users.name user_name,
  users.img_url user_img_url
FROM posts
LEFT JOIN users 
ON posts.user_id = users.id
WHERE posts.user_id = $1
ORDER BY posts.create_at DESC
LIMIT $2
OFFSET $3;

-- name: PublishPost :exec
UPDATE posts SET published = TRUE WHERE posts.id = $1;

-- name: UnpublishPost :exec
UPDATE posts SET published = FALSE WHERE posts.id = $1 AND posts.user_id = $2;

-- name: PostBySlugAndUser :one
SELECT 
  posts.*,
  users.name user_name,
  users.img_url user_img_url
FROM posts
LEFT JOIN users 
ON posts.user_id = users.id
WHERE posts.slug = $1 AND users.name = $2
ORDER BY posts.create_at DESC
LIMIT 1;

-- name: PostBySlugAndUserID :one
SELECT 
  posts.*,
  users.name user_name,
  users.img_url user_img_url
FROM posts
LEFT JOIN users 
ON posts.user_id = users.id
WHERE posts.slug = $1 AND users.ID = $2
ORDER BY posts.create_at DESC
LIMIT 1;

-- name: PostByID :one
SELECT posts.* FROM posts WHERE posts.id = $1;
