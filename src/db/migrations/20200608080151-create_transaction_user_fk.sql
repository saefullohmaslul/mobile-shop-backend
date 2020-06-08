
-- +migrate Up
ALTER TABLE transactions
ADD CONSTRAINT transaction_user_fk
FOREIGN KEY (user_id) REFERENCES users (id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE transactions
DROP CONSTRAINT transaction_user_fk;