-- Migration up script for add-confirm-id-and-date
-- Migration down script for add-confirm-id-and-date
SET statement_timeout = 0;

--bun:split

ALTER TABLE account.users ADD COLUMN confirm_id VARCHAR(255);
ALTER TABLE account.users ADD COLUMN confirm_date TIMESTAMP;

--bun:split

-- add unique index for confirm_id
CREATE UNIQUE INDEX users_confirm_id_uindex
ON account.users (confirm_id)
WHERE deleted_at IS NULL AND confirm_id IS NOT NULL;