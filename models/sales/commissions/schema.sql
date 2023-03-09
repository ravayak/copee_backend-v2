CREATE TABLE IF NOT EXISTS `sales_commissions` (
  `sc_id` int NOT NULL AUTO_INCREMENT,
  `sc_commissions_discount_id` int DEFAULT NULL,
  `sc_group_id` int DEFAULT NULL,
  PRIMARY KEY (`sc_id`),
  KEY `FK_sc_commissions_discount_id` (`sc_commissions_discount_id`),
  KEY `FK_sc_group_id` (`sc_group_id`),
  CONSTRAINT `FK_sc_commissions_discount_id` FOREIGN KEY (`sc_commissions_discount_id`) REFERENCES `sales_commissions_discount` (`sc_id`),
  CONSTRAINT `FK_sc_group_id` FOREIGN KEY (`sc_group_id`) REFERENCES `groups` (`g_id`)
)