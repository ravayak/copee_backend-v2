CREATE TABLE IF NOT EXISTS `logs` (
  `log_id` int NOT NULL AUTO_INCREMENT,
  `log_event_desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `log_event_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `log_event_date` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`log_id`),
  KEY `log_id` (`log_id`)
) 