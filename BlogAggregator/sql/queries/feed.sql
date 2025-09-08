-- name: AddFeed :one
INSERT INTO
  feeds (id, name, created_at, updated_at, url, user_id)
VALUES
  ($1, $2, $3, $4, $5, $6)
RETURNING
  *;

-- name: GetFeed :many
SELECT
  *
FROM
  feeds;

-- name: GetUserById :one
SELECT
  *
FROM
  users
WHERE
  id = $1;

-- name: GetFeedIdByUrl :one
SELECT
  id
FROM
  feeds
WHERE
  url = $1;

-- name: GetFeedFollowsForUser :many
SELECT
  feeds.name
FROM
  feeds
  LEFT JOIN users ON users.id = feeds.user_id
WHERE
  users.name = $1;
