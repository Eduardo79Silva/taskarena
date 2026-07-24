-- name: GetTask :one
SELECT id, title, description, priority, status, due_date, completed_at, created_at, updated_at
FROM task
WHERE id = ?;

-- name: ListTasks :many
SELECT id, title, description, priority, status, due_date, completed_at, created_at, updated_at
FROM task
ORDER BY priority DESC, created_at ASC;

-- name: ListTasksByStatus :many
SELECT id, title, description, priority, status, due_date, completed_at, created_at, updated_at
FROM task
WHERE status = ?
ORDER BY priority DESC, created_at ASC;

-- name: CreateTask :one
INSERT INTO task (title, description, priority, status, due_date, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, datetime('now'), datetime('now'))
RETURNING id, title, description, priority, status, due_date, completed_at, created_at, updated_at;

-- name: UpdateTask :one
UPDATE task
SET title = ?, description = ?, priority = ?, due_date = ?, updated_at = datetime('now')
WHERE id = ?
RETURNING id, title, description, priority, status, due_date, completed_at, created_at, updated_at;

-- name: UpdateTaskStatus :one
UPDATE task
SET status = ?, updated_at = datetime('now')
WHERE id = ?
RETURNING id, title, description, priority, status, due_date, completed_at, created_at, updated_at;

-- name: MarkTaskDone :one
UPDATE task
SET status = 'done', completed_at = datetime('now'), updated_at = datetime('now')
WHERE id = ?
RETURNING id, title, description, priority, status, due_date, completed_at, created_at, updated_at;

-- name: DeleteTask :exec
DELETE FROM task
WHERE id = ?;

-- name: CountTasksByStatus :one
SELECT COUNT(*) FROM task
WHERE status = ?;
