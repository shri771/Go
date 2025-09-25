-- name: AddChirp :one
INSERT INTO
  chiprs (id, created_at, updated_at, body, user_id)
VALUES
  (
    gen_random_uuid(),
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    $1,
    $2
  )
RETURNING
  *;

-- name: GetAllChirpAsc :many
SELECT
  *
FROM
  chiprs
ORDER BY
  created_at ASC;

-- name: GetChiprByID :one
SELECT
  *
FROM
  chiprs
WHERE
  id = $1;

-- name: DeleteChiprByID :exec
DELETE FROM chiprs
WHERE
  id = $1;

-- name: GetChiprByUserID :many
SELECT
  *
FROM
  chiprs
WHERE
  user_id = $1;
