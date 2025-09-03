-- name: GetTag :one
SELECT * FROM tags
WHERE id = ? LIMIT 1;

-- name: ListTags :many
SELECT * FROM tags
ORDER BY path;

-- name: GetTagWithName :one
SELECT * from tags
WHERE name = ?;

-- name: CreateTag :one
INSERT INTO tags (name) VALUES (?)
RETURNING *;
