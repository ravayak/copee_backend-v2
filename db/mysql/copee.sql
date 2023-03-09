-- --------------------------------------------------------
-- Hôte:                         127.0.0.1
-- Version du serveur:           8.0.30 - MySQL Community Server - GPL
-- SE du serveur:                Win64
-- HeidiSQL Version:             12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Listage de la structure de la base pour copee-v2
CREATE DATABASE IF NOT EXISTS `copee-v2` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `copee-v2`;

-- Listage de la structure de table copee-v2. clients
CREATE TABLE IF NOT EXISTS `clients` (
  `client_id` int NOT NULL AUTO_INCREMENT,
  `client_last_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `client_first_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `client_email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `client_phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `client_fiscal_year_income` int DEFAULT NULL,
  `client_date_created` datetime DEFAULT CURRENT_TIMESTAMP,
  `client_birthdate` date DEFAULT NULL,
  PRIMARY KEY (`client_id`),
  UNIQUE KEY `client_email` (`client_email`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.clients : ~3 rows (environ)
REPLACE INTO `clients` (`client_id`, `client_last_name`, `client_first_name`, `client_email`, `client_phone`, `client_fiscal_year_income`, `client_date_created`, `client_birthdate`) VALUES
	(1, 'test tt', 'test tt', 'toto@toto.com', '1', 0, '2023-02-28 10:41:10', NULL),
	(2, 'test2', 'test2', 'test2', 'test2', NULL, '2023-03-02 12:04:29', NULL),
	(11, 'test tt', 'test tt', 'aa', 'aa', 0, '2023-03-06 17:19:39', NULL);

-- Listage de la structure de table copee-v2. clients_ass
CREATE TABLE IF NOT EXISTS `clients_ass` (
  `ca_id` int NOT NULL AUTO_INCREMENT,
  `ca_client_id` int DEFAULT NULL,
  `ca_company_id` int DEFAULT NULL,
  `ca_call_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `ca_call_reason` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ca_intervention_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `ca_comment` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ca_is_resolved` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`ca_id`),
  KEY `FK_cs_client_id` (`ca_client_id`),
  KEY `FK_ca_company_id` (`ca_company_id`),
  CONSTRAINT `FK_ca_client_id` FOREIGN KEY (`ca_client_id`) REFERENCES `clients` (`client_id`),
  CONSTRAINT `FK_ca_company_id` FOREIGN KEY (`ca_company_id`) REFERENCES `companies` (`c_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Client after sale service table';

-- Listage des données de la table copee-v2.clients_ass : ~0 rows (environ)

-- Listage de la structure de table copee-v2. clients_homes
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.clients_homes : ~3 rows (environ)
REPLACE INTO `clients_homes` (`ch_id`, `ch_client_id`, `ch_geo_id`, `ch_construction_year`, `ch_area`, `ch_residents`, `ch_roof_positionning`, `ch_roof_slope`, `ch_label`, `ch_tr`, `ch_huc`, `ch_isolation`) VALUES
	(1, 1, NULL, 1990, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL),
	(2, 2, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

-- Listage de la structure de table copee-v2. clients_homes_bills
CREATE TABLE IF NOT EXISTS `clients_homes_bills` (
  `chb_id` int NOT NULL AUTO_INCREMENT,
  `chb_home_id` int DEFAULT NULL,
  `chb_electricity` int DEFAULT NULL,
  `chb_natural_gas` int DEFAULT NULL,
  `chb_propane_gas` int DEFAULT NULL,
  `chb_wood` int DEFAULT NULL,
  `chb_heating_oil` int DEFAULT NULL,
  `chb_year` year DEFAULT NULL,
  PRIMARY KEY (`chb_id`),
  UNIQUE KEY `chb_home_id` (`chb_home_id`),
  KEY `FK_home_bill_home_id` (`chb_home_id`),
  CONSTRAINT `FK_client_home_bill_home_id` FOREIGN KEY (`chb_home_id`) REFERENCES `clients_homes` (`ch_id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.clients_homes_bills : ~2 rows (environ)
REPLACE INTO `clients_homes_bills` (`chb_id`, `chb_home_id`, `chb_electricity`, `chb_natural_gas`, `chb_propane_gas`, `chb_wood`, `chb_heating_oil`, `chb_year`) VALUES
	(1, 1, 100, NULL, NULL, NULL, NULL, NULL),
	(3, 2, NULL, NULL, 100, NULL, NULL, NULL);

-- Listage de la structure de table copee-v2. clients_homes_equipments
CREATE TABLE IF NOT EXISTS `clients_homes_equipments` (
  `che_id` int NOT NULL AUTO_INCREMENT,
  `che_home_id` int DEFAULT NULL,
  `che_equipment_type` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `che_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  PRIMARY KEY (`che_id`),
  KEY `FK_che_home_id` (`che_home_id`),
  CONSTRAINT `FK_che_home_id` FOREIGN KEY (`che_home_id`) REFERENCES `clients_homes` (`ch_id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.clients_homes_equipments : ~3 rows (environ)
REPLACE INTO `clients_homes_equipments` (`che_id`, `che_home_id`, `che_equipment_type`, `che_description`) VALUES
	(1, 1, 'eqmt1', 'eqmt1 description'),
	(2, 2, 'eqmt2', 'eqmt2 description');

-- Listage de la structure de table copee-v2. clients_installations
CREATE TABLE IF NOT EXISTS `clients_installations` (
  `ci_id` int NOT NULL AUTO_INCREMENT,
  `ci_client_id` int DEFAULT NULL,
  `ci_user_id` int DEFAULT NULL,
  `ci_shared_user_id` int DEFAULT NULL,
  `ci_creation_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `ci_update_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `ci_status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `ci_comments` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  PRIMARY KEY (`ci_id`),
  KEY `FK_client_installation_client_id` (`ci_client_id`),
  KEY `FK_client_installation_user_id` (`ci_user_id`),
  KEY `FK_ci_shared_user_id` (`ci_shared_user_id`),
  CONSTRAINT `FK_ci_installation_client_id` FOREIGN KEY (`ci_client_id`) REFERENCES `clients` (`client_id`) ON DELETE CASCADE,
  CONSTRAINT `FK_ci_installation_user_id` FOREIGN KEY (`ci_user_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `FK_ci_shared_user_id` FOREIGN KEY (`ci_shared_user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.clients_installations : ~2 rows (environ)
REPLACE INTO `clients_installations` (`ci_id`, `ci_client_id`, `ci_user_id`, `ci_shared_user_id`, `ci_creation_date`, `ci_update_date`, `ci_status`, `ci_comments`) VALUES
	(23, 1, 1, NULL, '2023-03-02 12:04:10', '2023-03-02 12:04:10', NULL, NULL),
	(25, 2, 1, NULL, '2023-03-02 12:04:10', '2023-03-02 12:04:10', NULL, NULL);

-- Listage de la structure de table copee-v2. clients_installations_files
CREATE TABLE IF NOT EXISTS `clients_installations_files` (
  `cif_id` int NOT NULL AUTO_INCREMENT,
  `cif_installation_id` int DEFAULT NULL,
  `cif_file_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cif_file_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cif_file_creation_date` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`cif_id`),
  KEY `FK_installation_id` (`cif_installation_id`),
  CONSTRAINT `FK_installation_id` FOREIGN KEY (`cif_installation_id`) REFERENCES `clients_installations` (`ci_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.clients_installations_files : ~0 rows (environ)

-- Listage de la structure de table copee-v2. clients_installations_payments
CREATE TABLE IF NOT EXISTS `clients_installations_payments` (
  `cip_id` int NOT NULL AUTO_INCREMENT,
  `cip_installation_id` int DEFAULT NULL,
  `cip_amount` double DEFAULT NULL,
  `cip_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cip_date` datetime DEFAULT NULL,
  `cip_description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cip_transaction_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cip_vat_rate` double DEFAULT NULL,
  PRIMARY KEY (`cip_id`),
  KEY `FK_cp_installation_id` (`cip_installation_id`),
  CONSTRAINT `FK_cp_installation_id` FOREIGN KEY (`cip_installation_id`) REFERENCES `clients_installations` (`ci_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.clients_installations_payments : ~0 rows (environ)

-- Listage de la structure de table copee-v2. clients_installations_products
CREATE TABLE IF NOT EXISTS `clients_installations_products` (
  `cip_id` int NOT NULL AUTO_INCREMENT,
  `cip_installation_id` int DEFAULT NULL,
  `cip_product_id` int DEFAULT NULL,
  `cip_discount` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`cip_id`),
  KEY `FK_cip_products_id` (`cip_product_id`) USING BTREE,
  KEY `FK_cip_installations_id` (`cip_installation_id`) USING BTREE,
  CONSTRAINT `FK_cip_installations_id` FOREIGN KEY (`cip_installation_id`) REFERENCES `clients_installations` (`ci_id`) ON DELETE CASCADE,
  CONSTRAINT `FK_cip_products_id` FOREIGN KEY (`cip_product_id`) REFERENCES `equipments_products` (`ep_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.clients_installations_products : ~0 rows (environ)

-- Listage de la structure de table copee-v2. clients_loans
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.clients_loans : ~0 rows (environ)

-- Listage de la structure de table copee-v2. companies
CREATE TABLE IF NOT EXISTS `companies` (
  `c_id` int NOT NULL AUTO_INCREMENT,
  `c_geo_id` int DEFAULT NULL,
  `c_name` varchar(50) DEFAULT NULL,
  `c_rcs` varchar(50) DEFAULT NULL,
  `c_siret` varchar(50) DEFAULT NULL,
  `c_siren` varchar(50) DEFAULT NULL,
  `c_intra_eu_vat` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `c_legal_status` varchar(50) DEFAULT NULL,
  `c_creation_date` varchar(50) DEFAULT NULL,
  `c_capital` int DEFAULT NULL,
  PRIMARY KEY (`c_id`),
  KEY `FK_c_geo_id` (`c_geo_id`),
  CONSTRAINT `FK_c_geo_id` FOREIGN KEY (`c_geo_id`) REFERENCES `geo` (`chg_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.companies : ~0 rows (environ)

-- Listage de la structure de table copee-v2. energies_costs
CREATE TABLE IF NOT EXISTS `energies_costs` (
  `enc_id` int NOT NULL AUTO_INCREMENT,
  `enc_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `enc_date` date DEFAULT NULL,
  `enc_cost` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`enc_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.energies_costs : ~0 rows (environ)

-- Listage de la structure de table copee-v2. equipments_catalogue
CREATE TABLE IF NOT EXISTS `equipments_catalogue` (
  `ec_id` int NOT NULL AUTO_INCREMENT,
  `ec_product_id` int DEFAULT NULL,
  `ec_quantity` int DEFAULT NULL,
  `ec_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ec_date_added` datetime DEFAULT CURRENT_TIMESTAMP,
  `ec_date_updated` datetime DEFAULT NULL,
  PRIMARY KEY (`ec_id`),
  KEY `FKec_product_id` (`ec_product_id`),
  CONSTRAINT `FKec_product_id` FOREIGN KEY (`ec_product_id`) REFERENCES `equipments_products` (`ep_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.equipments_catalogue : ~0 rows (environ)

-- Listage de la structure de table copee-v2. equipments_categories
CREATE TABLE IF NOT EXISTS `equipments_categories` (
  `ec_id` int NOT NULL AUTO_INCREMENT,
  `ec_category_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `ec_description` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`ec_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.equipments_categories : ~0 rows (environ)

-- Listage de la structure de table copee-v2. equipments_designations
CREATE TABLE IF NOT EXISTS `equipments_designations` (
  `ea_id` int NOT NULL AUTO_INCREMENT,
  `ea_product_id` int DEFAULT NULL,
  `ea_article` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ea_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ea_title` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ea_id`),
  KEY `FK_ea_product_id` (`ea_product_id`),
  CONSTRAINT `FK_ea_product_id` FOREIGN KEY (`ea_product_id`) REFERENCES `equipments_products` (`ep_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Articles décrivant es équipements à inclure dans les devis par exemple.';

-- Listage des données de la table copee-v2.equipments_designations : ~0 rows (environ)

-- Listage de la structure de table copee-v2. equipments_installers
CREATE TABLE IF NOT EXISTS `equipments_installers` (
  `ei_id` int NOT NULL AUTO_INCREMENT,
  `ei_product_id` int DEFAULT NULL,
  `ei_company_id` int DEFAULT NULL,
  PRIMARY KEY (`ei_id`),
  KEY `FK_product_id` (`ei_product_id`),
  KEY `FK_company_id` (`ei_company_id`),
  CONSTRAINT `FK_company_id` FOREIGN KEY (`ei_company_id`) REFERENCES `companies` (`c_id`),
  CONSTRAINT `FK_product_id` FOREIGN KEY (`ei_product_id`) REFERENCES `equipments_products` (`ep_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.equipments_installers : ~0 rows (environ)

-- Listage de la structure de table copee-v2. equipments_products
CREATE TABLE IF NOT EXISTS `equipments_products` (
  `ep_id` int NOT NULL AUTO_INCREMENT,
  `ep_product_id` int DEFAULT NULL,
  `ep_price` decimal(10,2) DEFAULT NULL,
  `ep_vat` decimal(10,2) DEFAULT NULL,
  `ep_installation_price` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`ep_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.equipments_products : ~0 rows (environ)

-- Listage de la structure de table copee-v2. equipments_services
CREATE TABLE IF NOT EXISTS `equipments_services` (
  `es_id` int NOT NULL AUTO_INCREMENT,
  `es_description` text,
  `es_price` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`es_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.equipments_services : ~0 rows (environ)

-- Listage de la structure de table copee-v2. equipments_suppliers
CREATE TABLE IF NOT EXISTS `equipments_suppliers` (
  `es_id` int NOT NULL AUTO_INCREMENT,
  `es_companies_id` int DEFAULT NULL,
  `es_products_id` int DEFAULT NULL,
  PRIMARY KEY (`es_id`),
  KEY `FK_es_companies_id` (`es_companies_id`),
  KEY `FK_es_products_id` (`es_products_id`),
  CONSTRAINT `FK_es_companies_id` FOREIGN KEY (`es_companies_id`) REFERENCES `companies` (`c_id`),
  CONSTRAINT `FK_es_products_id` FOREIGN KEY (`es_products_id`) REFERENCES `equipments_products` (`ep_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.equipments_suppliers : ~0 rows (environ)

-- Listage de la structure de table copee-v2. geo
CREATE TABLE IF NOT EXISTS `geo` (
  `chg_id` int NOT NULL AUTO_INCREMENT,
  `chg_zone` tinytext,
  `chg_latitude` double DEFAULT NULL,
  `chg_longitude` double DEFAULT NULL,
  `chg_altitude` int DEFAULT NULL,
  `chg_department` int DEFAULT NULL,
  `chg_city` varchar(50) DEFAULT NULL,
  `chg_address` varchar(50) DEFAULT NULL,
  `chg_postcode` int DEFAULT NULL,
  PRIMARY KEY (`chg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.geo : ~0 rows (environ)

-- Listage de la structure de table copee-v2. groups
CREATE TABLE IF NOT EXISTS `groups` (
  `g_id` int NOT NULL AUTO_INCREMENT,
  `g_name` varchar(50) DEFAULT NULL,
  `g_description` text,
  PRIMARY KEY (`g_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.groups : ~5 rows (environ)
REPLACE INTO `groups` (`g_id`, `g_name`, `g_description`) VALUES
	(1, 'USER_STATUS', 'Users group defining privileges for normal users'),
	(2, 'ADMIN_STATUS', 'Users group defining privileges for admin users'),
	(3, 'SUPERADMIN_STATUS', 'Users group defining privileges for super admin users'),
	(4, 'MANAGER_POSITION', 'Group defining privileges for manager'),
	(5, 'VRP_POSITION', NULL),
	(6, 'TEST_TEAM', 'team managed by Test user'),
	(7, 'COMMERCIAL_DIRECTOR', NULL),
	(8, 'CEO', 'PDG');

-- Listage de la structure de table copee-v2. groups_privileges
CREATE TABLE IF NOT EXISTS `groups_privileges` (
  `gp_privilege_id` int NOT NULL AUTO_INCREMENT,
  `gp_privilege` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`gp_privilege_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.groups_privileges : ~4 rows (environ)
REPLACE INTO `groups_privileges` (`gp_privilege_id`, `gp_privilege`) VALUES
	(4, 'CLIENT_GET'),
	(5, 'CLIENT_DELETE'),
	(6, 'CLIENT_UPDATE'),
	(7, 'CLIENT_CREATE'),
	(8, 'CLIENT_LIST'),
	(9, 'COMPANIES_GET'),
	(10, 'COMPANIES_DELETE'),
	(11, 'COMPANIES_UPDATE'),
	(12, 'COMPANIES_CREATE'),
	(13, 'COMPANIES_LIST');

-- Listage de la structure de table copee-v2. groups_privileges_lnk
CREATE TABLE IF NOT EXISTS `groups_privileges_lnk` (
  `gpl_id` int NOT NULL AUTO_INCREMENT,
  `gpl_group_id` int DEFAULT NULL,
  `gpl_privilege_id` int DEFAULT NULL,
  PRIMARY KEY (`gpl_id`) USING BTREE,
  UNIQUE KEY `gpl_group_id_gpl_privilege_id` (`gpl_group_id`,`gpl_privilege_id`),
  KEY `FK_gpl_privilege_id` (`gpl_privilege_id`),
  CONSTRAINT `FK_gpl_group_id` FOREIGN KEY (`gpl_group_id`) REFERENCES `groups` (`g_id`),
  CONSTRAINT `FK_gpl_privilege_id` FOREIGN KEY (`gpl_privilege_id`) REFERENCES `groups_privileges` (`gp_privilege_id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.groups_privileges_lnk : ~11 rows (environ)
REPLACE INTO `groups_privileges_lnk` (`gpl_id`, `gpl_group_id`, `gpl_privilege_id`) VALUES
	(2, 1, 4),
	(6, 1, 6),
	(7, 1, 7),
	(16, 1, 8),
	(17, 1, 9),
	(18, 1, 13),
	(8, 2, 4),
	(9, 2, 5),
	(10, 2, 6),
	(11, 2, 7),
	(20, 2, 9),
	(19, 2, 10),
	(21, 2, 11),
	(22, 2, 12),
	(23, 2, 13),
	(12, 3, 4),
	(13, 3, 5),
	(14, 3, 6),
	(15, 3, 7),
	(24, 3, 9),
	(25, 3, 10),
	(26, 3, 11),
	(27, 3, 12),
	(28, 3, 13);

-- Listage de la structure de table copee-v2. groups_users_lnk
CREATE TABLE IF NOT EXISTS `groups_users_lnk` (
  `gu_id` int NOT NULL AUTO_INCREMENT,
  `gu_user_id` int DEFAULT NULL,
  `gu_group_id` int DEFAULT NULL,
  PRIMARY KEY (`gu_id`) USING BTREE,
  UNIQUE KEY `gu_user_id_gu_group_id` (`gu_user_id`,`gu_group_id`),
  KEY `FK_usu_user_id` (`gu_user_id`) USING BTREE,
  KEY `FK_usu_group_id` (`gu_group_id`) USING BTREE,
  CONSTRAINT `FK_gu_group_id` FOREIGN KEY (`gu_group_id`) REFERENCES `groups` (`g_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_gu_user_id` FOREIGN KEY (`gu_user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.groups_users_lnk : ~4 rows (environ)
REPLACE INTO `groups_users_lnk` (`gu_id`, `gu_user_id`, `gu_group_id`) VALUES
	(3, 1, 1),
	(6, 2, 6);

-- Listage de la structure de table copee-v2. logs
CREATE TABLE IF NOT EXISTS `logs` (
  `log_id` int NOT NULL AUTO_INCREMENT,
  `log_event_desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `log_event_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `log_event_date` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`log_id`),
  KEY `log_id` (`log_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.logs : ~0 rows (environ)

-- Listage de la structure de table copee-v2. sales_commissions
CREATE TABLE IF NOT EXISTS `sales_commissions` (
  `sc_id` int NOT NULL AUTO_INCREMENT,
  `sc_commissions_discount_id` int DEFAULT NULL,
  `sc_group_id` int DEFAULT NULL,
  PRIMARY KEY (`sc_id`),
  KEY `FK_sc_commissions_discount_id` (`sc_commissions_discount_id`),
  KEY `FK_sc_group_id` (`sc_group_id`),
  CONSTRAINT `FK_sc_commissions_discount_id` FOREIGN KEY (`sc_commissions_discount_id`) REFERENCES `sales_commissions_discount` (`sc_id`),
  CONSTRAINT `FK_sc_group_id` FOREIGN KEY (`sc_group_id`) REFERENCES `groups` (`g_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.sales_commissions : ~0 rows (environ)

-- Listage de la structure de table copee-v2. sales_commissions_discount
CREATE TABLE IF NOT EXISTS `sales_commissions_discount` (
  `sc_id` int NOT NULL AUTO_INCREMENT,
  `sc_discount_pct` double DEFAULT NULL,
  `sc_com_pct` double DEFAULT NULL,
  PRIMARY KEY (`sc_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.sales_commissions_discount : ~0 rows (environ)

-- Listage de la structure de table copee-v2. users
CREATE TABLE IF NOT EXISTS `users` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `user_firstname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `user_lastname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `user_email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `user_phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `user_is_active` tinyint(1) DEFAULT NULL,
  `user_date_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `user_recruitment_date` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `email` (`user_email`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Listage des données de la table copee-v2.users : ~2 rows (environ)
REPLACE INTO `users` (`user_id`, `user_firstname`, `user_lastname`, `user_email`, `user_phone`, `user_is_active`, `user_date_created`, `user_recruitment_date`) VALUES
	(1, 'test', 'est', 'ynnotna@gmail.com', NULL, 1, '2023-02-28 11:53:59', '2023-02-28 11:53:59'),
	(2, 'test2', 'test2', 'toto@gmail.com', NULL, NULL, '2023-02-28 13:02:09', '2023-02-28 13:02:09');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
