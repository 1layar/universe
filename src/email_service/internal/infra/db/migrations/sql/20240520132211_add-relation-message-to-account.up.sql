-- Migration up script for add-relation-message-to-account
SET statement_timeout = 0;

--bun:split
ALTER TABLE email.messages ADD COLUMN "account_id" INT;

--bun:split
ALTER TABLE email.messages ADD CONSTRAINT "messages_account_id_fkey" 
    FOREIGN KEY ("account_id") REFERENCES email.accounts ("id")
    ON DELETE CASCADE ON UPDATE CASCADE;