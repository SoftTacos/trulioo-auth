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

CREATE TABLE refresh_tokens (
  user_uuid uuid NOT NULL,
  header TEXT NOT NULL,
  payload TEXT NOT NULL,
  signature TEXT NOT NULL,
  expires_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE INDEX refresh_token_user_uuid_index ON refresh_tokens (user_uuid);
CREATE INDEX refresh_token_header_index ON refresh_tokens (header);
CREATE INDEX refresh_token_payload_index ON refresh_tokens (payload);
CREATE INDEX refresh_token_signature_index ON refresh_tokens (signature);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE passwords;