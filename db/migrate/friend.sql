-- +migrate Up
CREATE DATABASE IF NOT EXISTS willers;
CREATE TABLE IF NOT EXISTS willers.friends (
  name  VARCHAR(255)	NOT NULL,
  other VARCHAR(255) NOT NULL,
  index(name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +migrate Down
