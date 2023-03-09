CREATE TABLE IF NOT EXISTS `groups_privileges_lnk` (
  `gpl_id` int NOT NULL AUTO_INCREMENT,
  `gpl_group_id` int DEFAULT NULL,
  `gpl_privilege_id` int DEFAULT NULL,
  PRIMARY KEY (`gpl_id`) USING BTREE
)