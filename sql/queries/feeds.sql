-- name: AddFeed :one
INSERT INTO feeds (created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: ListFeeds :many
SELECT feeds.name AS feed_name, feeds.url AS feed_url, users.name AS user_name FROM feeds
JOIN users ON feeds.user_id = users.id;

-- name: GetFeedByID :one
SELECT * FROM feeds WHERE id = $1;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4
    )
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id
INNER JOIN users ON users.id = inserted_feed_follow.user_id;

-- name: GetFeedFollowsForUser :many
SELECT * FROM feed_follows WHERE user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2;

-- name: MarkFeedAsFetched :exec
UPDATE feeds 
SET 
    last_fetched_at = NOW(), 
    updated_at = NOW() 
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
WHERE last_fetched_at IS NULL
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;