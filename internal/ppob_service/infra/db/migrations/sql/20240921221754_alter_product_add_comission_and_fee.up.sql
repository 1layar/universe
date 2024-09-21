-- Migration up script for alter_product_add_comission_and_fee

-- bun:split
ALTER TABLE ppob.products ADD COLUMN IF NOT EXISTS comission DECIMAL(10, 2) NOT NULL DEFAULT 0;
ALTER TABLE ppob.products ADD COLUMN IF NOT EXISTS fee DECIMAL(10, 2) NOT NULL DEFAULT 0;