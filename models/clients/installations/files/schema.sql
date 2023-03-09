CREATE TABLE IF NOT EXISTS `clients_installations_files` (
  `cif_id` int NOT NULL AUTO_INCREMENT,
  `cif_installation_id` int DEFAULT NULL,
  `cif_file_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cif_file_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cif_file_creation_date` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`cif_id`),
  KEY `FK_installation_id` (`cif_installation_id`),
  CONSTRAINT `FK_installation_id` FOREIGN KEY (`cif_installation_id`) REFERENCES `clients_installations` (`ci_id`) ON DELETE CASCADE
)