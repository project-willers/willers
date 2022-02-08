-- +migrate Up
CREATE DATABASE IF NOT EXISTS willers;
CREATE TABLE IF NOT EXISTS willers.friends (
  id  BIGINT  NOT NULL  AUTO_INCREMENT,
  name  VARCHAR(255)	NOT NULL,
  other VARCHAR(255) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (name),
  UNIQUE (other)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +migrate Down
