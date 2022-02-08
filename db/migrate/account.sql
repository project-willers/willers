-- +migrate Up
CREATE DATABASE IF NOT EXISTS willers;
CREATE TABLE IF NOT EXISTS willers.accounts (
  name		VARCHAR(50)	NOT NULL,
  email VARCHAR(50) NOT NULL,
  password	VARCHAR(255) 	NOT NULL,
  PRIMARY KEY (name),
  UNIQUE (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +migrate Down
