CREATE TABLE IF NOT EXISTS `clients_installations` (
  `ci_id` int NOT NULL AUTO_INCREMENT,
  `ci_client_id` int DEFAULT NULL,
  `ci_user_id` int DEFAULT NULL,
  `ci_shared_user_id` int DEFAULT NULL,
  `ci_creation_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `ci_update_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `ci_status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `ci_comments` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  PRIMARY KEY (`ci_id`),
  KEY `FK_client_installation_client_id` (`ci_client_id`),
  KEY `FK_client_installation_user_id` (`ci_user_id`),
  KEY `FK_ci_shared_user_id` (`ci_shared_user_id`),
  CONSTRAINT `FK_ci_installation_client_id` FOREIGN KEY (`ci_client_id`) REFERENCES `clients` (`client_id`) ON DELETE CASCADE,
  CONSTRAINT `FK_ci_installation_user_id` FOREIGN KEY (`ci_user_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `FK_ci_shared_user_id` FOREIGN KEY (`ci_shared_user_id`) REFERENCES `users` (`user_id`)
)