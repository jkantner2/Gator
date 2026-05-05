-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8
)
RETURNING *;

-- name: GetPostsForUser :many
WITH selected_feed_follow AS (
    SELECT * FROM feed_follows
    WHERE feed_follows.user_id = $1
)
SELECT 
posts.*
FROM selected_feed_follow
INNER JOIN posts ON selected_feed_follow.feed_id = posts.feed_id
ORDER BY posts.published_at DESC
LIMIT $2;
