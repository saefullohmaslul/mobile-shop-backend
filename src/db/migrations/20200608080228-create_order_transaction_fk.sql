
-- +migrate Up
ALTER TABLE orders
ADD CONSTRAINT order_transaction_fk
FOREIGN KEY (transaction_id) REFERENCES transactions (id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE orders
DROP CONSTRAINT order_transaction_fk;