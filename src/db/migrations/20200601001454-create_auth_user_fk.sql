
-- +migrate Up
ALTER TABLE auth_informations 
ADD CONSTRAINT auth_user_fk
FOREIGN KEY (user_id) REFERENCES users (id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE auth_informations
DROP CONSTRAINT auth_user_fk;