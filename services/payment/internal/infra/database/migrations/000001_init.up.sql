-- -----------------------------------------------------
-- Schema commerce_payment
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `commerce_payment` DEFAULT CHARACTER SET utf8 ;
-- -----------------------------------------------------
-- Schema new_schema1
-- -----------------------------------------------------
USE `commerce_payment` ;
CREATE TABLE IF NOT EXISTS `commerce_payment`.`billing_addresses` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `street` VARCHAR(45) NULL,
  `number` VARCHAR(45) NULL,
  `zip_code` VARCHAR(45) NULL,
  `city` VARCHAR(45) NULL,
  `country` VARCHAR(45) NULL,
  `state` VARCHAR(45) NULL,
  `user_id` INT NOT NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_addresses_users1_idx` (`user_id` ASC) VISIBLE,
  )
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `commerce_payment`.`cards`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `commerce_payment`.`cards` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `card_name` VARCHAR(45) NULL,
  `number` VARCHAR(45) NULL,
  `holder_name` VARCHAR(45) NULL,
  `expirity_month` INT(2) NULL,
  `expirity_year` INT(4) NULL,
  `cvv` INT(3) NULL,
  `user_id` INT NOT NULL,
  `billing_address_id` INT NOT NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`, `user_id`),
  INDEX `fk_cards_users_idx` (`user_id` ASC) VISIBLE,
  INDEX `fk_cards_billing_addresses_idx` (`billing_address_id` ASC) VISIBLE,
  CONSTRAINT `fk_cards_addresses`
    FOREIGN KEY (`billing_address_id`)
    REFERENCES `commerce_payment`.`addresses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `commerce_payment`.`payments`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `commerce_payment`.`payments` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `type` VARCHAR(45) NULL,
  `installments` INT(2) NULL,
  `order_id` INT NOT NULL,
  `card_id` INT NOT NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`, `order_id`, `card_id`),
  INDEX `fk_payments_orders1_idx` (`order_id` ASC) VISIBLE,
  INDEX `fk_payments_cards1_idx` (`card_id` ASC) VISIBLE,
  CONSTRAINT `fk_payments_orders1`
    FOREIGN KEY (`order_id`)
    REFERENCES `commerce_payment`.`orders` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_payments_cards1`
    FOREIGN KEY (`card_id`)
    REFERENCES `commerce_payment`.`cards` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;