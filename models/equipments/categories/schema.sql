CREATE TABLE IF NOT EXISTS `equipments_categories` (
  `ec_id` int NOT NULL AUTO_INCREMENT,
  `ec_category_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `ec_description` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`ec_id`)
)