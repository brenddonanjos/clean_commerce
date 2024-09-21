-- -----------------------------------------------------
-- Schema commerce_user
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `commerce_user` DEFAULT CHARACTER SET utf8 ;
-- -----------------------------------------------------
-- Schema new_schema1
-- -----------------------------------------------------
USE `commerce_user` ;

-- -----------------------------------------------------
-- Table `commerce_user`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `commerce_user`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `username` VARCHAR(45) NOT NULL,
  `email` VARCHAR(45) NOT NULL,
  `password` TEXT NOT NULL,
  `birth_date` DATE NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;
