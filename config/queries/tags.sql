-- name: GetTag :one
SELECT * FROM tags
WHERE id = ? LIMIT 1;

-- name: ListTags :many
SELECT * FROM tags
ORDER BY name;

-- name: GetTagWithName :one
SELECT * from tags
WHERE name = ?;

-- name: CreateTag :one
INSERT INTO tags (name) VALUES (?)
RETURNING *;

-- name: GetPopularTags :many
SELECT tags.name,COUNT(tag_id) as "count"
FROM seal_tags INNER JOIN tags ON tags.id == tag_id
WHERE tags.name != 'jpeg' AND tags.name != 'png' AND tags.name != 'gif'
GROUP BY tag_id ORDER BY "count" DESC LIMIT 10;
