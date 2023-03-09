CREATE TABLE IF NOT EXISTS `clients_ass` (
  `ca_id` int NOT NULL AUTO_INCREMENT,
  `ca_client_id` int DEFAULT NULL,
  `ca_company_id` int DEFAULT NULL,
  `ca_call_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `ca_call_reason` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ca_intervention_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `ca_comment` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ca_is_resolved` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`ca_id`),
  KEY `FK_cs_client_id` (`ca_client_id`),
  KEY `FK_ca_company_id` (`ca_company_id`),
  CONSTRAINT `FK_ca_client_id` FOREIGN KEY (`ca_client_id`) REFERENCES `clients` (`client_id`),
  CONSTRAINT `FK_ca_company_id` FOREIGN KEY (`ca_company_id`) REFERENCES `companies` (`c_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Client after sale service table';
