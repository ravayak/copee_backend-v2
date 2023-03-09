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


-- Listage de la structure de la base pour copee-v2-test
CREATE DATABASE IF NOT EXISTS `copee-v2-test` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `copee-v2-test`;

-- Listage de la structure de table copee-v2-test. clients
CREATE TABLE IF NOT EXISTS `clients` (
  `client_id` int NOT NULL AUTO_INCREMENT,
  `client_last_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `client_first_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `client_email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `client_phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `client_fiscal_year_income` int DEFAULT NULL,
  `client_date_created` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`client_id`),
  UNIQUE KEY `client_email` (`client_email`)
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. clients_ass
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Client after sale service table';

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. clients_homes
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
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. clients_homes_bills
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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. clients_homes_equipments
CREATE TABLE IF NOT EXISTS `clients_homes_equipments` (
  `che_id` int NOT NULL AUTO_INCREMENT,
  `che_home_id` int DEFAULT NULL,
  `che_equipment_type` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `che_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  PRIMARY KEY (`che_id`),
  KEY `FK_che_home_id` (`che_home_id`),
  CONSTRAINT `FK_che_home_id` FOREIGN KEY (`che_home_id`) REFERENCES `clients_homes` (`ch_id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. clients_installations
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
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. clients_installations_files
CREATE TABLE IF NOT EXISTS `clients_installations_files` (
  `cif_id` int NOT NULL AUTO_INCREMENT,
  `cif_installation_id` int DEFAULT NULL,
  `cif_file_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cif_file_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cif_file_creation_date` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`cif_id`),
  KEY `FK_installation_id` (`cif_installation_id`),
  CONSTRAINT `FK_installation_id` FOREIGN KEY (`cif_installation_id`) REFERENCES `clients_installations` (`ci_id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. clients_installations_payments
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. clients_installations_products
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. clients_loans
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. companies
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. energies_costs
CREATE TABLE IF NOT EXISTS `energies_costs` (
  `enc_id` int NOT NULL AUTO_INCREMENT,
  `enc_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `enc_date` date DEFAULT NULL,
  `enc_cost` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`enc_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. equipments_catalogue
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. equipments_categories
CREATE TABLE IF NOT EXISTS `equipments_categories` (
  `ec_id` int NOT NULL AUTO_INCREMENT,
  `ec_category_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `ec_description` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`ec_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. equipments_designations
CREATE TABLE IF NOT EXISTS `equipments_designations` (
  `ea_id` int NOT NULL AUTO_INCREMENT,
  `ea_product_id` int DEFAULT NULL,
  `ea_article` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ea_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ea_title` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ea_id`),
  KEY `FK_ea_product_id` (`ea_product_id`),
  CONSTRAINT `FK_ea_product_id` FOREIGN KEY (`ea_product_id`) REFERENCES `equipments_products` (`ep_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Articles décrivant es équipements à inclure dans les devis par exemple.';

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. equipments_installers
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

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. equipments_products
CREATE TABLE IF NOT EXISTS `equipments_products` (
  `ep_id` int NOT NULL AUTO_INCREMENT,
  `ep_product_id` int DEFAULT NULL,
  `ep_price` decimal(10,2) DEFAULT NULL,
  `ep_vat` decimal(10,2) DEFAULT NULL,
  `ep_installation_price` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`ep_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. equipments_services
CREATE TABLE IF NOT EXISTS `equipments_services` (
  `es_id` int NOT NULL AUTO_INCREMENT,
  `es_description` text,
  `es_price` decimal(10,2) DEFAULT NULL,
  `es_title` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`es_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. equipments_suppliers
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

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. geo
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
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. groups
CREATE TABLE IF NOT EXISTS `groups` (
  `g_id` int NOT NULL AUTO_INCREMENT,
  `g_name` varchar(50) DEFAULT NULL,
  `g_description` text,
  PRIMARY KEY (`g_id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. groups_privileges
CREATE TABLE IF NOT EXISTS `groups_privileges` (
  `gp_privilege_id` int NOT NULL AUTO_INCREMENT,
  `gp_privilege` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`gp_privilege_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. groups_privileges_lnk
CREATE TABLE IF NOT EXISTS `groups_privileges_lnk` (
  `gpl_id` int NOT NULL AUTO_INCREMENT,
  `gpl_group_id` int DEFAULT NULL,
  `gpl_privilege_id` int DEFAULT NULL,
  PRIMARY KEY (`gpl_id`) USING BTREE,
  UNIQUE KEY `gpl_group_id_gpl_privilege_id` (`gpl_group_id`,`gpl_privilege_id`),
  KEY `FK_gpl_privilege_id` (`gpl_privilege_id`),
  CONSTRAINT `FK_gpl_group_id` FOREIGN KEY (`gpl_group_id`) REFERENCES `groups` (`g_id`),
  CONSTRAINT `FK_gpl_privilege_id` FOREIGN KEY (`gpl_privilege_id`) REFERENCES `groups_privileges_lnk` (`gpl_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. groups_users_lnk
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
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. logs
CREATE TABLE IF NOT EXISTS `logs` (
  `log_id` int NOT NULL AUTO_INCREMENT,
  `log_event_desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `log_event_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `log_event_date` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`log_id`),
  KEY `log_id` (`log_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. sales_commissions
CREATE TABLE IF NOT EXISTS `sales_commissions` (
  `sc_id` int NOT NULL AUTO_INCREMENT,
  `sc_commissions_discount_id` int DEFAULT NULL,
  `sc_group_id` int DEFAULT NULL,
  PRIMARY KEY (`sc_id`),
  KEY `FK_sc_commissions_discount_id` (`sc_commissions_discount_id`),
  KEY `FK_sc_group_id` (`sc_group_id`),
  CONSTRAINT `FK_sc_commissions_discount_id` FOREIGN KEY (`sc_commissions_discount_id`) REFERENCES `sales_commissions_discount` (`sc_id`),
  CONSTRAINT `FK_sc_group_id` FOREIGN KEY (`sc_group_id`) REFERENCES `groups` (`g_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. sales_commissions_discount
CREATE TABLE IF NOT EXISTS `sales_commissions_discount` (
  `sc_id` int NOT NULL AUTO_INCREMENT,
  `sc_discount_pct` double DEFAULT NULL,
  `sc_com_pct` double DEFAULT NULL,
  PRIMARY KEY (`sc_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de table copee-v2-test. users
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
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Les données exportées n'étaient pas sélectionnées.

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
