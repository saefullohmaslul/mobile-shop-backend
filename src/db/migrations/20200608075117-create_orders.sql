
-- +migrate Up
CREATE TABLE orders
(
  id UUID NOT NULL,
  user_id UUID NOT NULL,
  product_id UUID NOT NULL,
  transaction_id UUID NOT NULL,
  thumbnail VARCHAR(30),
  title VARCHAR(50) NOT NULL,
  description TEXT NOT NULL,
  price FLOAT NOT NULL,
  product_category_id UUID NOT NULL,
  amount INTEGER NOT NULL,
  sub_total FLOAT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);
-- +migrate Down
DROP TABLE orders;