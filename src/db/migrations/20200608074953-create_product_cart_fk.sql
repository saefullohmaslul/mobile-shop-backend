
-- +migrate Up
ALTER TABLE cart_products
ADD CONSTRAINT product_cart_fk
FOREIGN KEY (cart_id) REFERENCES carts (id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE cart_products
DROP CONSTRAINT product_cart_fk;
