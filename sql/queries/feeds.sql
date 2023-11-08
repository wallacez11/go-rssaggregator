-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, update_at, name, url, user_id)
values ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
Select * from feeds;

-- name: GetNextFeedsToFatch :many
Select * from feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
update feeds
set last_fetched_at = NOW(),
update_at = NOW()
WHERE id = $1
RETURNING *;



