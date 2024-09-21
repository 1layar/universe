-- Migration up script for product_categories

-- bun:split
CREATE SCHEMA IF NOT EXISTS ppob;

-- bun:split
CREATE TABLE ppob.product_categories (
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(225) NOT NULL UNIQUE
);