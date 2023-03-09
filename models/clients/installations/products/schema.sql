CREATE TABLE IF NOT EXISTS `clients_installations_products` (
  `cip_id` int NOT NULL AUTO_INCREMENT,
  `cip_installation_id` int DEFAULT NULL,
  `cip_product_id` int DEFAULT NULL,
  `cip_discount` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`cip_id`),
  KEY `FK_cip_products_id` (`cip_product_id`) USING BTREE,
  KEY `FK_cip_installations_id` (`cip_installation_id`) USING BTREE,
  CONSTRAINT `FK_cip_installations_id` FOREIGN KEY (`cip_installation_id`) REFERENCES `clients_installations` (`ci_id`),
  CONSTRAINT `FK_cip_products_id` FOREIGN KEY (`cip_product_id`) REFERENCES `equipments_products` (`ep_id`)
)