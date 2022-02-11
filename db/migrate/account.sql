-- +migrate Up
CREATE DATABASE IF NOT EXISTS willers;
CREATE TABLE IF NOT EXISTS willers.accounts (
  name		VARCHAR(50)	NOT NULL,
  email VARCHAR(50) NOT NULL,
  password	VARCHAR(513) 	NOT NULL,
  PRIMARY KEY (name),
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +migrate Down
