-- +goose Up
CREATE TABLE posts(
    id				uuid PRIMARY KEY,
    created_at 		TIMESTAMP	NOT NULL,
    updated_at		TIMESTAMP	NOT NULL,
    title			text,
    url				text		NOT NULL UNIQUE,
    description 	text,
    published_at	TIMESTAMP,
    feed_id			uuid		NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
