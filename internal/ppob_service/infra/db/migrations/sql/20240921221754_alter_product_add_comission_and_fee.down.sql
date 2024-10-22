-- Migration down script for alter_product_add_comission_and_fee

-- bun:split
ALTER TABLE ppob.products DROP COLUMN IF EXISTS comission;
ALTER TABLE ppob.products DROP COLUMN IF EXISTS fee