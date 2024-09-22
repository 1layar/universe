-- Migration up script for account

-- bun:split
CREATE SCHEMA IF NOT EXISTS payment;

-- bun:split
DROP TYPE IF EXISTS payment.account_type;
CREATE TYPE payment.account_type AS ENUM ('master', 'member', 'guest', 'vendor');

-- bun:split
CREATE TABLE IF NOT EXISTS payment.accounts (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(36) UNIQUE,
    name VARCHAR(50) NOT NULL,
    account_type payment.account_type NOT NULL,
    wallet_address VARCHAR(36) UNIQUE NOT NULL,
    balance DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
)