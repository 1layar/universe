-- Migration up script for product_types

-- bun:split
CREATE SCHEMA IF NOT EXISTS ppob;

-- bun:split
CREATE TABLE ppob.product_types (
    id SERIAL PRIMARY KEY,
    type_name VARCHAR(50) NOT NULL UNIQUE
);