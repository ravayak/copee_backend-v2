CREATE TABLE IF NOT EXISTS `clients_homes` (
  `ch_id` int NOT NULL AUTO_INCREMENT,
  `ch_client_id` int DEFAULT NULL,
  `ch_geo_id` int DEFAULT NULL,
  `ch_construction_year` int DEFAULT NULL,
  `ch_area` int DEFAULT NULL,
  `ch_residents` int DEFAULT NULL,
  `ch_roof_positionning` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `ch_roof_slope` int DEFAULT NULL,
  `ch_label` tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ch_tr` int DEFAULT NULL,
  `ch_huc` double DEFAULT NULL,
  `ch_isolation` tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  PRIMARY KEY (`ch_id`),
  KEY `FK_client_id` (`ch_client_id`),
  KEY `FK_geo_id` (`ch_geo_id`),
  CONSTRAINT `FK_client_home_id` FOREIGN KEY (`ch_client_id`) REFERENCES `clients` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_geo_id` FOREIGN KEY (`ch_geo_id`) REFERENCES `geo` (`chg_id`) ON DELETE CASCADE
)