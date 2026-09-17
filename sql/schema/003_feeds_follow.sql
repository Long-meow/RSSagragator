-- +goose Up 
create table feeds_follow (
    id INT generated always as identity primary key, 
    created_at timestamp not null, 
    updated_at timestamp not null, 
    user_id uuid not null, 
    feed_id uuid not null, 
    constraint fk_user  
        foreign key (user_id) 
        references users (id) 
        on delete cascade, 
    constraint fk_feed_follow
        foreign key (feed_id)
        references feeds (id)
        on delete cascade,
    constraint fk_feed_user 
        unique (user_id, feed_id) 
); 

-- +goose Down 
DROP TABLE feeds_follow; 