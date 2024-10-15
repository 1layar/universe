-- Migration up script for product_category_relations
SET statement_timeout = 0;

--bun:split
CREATE TABLE product_catalog.product_category_relations (
    product_id INT NOT NULL REFERENCES product_catalog.products(id) ON DELETE CASCADE,
    category_id INT NOT NULL REFERENCES product_catalog.categories(id) ON DELETE CASCADE
);

--bun:split
CREATE UNIQUE INDEX product_category_relations_pkey ON product_catalog.product_category_relations (product_id, category_id);