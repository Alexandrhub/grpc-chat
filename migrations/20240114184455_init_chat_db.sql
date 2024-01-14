-- +goose Up
CREATE TABLE chats (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    msg TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS chats;
