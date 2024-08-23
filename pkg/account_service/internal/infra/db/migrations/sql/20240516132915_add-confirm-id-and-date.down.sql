-- Migration down script for add-confirm-id-and-date
SET statement_timeout = 0;

--bun:split

ALTER TABLE account.users DROP COLUMN confirm_id;
ALTER TABLE account.users DROP COLUMN confirm_date;

--bun:split

-- drop unique index for confirm_id
DROP INDEX IF EXISTS users_confirm_id_uindex;