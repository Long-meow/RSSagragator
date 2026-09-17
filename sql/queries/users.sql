-- name: CreateUser :one
insert into users (id, created_at, updated_at, name) 
values (
    $1, 
    $2, 
    $3, 
    $4
)
returning *; 

-- name: GetUser :one 
SELECT * FROM users WHERE name=$1; 

-- name: RemoveUser :one 
delete from users where name=$1 returning *; 

-- name: ResetUsers :exec 
delete from users; 

-- name: GetUsers :many
select name from users;