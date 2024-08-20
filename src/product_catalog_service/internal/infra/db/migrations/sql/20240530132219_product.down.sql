-- Migration down script for product
SET statement_timeout = 0;

--bun:split
DROP TABLE product_catalog.products;