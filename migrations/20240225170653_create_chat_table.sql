-- +goose Up
CREATE TABLE chat
(
    id        BIGINT PRIMARY KEY,
    from_user VARCHAR(20),
    text      TEXT,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
drop table chat;