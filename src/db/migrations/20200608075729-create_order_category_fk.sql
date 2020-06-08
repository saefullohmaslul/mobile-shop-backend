
-- +migrate Up
ALTER TABLE orders
ADD CONSTRAINT order_category_fk
FOREIGN KEY (product_category_id) REFERENCES categories (id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE orders
DROP CONSTRAINT order_category_fk;