CREATE TABLE IF NOT EXISTS `equipments_products` (
  `ep_id` int NOT NULL AUTO_INCREMENT,
  `ep_product_id` int DEFAULT NULL,
  `ep_price` decimal(10,2) DEFAULT NULL,
  `ep_vat` decimal(10,2) DEFAULT NULL,
  `ep_installation_price` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`ep_id`) USING BTREE
)