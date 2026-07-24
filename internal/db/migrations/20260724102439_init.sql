-- +goose Up
CREATE TABLE tag (
    id            INTEGER PRIMARY KEY,
    name          TEXT NOT NULL UNIQUE,
    created_at    TEXT NOT NULL DEFAULT (datetime('now'))
);

CREATE TABLE task (
    id            INTEGER PRIMARY KEY,
    title         TEXT NOT NULL,
    description   TEXT,
    priority      INTEGER NOT NULL DEFAULT 2 CHECK(priority IN (1, 2, 3)),
    status        TEXT NOT NULL DEFAULT 'pending' CHECK(status IN ('pending', 'in_progress', 'done')),
    due_date      TEXT,
    completed_at  TEXT,
    created_at    TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at    TEXT NOT NULL DEFAULT (datetime('now'))
);

CREATE TABLE task_tag (
    task_id       INTEGER NOT NULL REFERENCES task(id) ON DELETE CASCADE,
    tag_id        INTEGER NOT NULL REFERENCES tag(id) ON DELETE CASCADE,
    PRIMARY KEY (task_id, tag_id)
);

CREATE INDEX idx_task_status_priority ON task(status, priority);
CREATE INDEX idx_task_tag_tag_id ON task_tag(tag_id);

-- +goose Down
DROP TABLE task_tag;
DROP TABLE task;
DROP TABLE tag;
