CREATE TABLE IF NOT EXISTS `clients_installations_payments` (
  `cip_id` int NOT NULL AUTO_INCREMENT,
  `cip_installation_id` int DEFAULT NULL,
  `cip_amount` double DEFAULT NULL,
  `cip_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cip_date` datetime DEFAULT NULL,
  `cip_description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cip_transaction_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cip_vat_rate` double DEFAULT NULL,
  PRIMARY KEY (`cip_id`),
  KEY `FK_cp_installation_id` (`cip_installation_id`),
  CONSTRAINT `FK_cp_installation_id` FOREIGN KEY (`cip_installation_id`) REFERENCES `clients_installations` (`ci_id`) ON DELETE CASCADE
) 