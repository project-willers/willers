-- +migrate Up
CREATE DATABASE IF NOT EXISTS willers;
CREATE TABLE IF NOT EXISTS willers.userInfo (
  name  VARCHAR(255)	NOT NULL,
  content VARCHAR(200)  NOT NULL,
  select_at DATETIME 	NOT NULL ,
  updated_at  DATETIME 	NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (name),
  UNIQUE (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +migrate Down
