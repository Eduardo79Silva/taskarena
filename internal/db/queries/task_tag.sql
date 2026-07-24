-- name: AddTagToTask :exec
INSERT INTO task_tag (task_id, tag_id)
VALUES (?, ?)
ON CONFLICT DO NOTHING;

-- name: RemoveTagFromTask :exec
DELETE FROM task_tag
WHERE task_id = ? AND tag_id = ?;

-- name: ListTagsForTask :many
SELECT t.id, t.name, t.created_at
FROM tag t
JOIN task_tag tt ON tt.tag_id = t.id
WHERE tt.task_id = ?
ORDER BY t.name ASC;

-- name: ListTasksForTag :many
SELECT t.id, t.title, t.description, t.priority, t.status, t.due_date, t.completed_at, t.created_at, t.updated_at
FROM task t
JOIN task_tag tt ON tt.task_id = t.id
WHERE tt.tag_id = ?
ORDER BY t.priority DESC, t.created_at ASC;
