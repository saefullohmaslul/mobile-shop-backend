
-- +migrate Up
ALTER TABLE orders
ADD CONSTRAINT order_product_fk
FOREIGN KEY (product_id) REFERENCES products (id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE orders
DROP CONSTRAINT order_product_fk;