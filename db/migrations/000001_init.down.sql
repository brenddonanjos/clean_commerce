-- -----------------------------------------------------
-- Rollback actions in reverse order
-- -----------------------------------------------------

-- Drop table `commerce_ai`.`payments`
DROP TABLE IF EXISTS `commerce_ai`.`payments`;

-- Drop table `commerce_ai`.`orders_items`
DROP TABLE IF EXISTS `commerce_ai`.`orders_items`;

-- Drop table `commerce_ai`.`orders`
DROP TABLE IF EXISTS `commerce_ai`.`orders`;

-- Drop table `commerce_ai`.`products_categories`
DROP TABLE IF EXISTS `commerce_ai`.`products_categories`;

-- Drop table `commerce_ai`.`categories`
DROP TABLE IF EXISTS `commerce_ai`.`categories`;

-- Drop table `commerce_ai`.`addresses`
DROP TABLE IF EXISTS `commerce_ai`.`addresses`;

-- Drop table `commerce_ai`.`cards`
DROP TABLE IF EXISTS `commerce_ai`.`cards`;

-- Drop table `commerce_ai`.`products`
DROP TABLE IF EXISTS `commerce_ai`.`products`;

-- Drop table `commerce_ai`.`users`
DROP TABLE IF EXISTS `commerce_ai`.`users`;