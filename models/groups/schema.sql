CREATE TABLE IF NOT EXISTS `groups` (
  `g_id` int NOT NULL AUTO_INCREMENT,
  `g_name` varchar(50) DEFAULT NULL,
  `g_description` text,
  PRIMARY KEY (`g_id`)
)