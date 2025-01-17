-- MySQL Script generated by MySQL Workbench
-- Sat 25 May 2024 02:57:19 PM JST
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `goweb` DEFAULT CHARACTER SET utf8 ;
USE `goweb` ;

-- -----------------------------------------------------
-- Table `items`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `goweb`.`items` ;

CREATE TABLE IF NOT EXISTS `goweb`.`items` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME NULL,
  `name` VARCHAR(40) NOT NULL,
  `price` FLOAT NOT NULL,
  `quantity` INT UNSIGNED NOT NULL,
  `onSale` TINYINT NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `stores`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `goweb`.`stores` ;

CREATE TABLE IF NOT EXISTS `goweb`.`stores` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME NULL,
  `name` VARCHAR(128) NOT NULL,
  `owner` ENUM('state', 'private') NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;

-- -----------------------------------------------------
-- Table `stores`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `goweb`.`stores_items` ;

CREATE TABLE IF NOT EXISTS `goweb`.`stores_items` (
  `store_id` int unsigned NOT NULL,
  `item_id` int unsigned NOT NULL,
  UNIQUE KEY `uq_store_item_combination` (`item_id`,`store_id`),
  KEY `fk_store_id_idx` (`store_id`),
  KEY `fk_item_id_idx` (`item_id`),
  CONSTRAINT `fk_item_id` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `fk_store_id` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

INSERT INTO `goweb`.`stores`
(`created_at`,
`updated_at`,
`deleted_at`,
`name`,
`owner`)
VALUES
(now(),
now(),
null,
'Cold Trading Ltd.',
'private');

INSERT INTO `goweb`.`stores`
(`created_at`,
`updated_at`,
`deleted_at`,
`name`,
`owner`)
VALUES
(now(),
now(),
null,
'Cheap Union',
'state');

INSERT INTO `goweb`.`items`
(`created_at`,
`updated_at`,
`name`,
`price`,
`quantity`,
`onSale`)
VALUES
(now(),
now(),
'broom',
3.50,
143,
false);

INSERT INTO `goweb`.`items`
(`created_at`,
`updated_at`,
`name`,
`price`,
`quantity`,
`onSale`)
VALUES
(now(),
now(),
'bread',
1.10,
76,
true);

INSERT INTO `goweb`.`items`
(`created_at`,
`updated_at`,
`name`,
`price`,
`quantity`,
`onSale`)
VALUES
(now(),
now(),
'canned beef',
2.0,
9,
false);

INSERT INTO `goweb`.`stores_items`
(`store_id`,
`item_id`)
VALUES
(1,
1);

INSERT INTO `goweb`.`stores_items`
(`store_id`,
`item_id`)
VALUES
(2,
1);

INSERT INTO `goweb`.`stores_items`
(`store_id`,
`item_id`)
VALUES
(2,
2);
