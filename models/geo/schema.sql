CREATE TABLE `geo` (
  `chg_id` int NOT NULL AUTO_INCREMENT,
  `chg_zone` tinytext,
  `chg_latitude` double,
  `chg_longitude` double,
  `chg_altitude` int,
  `chg_department` int,
  `chg_city` varchar(50),
  `chg_address` varchar(50),
  `chg_postcode` int,
  PRIMARY KEY (`chg_id`)
) 