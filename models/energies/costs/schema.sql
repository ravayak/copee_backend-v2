CREATE TABLE IF NOT EXISTS `energies_costs` (
  `enc_id` int NOT NULL AUTO_INCREMENT,
  `enc_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `enc_date` date,
  `enc_cost` decimal(10,2),
  PRIMARY KEY (`enc_id`)
)