CREATE DATABASE soccer;
USE soccer;

DROP TABLE IF EXISTS players;
CREATE TABLE players (
  player_id INT KEY AUTO_INCREMENT,
  first_name VARCHAR(20) NOT NULL,
  last_name VARCHAR(20) NOT NULL,
  birth_date DATE,
  gender VARCHAR(20) NOT NULL,
  phone_number VARCHAR(20) NOT NULL,
  email_address VARCHAR(50)NOT NULL,
  jerse_number int(11),
  team VARCHAR(20),
  address_1 VARCHAR(50) NOT NULL,
  address_2 VARCHAR(20),
  city VARCHAR(20) NOT NULL,
  state VARCHAR(20) NOT NULL,
  zip_code VARCHAR(20) NOT NULL
);
INSERT INTO players VALUES
  ("1001","Harry","Kane","11-11-11","Male","508-320-7922","harryK@google.com","10","Tottenham","44 Hickery Lyne","2","Burlington","MA","1803"),
  ("1002","Christiano","Rolando","05-12-09","Male","508-652-7975","christianoR@gmail.com","7","Man United","39 Hiltop DR","","Burlington","MA","1803"),
  ("1003","Neymar","da Silva","09-12-09","Male","987-652-7975","neymarda@yahoo.com","11","PSG","12 Sandybrook DR","","Burlington","MA","1803"),
  ("1004","Lionel","Messi","01-05-05","Male","987-652-8544","lionelmessi@yahoo.com","10","FC Barcelona","85 Church hill DR","","Lexington","MA","1803"),
  ("1005","Paul","Pogba","09-12-09","Male","987-652-7975","paulP@yahoo.com","11","Man United","12 Sandybrook DR","","Burlington","MA","1803");

DROP TABLE IF EXISTS teams;
CREATE TABLE teams (
    team_id INT KEY AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL,
    address_1 VARCHAR(50) NOT NULL,
    address_2 VARCHAR(20),
    city VARCHAR(20) NOT NULL,
    state VARCHAR(20) NOT NULL,
    zip_code INT NOT NULL
  );
INSERT INTO teams VALUES
    ("1001","Man United","44 Hickery Lyne","","Woburn","MA","01804"),
    ("1002","FC Barcelona","44 Hickery Lyne","","Burlington","MA","01803"),
    ("1003","PSG","15 Circle Lyne","","Burlington","MA","01802"),
    ("1004","Tottenham","44 Hickery Lyne","","Belmont","MA","01804");

DROP TABLE IF EXISTS coaches;
CREATE TABLE coaches (
  id INT KEY AUTO_INCREMENT,
  first_name VARCHAR(20) NOT NULL,
  last_name VARCHAR(20) NOT NULL,
  gender VARCHAR(20) NOT NULL,
  phone_number VARCHAR(20) NOT NULL,
  email_address VARCHAR(50)NOT NULL,
  address_1 VARCHAR(50) NOT NULL,
  address_2 VARCHAR(20),
  city VARCHAR(20) NOT NULL,
  state VARCHAR(20) NOT NULL,
  zip_code INT NOT NULL,
  role VARCHAR(20) NOT NULL,
  team VARCHAR(20) NOT NULL
  );
INSERT INTO coaches VALUES
    ("1001","Ole","Gunnar","M","781-112-7922","OleG@yahoo.com","408 Winn Street","","Woburn","MA","01803","Head Coach","Man United"),
    ("1002","Xavi","hernandez","M","227-342-7922","XaviH@yahoo.com","408 Winn Street","","Woburn","MA","01803","Assistant Coach","FC Barcelona"),
    ("1003","Mauricio","Pochettino","M","227-342-7922","MauricioP@gmail.com","408 Winn Street","","Woburn","MA","01803","Head Coach","PSG"),
    ("1004","Antonio","Conte","M","227-342-7922","AntonioA@gmail.com","408 Winn Street","","Woburn","MA","01803","Head Coach","Tottenham");

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `username` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  `role` varchar(20) NOT NULL,
  `player_id` int(11) DEFAULT NULL,
  `team_id` int(11) DEFAULT NULL,
  `created_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
INSERT INTO `users` VALUES
  ('admin','abc123','admin',NULL,NULL,'2021-08-09 09:27:22'),
  ('2001','abc123','user', 1001, 1001,'2021-05-09 10:27:22'),
  ('2002','abc123','user', 1002, 1002,'2021-03-09 01:27:22'),
  ('2003','abc123','user', 1003, 1003,'2021-09-09 08:27:22'),
  ('2004','abc123','user', 1004, 1004,'2020-01-09 03:27:22');

DROP TABLE IF EXISTS `refresh_token_store`;
CREATE TABLE `refresh_token_store` (
    `refresh_token` varchar(300) NOT NULL,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`refresh_token`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
