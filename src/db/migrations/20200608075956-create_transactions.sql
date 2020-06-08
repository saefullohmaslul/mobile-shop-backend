
-- +migrate Up
CREATE TABLE transactions
(
  id UUID NOT NULL,
  user_id UUID NOT NULL,
  total FLOAT NOT NULL,
  status VARCHAR(10) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);
-- +migrate Down
DROP TABLE transactions;