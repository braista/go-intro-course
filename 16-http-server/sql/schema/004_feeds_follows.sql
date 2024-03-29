-- +goose Up
CREATE TABLE feeds_follows (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,
  UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feeds_follows;