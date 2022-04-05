-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE passwords (
  user_uuid uuid,
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE,
  deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE UNIQUE INDEX one_active_password ON passwords (user_uuid) WHERE deleted_at IS NULL; 

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE passwords;