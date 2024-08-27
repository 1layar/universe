-- Migration up script for products

-- bun:split
CREATE SCHEMA IF NOT EXISTS ppob;

-- bun:split
CREATE TABLE ppob.products (
    id SERIAL PRIMARY KEY,
    kind VARCHAR(50) NOT NULL CHECK (kind IN ('prepaid', 'postpaid')),
    product_code VARCHAR(50) NOT NULL,
    product_description TEXT NOT NULL,
    product_nominal VARCHAR(50) NOT NULL,
    product_details TEXT DEFAULT '-',
    product_price DECIMAL(10, 2) NOT NULL,
    product_type_id INTEGER NOT NULL REFERENCES product_types(id),
    active_period INTEGER NOT NULL DEFAULT 0,
    status VARCHAR(20) NOT NULL CHECK (status IN ('active', 'inactive')),
    icon_url TEXT DEFAULT '-',
    product_category_id INTEGER NOT NULL REFERENCES product_categories(id),
    billing_cycle INTEGER,            -- New column to support postpaid products
    due_date DATE,                    -- New column for the due date of payment
    grace_period INTEGER              -- New column for the grace period in days
);