
-- +migrate Up
ALTER TABLE carts
ADD CONSTRAINT cart_user_fk
FOREIGN KEY (user_id) REFERENCES users (id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE carts
DROP CONSTRAINT cart_user_fk;