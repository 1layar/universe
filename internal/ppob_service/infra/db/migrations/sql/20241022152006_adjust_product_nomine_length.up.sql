-- Migration up script for adjust_product_nomine_length

-- bun:split
ALTER TABLE ppob.products ALTER COLUMN product_nominal TYPE VARCHAR(225);