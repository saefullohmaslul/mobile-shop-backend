
-- +migrate Up
CREATE TABLE products
(
  id UUID NOT NULL,
  thumbnail TEXT,
  title VARCHAR(50) NOT NULL,
  description TEXT NOT NULL,
  price FLOAT NOT NULL,
  product_category_id UUID NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE products;