-- name: CreateFeed :one 
insert into feeds (id, created_at, updated_at, name, url, user_id) values (
    $1, 
    $2, 
    $3, 
    $4, 
    $5, 
    $6
) returning *; 

-- name: GetFeeds :many 
select feeds.name, url, users.name as username from feeds INNER JOIN users 
on feeds.user_id = users.id;

-- name: CreateFeedFollow :one 
with inserted_feed_follow as (
    insert into feeds_follow (created_at, updated_at, user_id, feed_id) values (
        $1, 
        $2, 
        $3, 
        $4
    ) returning *
)
select inserted_feed_follow.*, users.name, feeds.name from inserted_feed_follow
INNER JOIN users on inserted_feed_follow.user_id = users.id 
INNER JOIN feeds on inserted_feed_follow.feed_id = feeds.id; 

-- name: LookupFeed :one 
select * FROM feeds where url = $1; 

-- name: GetFeedsFollow :many 
select feeds.name as feed_name from feeds_follow INNER JOIN feeds on feeds.id = feeds_follow.feed_id where feeds_follow.user_id = $1; 

-- name: UnfollowFeed :exec
delete from feeds_follow where feeds_follow.user_id=$1 and feeds_follow.feed_id = (select id from feeds where url=$2); 

