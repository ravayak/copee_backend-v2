CREATE TABLE IF NOT EXISTS `clients_homes_equipments` (
  `che_id` int NOT NULL AUTO_INCREMENT,
  `che_home_id` int DEFAULT NULL,
  `che_equipment_type` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `che_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  PRIMARY KEY (`che_id`),
  KEY `FK_che_home_id` (`che_home_id`),
  CONSTRAINT `FK_che_home_id` FOREIGN KEY (`che_home_id`) REFERENCES `clients_homes` (`ch_id`) ON DELETE CASCADE
)