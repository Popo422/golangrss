-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE  /* references means it must be relational you cant create a entry if that user_id 
    does not exist in users , On delete cascade basically means  once a user is deleted it deletes the feed alongside with it*/

);




-- +goose Down
DROP TABLE feeds;