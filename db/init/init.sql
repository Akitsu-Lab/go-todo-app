DROP DATABASE IF EXISTS `todo`;
CREATE DATABASE `todo`;

USE `todo`;

DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks` (
    id INT NOT NULL ,
    name VARCHAR(80) NOT NULL ,
    status INT NOT NULL ,
    PRIMARY KEY (id),
    CHECK (status IN (0, 1, 2))
);

INSERT INTO `tasks` VALUES (1,'掃除をする',0);
INSERT INTO `tasks` VALUES (2,'勉強をする',1);
INSERT INTO `tasks` VALUES (3,'買い物をする',2);
INSERT INTO `tasks` VALUES (4,'本を読む',0);