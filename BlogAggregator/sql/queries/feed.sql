-- name: AddFeed :one
INSERT INTO feeds (id,name,created_at,updated_at,url,user_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetFeed :many
SELECT * FROM feeds;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;
