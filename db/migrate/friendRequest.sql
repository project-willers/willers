-- +migrate Up
CREATE DATABASE IF NOT EXISTS willers;
CREATE TABLE IF NOT EXISTS willers.friendrequest (
  name  VARCHAR(255)	NOT NULL,
  other VARCHAR(255) NOT NULL,
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +migrate Down
