CREATE TABLE IF NOT EXISTS `groups_privileges` (
  `gp_privilege_id` int NOT NULL AUTO_INCREMENT,
  `gp_privilege` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`gp_privilege_id`) USING BTREE
)