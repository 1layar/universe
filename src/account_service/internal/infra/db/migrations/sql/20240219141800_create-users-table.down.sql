SET statement_timeout = 0;

--bun:split

delete table account.users;

--bun:split

delete index users_id_uindex;
delete index users_email_uindex;
delete index users_username_uindex;
delete index users_created_at_index;