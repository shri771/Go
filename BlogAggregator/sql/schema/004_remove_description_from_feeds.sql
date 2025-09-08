-- +goose Up
ALTER TABLE feeds
DROP COLUMN user_id;

-- +goose Down
ALTER TABLE feeds
ADD COLUMN user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE;
