-- Migration up script for transaction

-- bun:split
CREATE SCHEMA IF NOT EXISTS payment;

-- bun:split
DROP TYPE IF EXISTS payment.transaction_type;
CREATE TYPE payment.transaction_type AS ENUM ('sell', 'refund', 'buy');

-- bun:split
DROP TYPE IF EXISTS payment.transaction_status;
CREATE TYPE payment.transaction_status AS ENUM ('pending', 'success', 'failed');

-- bun:split
CREATE TABLE payment.transactions (
    id SERIAL PRIMARY KEY,
    reference_id VARCHAR(36) NOT NULL UNIQUE,
    payment_id VARCHAR(36) NOT NULL UNIQUE,
    channel_id INTEGER NOT NULL REFERENCES payment.channels(id),
    user_id VARCHAR(36) NOT NULL,
    order_code INT NOT NULL,
    origin VARCHAR(255) NOT NULL,
    type payment.transaction_type NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    status payment.transaction_status NOT NULL,
    note TEXT,
    metadata JSON,
    from_account INTEGER NOT NULL REFERENCES payment.accounts(id),
    to_account INTEGER NOT NULL REFERENCES payment.accounts(id),
    fee DECIMAL(10, 2) NOT NULL,
    parent_id INTEGER REFERENCES payment.transactions(id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
)