-- Migration up script for category
SET statement_timeout = 0;

-- bun:split
CREATE TABLE product_catalog.categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    picture_url VARCHAR(255),
    parent_id INT DEFAULT NULL REFERENCES product_catalog.categories(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
