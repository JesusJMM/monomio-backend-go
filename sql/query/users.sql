-- name: CreateUser :one
INSERT INTO users (
  name, password, img_url
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET name=$1,password=$2,img_url=$3 
WHERE id=$4
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByName :one
SELECT *
FROM users
WHERE name = $1
LIMIT 1;

-- name: CreateBio :one
INSERT INTO bios (
  user_id, bio
) VALUES (
$1, $2
)
RETURNING *;

-- name: UpdateBio :one
UPDATE bios
SET bio=$1
WHERE id=$2
RETURNING *;

-- name: GetUsersAndBio :many
SELECT u.id, u.name, u.img_url, b.bio 
FROM users u
LEFT JOIN bios b
ON b.user_id = u.id;

-- name: GetUserAndBioByName :one
SELECT u.id, u.name, u.img_url, b.bio
FROM users u
LEFT JOIN bios b
ON b.user_id = u.id
WHERE u.name = $1;
