CREATE DATABASE qonocloud;
use qonocloud;

CREATE TABLE `player`(
    `player_id` INT UNSIGNED AUTO_INCREMENT,
    `player_username` VARCHAR(32) NOT NULL,
    `player_displayname` VARCHAR(32) NOT NULL,
    PRIMARY KEY (`player_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
