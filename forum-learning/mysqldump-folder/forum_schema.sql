CREATE DATABASE IF NOT EXISTS `forum`;

USE `forum`;

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

DROP TABLE IF EXISTS `forum`;

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
  CONSTRAINT `forum_FK_author_user_id` FOREIGN KEY (`author_user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT `forum_FK_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON UPDATE CASCADE ON DELETE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `forum_reaction`;

CREATE TABLE `forum_reaction` (
  `user_id` bigint NOT NULL,
  `forum_id` varchar(50) NOT NULL,
  `up_vote` tinyint(1) DEFAULT NULL,
  `down_vote` tinyint(1) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `forum_reaction_FK_user_id` (`user_id`),
  KEY `forum_reaction_FK` (`forum_id`),
  CONSTRAINT `forum_reaction_FK` FOREIGN KEY (`forum_id`) REFERENCES `forum` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT `forum_reaction_FK_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE ON DELETE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `forum_replies`;

CREATE TABLE `forum_replies` (
  `id` varchar(50) NOT NULL,
  `forum_id` varchar(50) NOT NULL,
  `author_user_id` bigint NOT NULL,
  `answer` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `forum_replies_FK_forum_id` (`forum_id`),
  KEY `forum_replies_FK` (`author_user_id`),
  CONSTRAINT `forum_replies_FK` FOREIGN KEY (`author_user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT `forum_replies_FK_forum_id` FOREIGN KEY (`forum_id`) REFERENCES `forum` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


DROP TABLE IF EXISTS `forum_replies_reactions`;

CREATE TABLE `forum_replies_reactions` (
  `user_id` bigint NOT NULL,
  `forum_id` varchar(50) NOT NULL,
  `forum_replies_id` varchar(50) NOT NULL,
  `up_vote` tinyint(1) DEFAULT NULL,
  `down_vote` tinyint(1) DEFAULT NULL,
  `agree` tinyint(1) DEFAULT NULL,
  `skeptic` tinyint(1) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `forum_replies_reactions_FK_user_id` (`user_id`),
  KEY `forum_replies_reactions_FK_forum_id` (`forum_id`),
  KEY `forum_replies_reactions_FKforum_replies_id` (`forum_replies_id`),
  CONSTRAINT `forum_replies_reactions_FK_forum_id` FOREIGN KEY (`forum_id`) REFERENCES `forum` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT `forum_replies_reactions_FK_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT `forum_replies_reactions_FKforum_replies_id` FOREIGN KEY (`forum_replies_id`) REFERENCES `forum_replies` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;