-- name: CreateRefreshToken :one
INSERT INTO
  refresh_tokens (
    token,
    created_at,
    updated_at,
    user_id,
    expires_at,
    revoked_at
  )
VALUES
  (
    $1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    $2,
    $3,
    NULL
  )
RETURNING
  *;

-- name: GetUserFromRefreshToken :one
SELECT
  user_id,
  expires_at
FROM
  refresh_tokens
WHERE
  token = $1;

-- name: RevokeByToken :one
UPDATE refresh_tokens
SET
  updated_at = CURRENT_TIMESTAMP,
  revoked_at = CURRENT_TIMESTAMP
WHERE
  token = $1
RETURNING
  *;
