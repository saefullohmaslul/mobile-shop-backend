
-- +migrate Up
ALTER TABLE orders
ADD CONSTRAINT order_user_fk
FOREIGN KEY (user_id) REFERENCES users (id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE orders
DROP CONSTRAINT order_user_fk;