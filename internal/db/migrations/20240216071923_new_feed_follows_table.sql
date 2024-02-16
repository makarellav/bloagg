-- +goose Up
-- +goose StatementBegin
CREATE TABLE feed_follows
(
    id         UUID PRIMARY KEY,
    user_id    UUID REFERENCES users (id) ON DELETE CASCADE,
    feed_id    UUID REFERENCES feeds (id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    UNIQUE (user_id, feed_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feed_follows;
-- +goose StatementEnd
