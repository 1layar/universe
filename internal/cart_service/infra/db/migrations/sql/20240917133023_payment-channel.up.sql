-- Migration up script for payment_channel

-- bun:split
CREATE SCHEMA IF NOT EXISTS payment;

-- bun:split
DROP TYPE IF EXISTS payment.fee_type;
CREATE TYPE payment.fee_type AS ENUM ('flat', 'percentage');

-- bun:split
CREATE TABLE IF NOT EXISTS payment.channels (
    id SERIAL PRIMARY KEY,
    channel_code VARCHAR(50) NOT NULL,
    channel_method VARCHAR(50) NOT NULL,
    channel_name VARCHAR(50) NOT NULL,
    active BOOLEAN NOT NULL,
    fee DECIMAL(10, 2) NOT NULL,
    fee_type payment.fee_type NOT NULL,
    additional_fee DECIMAL(10, 2)
)