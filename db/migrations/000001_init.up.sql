-- -----------------------------------------------------
-- Schema commerce_ai
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `commerce_ai` DEFAULT CHARACTER SET utf8 ;
-- -----------------------------------------------------
-- Schema new_schema1
-- -----------------------------------------------------
USE `commerce_ai` ;

-- -----------------------------------------------------
-- Table `commerce_ai`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `commerce_ai`.`users` (
  `id` INT NOT NULL COMMENT '	',
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

-- -----------------------------------------------------
-- Table `commerce_ai`.`products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `commerce_ai`.`products` (
  `id` INT NOT NULL,
  `title` VARCHAR(45) NOT NULL,
  `description` VARCHAR(45) NULL,
  `barcode` TEXT NOT NULL,
  `price` DECIMAL(13,2) NOT NULL,
  `amount` INT(42) NOT NULL DEFAULT 0,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `commerce_ai`.`addresses`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `commerce_ai`.`addresses` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `street` VARCHAR(45) NULL,
  `number` VARCHAR(45) NULL,
  `zip_code` VARCHAR(45) NULL,
  `city` VARCHAR(45) NULL,
  `country` VARCHAR(45) NULL,
  `state` VARCHAR(45) NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `commerce_ai`.`cards`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `commerce_ai`.`cards` (
  `id` INT NOT NULL,
  `card_name` VARCHAR(45) NULL,
  `number` VARCHAR(45) NULL,
  `holder_name` VARCHAR(45) NULL,
  `expirity_month` INT(2) NULL,
  `expirity_year` INT(4) NULL,
  `cvv` INT(3) NULL,
  `user_id` INT NOT NULL,
  `address_id` INT NOT NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`, `user_id`),
  INDEX `fk_cards_users_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_cards_users`
    FOREIGN KEY (`user_id`)
    REFERENCES `commerce_ai`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_cards_addresses`
    FOREIGN KEY (`address_id`)
    REFERENCES `commerce_ai`.`addresses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `commerce_ai`.`categories`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `commerce_ai`.`categories` (
  `id` INT NOT NULL,
  `title` VARCHAR(45) NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `commerce_ai`.`products_categories`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `commerce_ai`.`products_categories` (
  `id` INT NOT NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  `product_id` INT NOT NULL,
  `category_id` INT NOT NULL,
  PRIMARY KEY (`id`, `product_id`, `category_id`),
  INDEX `fk_products_categories_products1_idx` (`product_id` ASC) VISIBLE,
  INDEX `fk_products_categories_categories1_idx` (`category_id` ASC) VISIBLE,
  CONSTRAINT `fk_products_categories_products1`
    FOREIGN KEY (`product_id`)
    REFERENCES `commerce_ai`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_products_categories_categories1`
    FOREIGN KEY (`category_id`)
    REFERENCES `commerce_ai`.`categories` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `commerce_ai`.`orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `commerce_ai`.`orders` (
  `id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `created_at` DATETIME NULL,
  `finished` TINYINT(1) NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`, `user_id`),
  INDEX `fk_orders_users1_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_orders_users1`
    FOREIGN KEY (`user_id`)
    REFERENCES `commerce_ai`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `commerce_ai`.`orders_items`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `commerce_ai`.`orders_items` (
  `id` INT NOT NULL,
  `price` DECIMAL(13,2) NULL,
  `amount` INT NULL,
  `product_id` INT NOT NULL,
  `order_id` INT NOT NULL,
  PRIMARY KEY (`id`, `product_id`, `order_id`),
  INDEX `fk_orders_items_orders1_idx` (`order_id` ASC) VISIBLE,
  INDEX `fk_orders_items_products1_idx` (`product_id` ASC) VISIBLE,
  CONSTRAINT `fk_orders_items_orders1`
    FOREIGN KEY (`order_id`)
    REFERENCES `commerce_ai`.`orders` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_orders_items_products1`
    FOREIGN KEY (`product_id`)
    REFERENCES `commerce_ai`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `commerce_ai`.`payments`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `commerce_ai`.`payments` (
  `id` INT NOT NULL,
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
    REFERENCES `commerce_ai`.`orders` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_payments_cards1`
    FOREIGN KEY (`card_id`)
    REFERENCES `commerce_ai`.`cards` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


