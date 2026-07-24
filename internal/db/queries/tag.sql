-- name: GetTag :one
SELECT id, name, created_at FROM tag
WHERE id = ?;

-- name: GetTagByName :one
SELECT id, name, created_at FROM tag
WHERE name = ?;

-- name: ListTags :many
SELECT id, name, created_at FROM tag
ORDER BY name ASC;

-- name: CreateTag :one
INSERT INTO tag (name, created_at)
VALUES (?, datetime('now'))
RETURNING id, name, created_at;

-- name: DeleteTag :exec
DELETE FROM tag
WHERE id = ?;
