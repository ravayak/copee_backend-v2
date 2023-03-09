CREATE TABLE IF NOT EXISTS `companies` (
  `c_id` int NOT NULL AUTO_INCREMENT,
  `c_geo_id` int,
  `c_name` varchar(50),
  `c_rcs` varchar(50),
  `c_siret` varchar(50),
  `c_siren` varchar(50),
  `c_intra_eu_vat` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `c_legal_status` varchar(50),
  `c_creation_date` varchar(50),
  `c_capital` int,
  PRIMARY KEY (`c_id`),
  KEY `FK_c_geo_id` (`c_geo_id`),
  CONSTRAINT `FK_c_geo_id` FOREIGN KEY (`c_geo_id`) REFERENCES `geo` (`chg_id`)
)