
-- +migrate Up
CREATE TABLE cart_products
(
  id UUID NOT NULL,
  cart_id UUID NOT NULL,
  product_id UUID NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);
-- +migrate Down
DROP TABLE cart_products;