
-- +migrate Up
ALTER TABLE products
ADD CONSTRAINT product_category_fk
FOREIGN KEY (product_category_id) REFERENCES categories (id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE products
DROP CONSTRAINT product_category_fk;