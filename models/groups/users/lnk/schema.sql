CREATE TABLE IF NOT EXISTS `groups_users_lnk` (
  `gu_id` int NOT NULL AUTO_INCREMENT,
  `gu_user_id` int DEFAULT NULL,
  `gu_group_id` int DEFAULT NULL,
  PRIMARY KEY (`gu_id`) USING BTREE,
  KEY `FK_usu_user_id` (`gu_user_id`) USING BTREE,
  KEY `FK_usu_group_id` (`gu_group_id`) USING BTREE,
  CONSTRAINT `FK_gu_group_id` FOREIGN KEY (`gu_group_id`) REFERENCES `groups` (`g_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_gu_user_id` FOREIGN KEY (`gu_user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
)