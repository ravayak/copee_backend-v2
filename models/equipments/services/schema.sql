CREATE TABLE IF NOT EXISTS `equipments_services` (
  `es_id` int NOT NULL AUTO_INCREMENT,
  `es_description` text,
  `es_price` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`es_id`)
)