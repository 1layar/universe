-- Migration up script for product
SET statement_timeout = 0;

-- bun:split
CREATE SCHEMA IF NOT EXISTS product_catalog;

--bun:split
CREATE TABLE product_catalog.products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    status TINYINT NOT NULL,
    fee MONEY NOT NULL,
    komisi MONEY NOT NULL,
    sku VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    picture_url VARCHAR(255),
    price MONEY NOT NULL,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);