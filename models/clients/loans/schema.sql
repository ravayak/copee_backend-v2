CREATE TABLE IF NOT EXISTS `clients_loans` (
  `cl_id` int NOT NULL AUTO_INCREMENT,
  `cl_client_id` int DEFAULT NULL,
  `cl_amount` double DEFAULT NULL,
  `cl_installments` double DEFAULT NULL,
  `cl_client_provision` double DEFAULT NULL,
  `cl_client_prepayment` double DEFAULT NULL,
  `cl_insured` tinyint DEFAULT NULL,
  `cl_funding_agency` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`cl_id`),
  KEY `FK_loan_client_id` (`cl_client_id`),
  CONSTRAINT `FK_loan_client_id` FOREIGN KEY (`cl_client_id`) REFERENCES `clients` (`client_id`) ON DELETE CASCADE
)