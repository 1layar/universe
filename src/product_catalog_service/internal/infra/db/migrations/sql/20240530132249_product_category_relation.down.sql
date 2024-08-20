-- Migration down script for product_category_relation
SET statement_timeout = 0;

--bun:split
DROP TABLE IF EXISTS product_catalog.product_category_relations;