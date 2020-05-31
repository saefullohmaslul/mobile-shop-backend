
-- +migrate Up
CREATE TABLE auth_informations
(
  id UUID NOT NULL,
  user_id UUID NOT NULL,
  refresh_token VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT user_id_unique UNIQUE (user_id)
);

-- +migrate Down
DROP TABLE auth_informations;