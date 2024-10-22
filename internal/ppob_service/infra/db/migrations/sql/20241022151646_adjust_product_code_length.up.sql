-- Migration up script for adjust_product_code_length

-- bun:split
ALTER TABLE ppob.products ALTER COLUMN product_code TYPE VARCHAR(150);