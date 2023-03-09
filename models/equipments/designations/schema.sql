CREATE TABLE IF NOT EXISTS `equipments_designations` (
  `ea_id` int NOT NULL AUTO_INCREMENT,
  `ea_product_id` int DEFAULT NULL,
  `ea_article` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ea_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ea_title` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ea_id`),
  KEY `FK_ea_product_id` (`ea_product_id`),
  CONSTRAINT `FK_ea_product_id` FOREIGN KEY (`ea_product_id`) REFERENCES `equipments_products` (`ep_id`)
)