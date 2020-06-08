
-- +migrate Up
CREATE TABLE users
(
  id UUID CONSTRAINT user_pk PRIMARY KEY,
  username VARCHAR(50) NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(100),
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP,
  CONSTRAINT username_unique UNIQUE (username),
  CONSTRAINT email_unique UNIQUE (email)
);

-- +migrate Down
DROP TABLE users;