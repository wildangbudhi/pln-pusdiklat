CREATE DATABASE IF NOT EXISTS `discussion`;

USE `discussion`;

DROP TABLE IF EXISTS `user_auth`;

CREATE TABLE `user_auth` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `full_name` varchar(255) DEFAULT NULL,
  `avatar_file` varchar(64) DEFAULT NULL,
  `email` varchar(255) NOT NULL,
  `username` varchar(64) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_auth_UN_email` (`email`),
  UNIQUE KEY `user_auth_UN_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `category_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `forum` (
  `id` varchar(50) NOT NULL,
  `title` varchar(100) NOT NULL,
  `question` text DEFAULT NULL,
  `author_user_id` bigint NOT NULL,
  `category_id` bigint NOT NULL,
  `status` varchar(100) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `forum_FK_author_user_id` (`author_user_id`),
  KEY `forum_FK_category_id` (`category_id`),
  CONSTRAINT `forum_FK_author_user_id` FOREIGN KEY (`author_user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `forum_FK_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `expertises`;

CREATE TABLE `expertises` (
  `id` varchar(50) NOT NULL,
  `identity_id` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `tempat_tanggal_lahir` varchar(200) DEFAULT NULL,
  `telp_number` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `category_id` bigint DEFAULT NULL,
  `pendidikan_terakhir` varchar(100) DEFAULT NULL,
  `jenis_kelamin` varchar(20) DEFAULT NULL,
  `rekening_number` varchar(100) DEFAULT NULL,
  `rekening_owner_name` varchar(100) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `expertises_FK_category_id` (`category_id`),
  CONSTRAINT `expertises_FK_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `discussion_requests`;

CREATE TABLE `discussion_requests` (
  `id` varchar(50) NOT NULL,
  `forum_id` varchar(50) NOT NULL,
  `accepted` tinyint(1) DEFAULT NULL,
  `accepted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `reqeuster_user_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `discussion_requests_FK` (`forum_id`),
  KEY `discussion_requests_FK_1` (`reqeuster_user_id`),
  CONSTRAINT `discussion_requests_FK` FOREIGN KEY (`forum_id`) REFERENCES `forum` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `discussion_requests_FK_1` FOREIGN KEY (`reqeuster_user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;