
-- +migrate Up
ALTER TABLE cart_products
ADD CONSTRAINT cart_product_fk
FOREIGN KEY (product_id) REFERENCES products (id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE cart_products
DROP CONSTRAINT cart_product_fk;
