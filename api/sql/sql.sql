CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
  id int auto_increment PRIMARY KEY,
  name varchar(100) not null,
  nick varchar(50) not null UNIQUE,
  email varchar(50) not null UNIQUE,
  password varchar(50) not null UNIQUE,
  create_at TIMESTAMP default CURRENT_TIMESTAMP()

) ENGINE=INNODB;