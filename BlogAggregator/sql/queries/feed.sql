-- name: AddFeed :one
INSERT INTO
  feeds (id, name, created_at, updated_at, url)
VALUES
  ($1, $2, $3, $4, $5)
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
  feeds.*
FROM
  users
  JOIN feed_follows ON users.id = feed_follows.user_id
  JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE
  users.name = $1;
