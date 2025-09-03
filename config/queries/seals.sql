-- name: GetSeal :one
SELECT * FROM seals
WHERE id = ? LIMIT 1;

-- name: CountSeals :one
SELECT count(*) FROM seals;

-- name: CountSealsByTag :one
SELECT count(*) FROM seals
INNER JOIN seal_tags ON seal_tags.seal_id = seals.id
INNER JOIN tags ON seal_tags.tag_id = tags.id
WHERE tags.name = ?;

-- name: ListSeals :many
SELECT * FROM seals
ORDER BY path;

-- name: ListSealTags :many
SELECT tags.* FROM tags
INNER JOIN seal_tags ON seal_tags.tag_id = tags.id
WHERE seal_tags.seal_id = ?;

-- name: CreateSeal :one
INSERT INTO seals (
  path, mime_type
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateSeal :one
UPDATE seals
set path = ?,
mime_type = ?
WHERE id = ?
RETURNING *;

-- name: DeleteSeal :exec
DELETE FROM seals
WHERE id = ?;

-- name: ListSealsWithTag :many
SELECT seals.* FROM seals
INNER JOIN seal_tags ON seal_tags.seal_id = seals.id
INNER JOIN tags ON seal_tags.tag_id = tags.id
WHERE tags.name = ?;

-- name: RandomSeal :one
SELECT * from seals
ORDER BY RANDOM() LIMIT 1;

-- name: RandomSealWithTag :one
SELECT seals.* FROM seals
INNER JOIN seal_tags ON seal_tags.seal_id = seals.id
INNER JOIN tags ON seal_tags.tag_id = tags.id
WHERE tags.name = ?
ORDER BY RANDOM() LIMIT 1;

-- name: AddSealTag :one
INSERT INTO seal_tags (
  seal_id, tag_id
) VALUES (
  ?, ?
)
RETURNING *;

