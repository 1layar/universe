-- Migration down script for create-sessions-table
SET statement_timeout = 0;

--bun:split
DROP TABLE IF EXISTS auth.sessions;
DROP TYPE IF EXISTS auth_session_kind;