-- +goose Up
CREATE TABLE feeds(
    id			uuid PRIMARY KEY,
    created_at	TIMESTAMP	NOT NULL,
    updated_at	TIMESTAMP	NOT NULL,
    name		text		NOT NULL,
    url			text		NOT NULL UNIQUE,
    user_id		uuid		NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
