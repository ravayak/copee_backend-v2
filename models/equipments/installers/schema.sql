CREATE TABLE IF NOT EXISTS `equipments_installers` (
  `ei_id` int NOT NULL AUTO_INCREMENT,
  `ei_product_id` int DEFAULT NULL,
  `ei_company_id` int DEFAULT NULL,
  PRIMARY KEY (`ei_id`),
  KEY `FK_product_id` (`ei_product_id`),
  KEY `FK_company_id` (`ei_company_id`),
  CONSTRAINT `FK_company_id` FOREIGN KEY (`ei_company_id`) REFERENCES `companies` (`c_id`),
  CONSTRAINT `FK_product_id` FOREIGN KEY (`ei_product_id`) REFERENCES `equipments_products` (`ep_id`)
)