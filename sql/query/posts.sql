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

-- name: GetPosts :many
SELECT * FROM posts
ORDER BY created_at;

-- name: GetPostsByUser :many
SELECT * FROM posts
WHERE user_id = $1
ORDER BY created_at;

-- name: DeletePost :exec
DELETE FROM posts WHERE id=$1;
