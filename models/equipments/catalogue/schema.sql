CREATE TABLE IF NOT EXISTS `equipments_catalogue` (
  `ec_id` int NOT NULL AUTO_INCREMENT,
  `ec_product_id` int DEFAULT NULL,
  `ec_quantity` int DEFAULT NULL,
  `ec_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ec_date_added` datetime DEFAULT CURRENT_TIMESTAMP,
  `ec_date_updated` datetime DEFAULT NULL,
  PRIMARY KEY (`ec_id`),
  KEY `FKec_product_id` (`ec_product_id`),
  CONSTRAINT `FKec_product_id` FOREIGN KEY (`ec_product_id`) REFERENCES `equipments_products` (`ep_id`)
)