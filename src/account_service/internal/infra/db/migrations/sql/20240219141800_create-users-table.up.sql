SET statement_timeout = 0;

--bun:split
CREATE SCHEMA IF NOT EXISTS account;

--bun:split

CREATE TABLE account.users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

--bun:split

CREATE UNIQUE INDEX users_id_uindex ON account.users (id)
WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX users_email_uindex ON account.users (email)
WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX users_username_uindex ON account.users (username)
WHERE deleted_at IS NULL;

CREATE INDEX users_created_at_index ON account.users (created_at);
