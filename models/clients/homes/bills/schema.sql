CREATE TABLE IF NOT EXISTS `clients_homes_bills` (
  `chb_id` int NOT NULL AUTO_INCREMENT,
  `chb_home_id` int DEFAULT NULL,
  `chb_electricity` int DEFAULT NULL,
  `chb_natural_gas` int DEFAULT NULL,
  `chb_propane_gas` int DEFAULT NULL,
  `chb_wood` int DEFAULT NULL,
  `chb_heating_oil` int DEFAULT NULL,
  `chb_year` year DEFAULT NULL,
  PRIMARY KEY (`chb_id`),
  UNIQUE KEY `chb_home_id` (`chb_home_id`),
  KEY `FK_home_bill_home_id` (`chb_home_id`),
  CONSTRAINT `FK_client_home_bill_home_id` FOREIGN KEY (`chb_home_id`) REFERENCES `clients_homes` (`ch_id`) ON DELETE CASCADE
)