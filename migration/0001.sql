CREATE TABLE todos (
    todo_id     UUID          PRIMARY KEY,
    task        VARCHAR(166)  NOT NULL,
    is_completed BOOL         DEFAULT false,
    created_at  TIMESTAMP     DEFAULT current_timestamp,
    updated_at  TIMESTAMP     DEFAULT NULL,
    deleted_at  TIMESTAMP     DEFAULT NULL
);

