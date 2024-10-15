-- Migration up script for create-sessions-table
SET statement_timeout = 0;

--bun:split
CREATE SCHEMA IF NOT EXISTS auth;

--bun:split

CREATE TYPE auth_session_kind AS ENUM (
    'LoginKind', 
    'RefereshKind', 
    'LogoutKind', 
    'IddleKind'
);

CREATE TABLE auth.sessions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    ip VARCHAR(255) NOT NULL,
    user_agent VARCHAR(255) NOT NULL,
    kind auth_session_kind NOT NULL,
    retry INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);