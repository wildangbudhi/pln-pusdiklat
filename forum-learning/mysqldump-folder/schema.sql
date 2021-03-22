-- MySQL dump 10.13  Distrib 5.7.32, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: forum_learning
-- ------------------------------------------------------
-- Server version	8.0.22

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `activities`
--

DROP TABLE IF EXISTS `activities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `activities` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `module_name` varchar(100) DEFAULT NULL,
  `method` varchar(4) DEFAULT NULL,
  `url` varchar(100) DEFAULT NULL,
  `url_regex` varchar(100) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `category` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `category_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `discussion`
--

DROP TABLE IF EXISTS `discussion`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discussion` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `forum_id` varchar(50) NOT NULL,
  `title` varchar(100) DEFAULT NULL,
  `zoom_invitation_url` varchar(100) DEFAULT NULL,
  `zoom_host_email` varchar(100) DEFAULT NULL,
  `zoom_host_password` varchar(100) DEFAULT NULL,
  `start_datetime` timestamp NULL DEFAULT NULL,
  `end_datetime` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `discussion_request_id` bigint DEFAULT NULL,
  `status` varchar(100) DEFAULT NULL,
  `category_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `discussion_FK_forum_id` (`forum_id`),
  KEY `discussion_FK` (`discussion_request_id`),
  KEY `discussion_FK_1_category_id` (`category_id`),
  CONSTRAINT `discussion_FK` FOREIGN KEY (`discussion_request_id`) REFERENCES `discussion_requests` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `discussion_FK_1_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `discussion_FK_forum_id` FOREIGN KEY (`forum_id`) REFERENCES `forum` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `discussion_expertises`
--

DROP TABLE IF EXISTS `discussion_expertises`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discussion_expertises` (
  `expertise_id` bigint DEFAULT NULL,
  `discussion_id` bigint DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `discussion_expertises_FK_user_id` (`expertise_id`),
  KEY `discussion_expertises_FK` (`discussion_id`),
  CONSTRAINT `discussion_expertises_FK` FOREIGN KEY (`discussion_id`) REFERENCES `discussion` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `discussion_expertises_FK_expertise_id` FOREIGN KEY (`expertise_id`) REFERENCES `expertises` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `discussion_participants`
--

DROP TABLE IF EXISTS `discussion_participants`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discussion_participants` (
  `user_id` bigint DEFAULT NULL,
  `discussion_id` bigint DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `discussion_participants_FK_user_id` (`user_id`),
  KEY `discussion_participants_FK_discussion_id` (`discussion_id`),
  CONSTRAINT `discussion_participants_FK_discussion_id` FOREIGN KEY (`discussion_id`) REFERENCES `discussion` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `discussion_participants_FK_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `discussion_participants_waiting_list`
--

DROP TABLE IF EXISTS `discussion_participants_waiting_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discussion_participants_waiting_list` (
  `discussion_id` bigint DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  `accepted` tinyint(1) DEFAULT NULL,
  `accapted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `discussion_participants_waiting_list_FK` (`discussion_id`),
  KEY `discussion_participants_waiting_list_FK_1` (`user_id`),
  CONSTRAINT `discussion_participants_waiting_list_FK` FOREIGN KEY (`discussion_id`) REFERENCES `discussion` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `discussion_participants_waiting_list_FK_1` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `discussion_reaction`
--

DROP TABLE IF EXISTS `discussion_reaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discussion_reaction` (
  `user_id` bigint DEFAULT NULL,
  `discussion` bigint DEFAULT NULL,
  `up_vote` tinyint(1) DEFAULT NULL,
  `down_vote` tinyint(1) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `discussion_reaction_FK_user_id` (`user_id`),
  KEY `discussion_reaction_FK_discussion_id` (`discussion`),
  CONSTRAINT `discussion_reaction_FK_discussion_id` FOREIGN KEY (`discussion`) REFERENCES `discussion` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `discussion_reaction_FK_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `discussion_request_expertise_sugestions`
--

DROP TABLE IF EXISTS `discussion_request_expertise_sugestions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discussion_request_expertise_sugestions` (
  `discussion_request_id` bigint DEFAULT NULL,
  `expertise_id` bigint DEFAULT NULL,
  KEY `discussion_request_expertise_sugestions_FK_1` (`expertise_id`),
  KEY `discussion_request_expertise_sugestions_FK` (`discussion_request_id`),
  CONSTRAINT `discussion_request_expertise_sugestions_FK` FOREIGN KEY (`discussion_request_id`) REFERENCES `discussion_requests` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `discussion_request_expertise_sugestions_FK_expertise_id` FOREIGN KEY (`expertise_id`) REFERENCES `expertises` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `discussion_requests`
--

DROP TABLE IF EXISTS `discussion_requests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discussion_requests` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `forum_id` varchar(50) DEFAULT NULL,
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
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `discussion_requests_start_datetime_sugestions`
--

DROP TABLE IF EXISTS `discussion_requests_start_datetime_sugestions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discussion_requests_start_datetime_sugestions` (
  `discussion_request_id` bigint DEFAULT NULL,
  `start_datetime` timestamp NULL DEFAULT NULL,
  KEY `discussion_requests_start_datetime_sugestions_FK` (`discussion_request_id`),
  CONSTRAINT `discussion_requests_start_datetime_sugestions_FK` FOREIGN KEY (`discussion_request_id`) REFERENCES `discussion_requests` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `expertises`
--

DROP TABLE IF EXISTS `expertises`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `expertises` (
  `id` bigint NOT NULL AUTO_INCREMENT,
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
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `forum`
--

DROP TABLE IF EXISTS `forum`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
  CONSTRAINT `forum_FK_author_user_id` FOREIGN KEY (`author_user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT `forum_FK_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `forum_reaction`
--

DROP TABLE IF EXISTS `forum_reaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
  CONSTRAINT `forum_reaction_FK_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `forum_replies`
--

DROP TABLE IF EXISTS `forum_replies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
  CONSTRAINT `forum_replies_FK` FOREIGN KEY (`author_user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT `forum_replies_FK_forum_id` FOREIGN KEY (`forum_id`) REFERENCES `forum` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `forum_replies_reactions`
--

DROP TABLE IF EXISTS `forum_replies_reactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `forum_replies_reactions` (
  `user_id` bigint NOT NULL,
  `forum_replies_id` varchar(50) NOT NULL,
  `up_vote` tinyint(1) DEFAULT NULL,
  `down_vote` tinyint(1) DEFAULT NULL,
  `agree` tinyint(1) DEFAULT NULL,
  `skeptic` tinyint(1) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `forum_replies_reactions_FK_user_id` (`user_id`),
  KEY `forum_replies_reactions_FKforum_replies_id` (`forum_replies_id`),
  CONSTRAINT `forum_replies_reactions_FK_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT `forum_replies_reactions_FKforum_replies_id` FOREIGN KEY (`forum_replies_id`) REFERENCES `forum_replies` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `posts`
--

DROP TABLE IF EXISTS `posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `posts` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT NULL,
  `discussion_id` bigint DEFAULT NULL,
  `author_user_id` bigint DEFAULT NULL,
  `category_id` bigint DEFAULT NULL,
  `articles` text,
  `thumbnail` varchar(100) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `posts_FK_discussion_id` (`discussion_id`),
  KEY `posts_FK_author_user_id` (`author_user_id`),
  KEY `posts_FK_category_id` (`category_id`),
  CONSTRAINT `posts_FK_author_user_id` FOREIGN KEY (`author_user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `posts_FK_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `posts_FK_discussion_id` FOREIGN KEY (`discussion_id`) REFERENCES `discussion` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `posts_reactions`
--

DROP TABLE IF EXISTS `posts_reactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `posts_reactions` (
  `user_id` bigint DEFAULT NULL,
  `posts_id` bigint DEFAULT NULL,
  `up_vote` tinyint(1) DEFAULT NULL,
  `down_vote` tinyint(1) DEFAULT NULL,
  `agree` tinyint(1) DEFAULT NULL,
  `skeptic` tinyint(1) DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  KEY `posts_reactions_FK_user_id` (`user_id`),
  KEY `posts_reactions_FK_posts_id` (`posts_id`),
  CONSTRAINT `posts_reactions_FK_posts_id` FOREIGN KEY (`posts_id`) REFERENCES `posts` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `posts_reactions_FK_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `role_activities`
--

DROP TABLE IF EXISTS `role_activities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_activities` (
  `role_id` bigint DEFAULT NULL,
  `activities_id` bigint DEFAULT NULL,
  KEY `role_activities_FK` (`role_id`),
  KEY `role_activities_FK_1` (`activities_id`),
  CONSTRAINT `role_activities_FK` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `role_activities_FK_1` FOREIGN KEY (`activities_id`) REFERENCES `activities` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `roles_name` varchar(100) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_auth`
--

DROP TABLE IF EXISTS `user_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_auth` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `full_name` varchar(255) DEFAULT NULL,
  `avatar_file` varchar(64) DEFAULT NULL,
  `username` varchar(64) NOT NULL,
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `employee_no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `is_employee` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_auth_UN_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_roles`
--

DROP TABLE IF EXISTS `user_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_roles` (
  `user_id` bigint DEFAULT NULL,
  `role_id` bigint DEFAULT NULL,
  KEY `user_roles_FK` (`user_id`),
  KEY `user_roles_FK_1` (`role_id`),
  CONSTRAINT `user_roles_FK` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `user_roles_FK_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-01-07  8:54:28
