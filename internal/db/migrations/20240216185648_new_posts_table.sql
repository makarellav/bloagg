-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts
(
    id           UUID PRIMARY KEY,
    title        TEXT      NOT NULL,
    url          TEXT      NOT NULL UNIQUE,
    description  TEXT,
    published_at TIMESTAMP NOT NULL,
    created_at   TIMESTAMP NOT NULL,
    updated_at   TIMESTAMP NOT NULL,
    feed_id      UUID REFERENCES feeds (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
