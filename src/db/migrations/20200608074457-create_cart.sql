
-- +migrate Up
CREATE TABLE carts
(
  id UUID NOT NULL,
  user_id UUID NOT NULL,
  amount INTEGER NOT NULL,
  sub_total FLOAT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);
-- +migrate Down
DROP TABLE carts;