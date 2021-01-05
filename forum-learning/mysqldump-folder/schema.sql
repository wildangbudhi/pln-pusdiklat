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
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `activities`
--

LOCK TABLES `activities` WRITE;
/*!40000 ALTER TABLE `activities` DISABLE KEYS */;
/*!40000 ALTER TABLE `activities` ENABLE KEYS */;
UNLOCK TABLES;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `discussion`
--

DROP TABLE IF EXISTS `discussion`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discussion` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `forum_id` bigint NOT NULL,
  `title` varchar(100) DEFAULT NULL,
  `zoom_invitation_url` varchar(100) DEFAULT NULL,
  `zoom_host_email` varchar(100) DEFAULT NULL,
  `zoom_host_password` varchar(100) DEFAULT NULL,
  `start_datetime` timestamp NULL DEFAULT NULL,
  `end_datetime` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discussion`
--

LOCK TABLES `discussion` WRITE;
/*!40000 ALTER TABLE `discussion` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `discussion_expertises`
--

DROP TABLE IF EXISTS `discussion_expertises`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discussion_expertises` (
  `expertise_id` bigint DEFAULT NULL,
  `discussion_id` bigint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `discussion_expertises_FK_user_id` (`expertise_id`),
  KEY `discussion_expertises_FK` (`discussion_id`),
  CONSTRAINT `discussion_expertises_FK` FOREIGN KEY (`discussion_id`) REFERENCES `discussion` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `discussion_expertises_FK_expertise_id` FOREIGN KEY (`expertise_id`) REFERENCES `expertises` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discussion_expertises`
--

LOCK TABLES `discussion_expertises` WRITE;
/*!40000 ALTER TABLE `discussion_expertises` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_expertises` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `discussion_participants`
--

DROP TABLE IF EXISTS `discussion_participants`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discussion_participants` (
  `user_id` bigint DEFAULT NULL,
  `discussion_id` bigint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `discussion_participants_FK_user_id` (`user_id`),
  KEY `discussion_participants_FK_discussion_id` (`discussion_id`),
  CONSTRAINT `discussion_participants_FK_discussion_id` FOREIGN KEY (`discussion_id`) REFERENCES `discussion` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `discussion_participants_FK_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discussion_participants`
--

LOCK TABLES `discussion_participants` WRITE;
/*!40000 ALTER TABLE `discussion_participants` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_participants` ENABLE KEYS */;
UNLOCK TABLES;

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
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `discussion_participants_waiting_list_FK` (`discussion_id`),
  KEY `discussion_participants_waiting_list_FK_1` (`user_id`),
  CONSTRAINT `discussion_participants_waiting_list_FK` FOREIGN KEY (`discussion_id`) REFERENCES `discussion` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `discussion_participants_waiting_list_FK_1` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discussion_participants_waiting_list`
--

LOCK TABLES `discussion_participants_waiting_list` WRITE;
/*!40000 ALTER TABLE `discussion_participants_waiting_list` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_participants_waiting_list` ENABLE KEYS */;
UNLOCK TABLES;

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
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `discussion_reaction_FK_user_id` (`user_id`),
  KEY `discussion_reaction_FK_discussion_id` (`discussion`),
  CONSTRAINT `discussion_reaction_FK_discussion_id` FOREIGN KEY (`discussion`) REFERENCES `discussion` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `discussion_reaction_FK_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discussion_reaction`
--

LOCK TABLES `discussion_reaction` WRITE;
/*!40000 ALTER TABLE `discussion_reaction` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_reaction` ENABLE KEYS */;
UNLOCK TABLES;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discussion_request_expertise_sugestions`
--

LOCK TABLES `discussion_request_expertise_sugestions` WRITE;
/*!40000 ALTER TABLE `discussion_request_expertise_sugestions` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_request_expertise_sugestions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `discussion_requests`
--

DROP TABLE IF EXISTS `discussion_requests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discussion_requests` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `forum_id` bigint DEFAULT NULL,
  `accepted` tinyint(1) DEFAULT NULL,
  `accepted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `reqeuster_user_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `discussion_requests_FK` (`forum_id`),
  KEY `discussion_requests_FK_1` (`reqeuster_user_id`),
  CONSTRAINT `discussion_requests_FK` FOREIGN KEY (`forum_id`) REFERENCES `forum` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `discussion_requests_FK_1` FOREIGN KEY (`reqeuster_user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discussion_requests`
--

LOCK TABLES `discussion_requests` WRITE;
/*!40000 ALTER TABLE `discussion_requests` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_requests` ENABLE KEYS */;
UNLOCK TABLES;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discussion_requests_start_datetime_sugestions`
--

LOCK TABLES `discussion_requests_start_datetime_sugestions` WRITE;
/*!40000 ALTER TABLE `discussion_requests_start_datetime_sugestions` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_requests_start_datetime_sugestions` ENABLE KEYS */;
UNLOCK TABLES;

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
  PRIMARY KEY (`id`),
  KEY `expertises_FK_category_id` (`category_id`),
  CONSTRAINT `expertises_FK_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `expertises`
--

LOCK TABLES `expertises` WRITE;
/*!40000 ALTER TABLE `expertises` DISABLE KEYS */;
/*!40000 ALTER TABLE `expertises` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `forum`
--

DROP TABLE IF EXISTS `forum`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `forum` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `author_user_id` bigint NOT NULL,
  `category_id` bigint NOT NULL,
  `status` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `forum_FK_author_user_id` (`author_user_id`),
  KEY `forum_FK_category_id` (`category_id`),
  CONSTRAINT `forum_FK_author_user_id` FOREIGN KEY (`author_user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `forum_FK_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `forum`
--

LOCK TABLES `forum` WRITE;
/*!40000 ALTER TABLE `forum` DISABLE KEYS */;
/*!40000 ALTER TABLE `forum` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `forum_reaction`
--

DROP TABLE IF EXISTS `forum_reaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `forum_reaction` (
  `user_id` bigint DEFAULT NULL,
  `forum_id` bigint DEFAULT NULL,
  `up_vote` tinyint(1) DEFAULT NULL,
  `down_vote` tinyint(1) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `forum_reaction_FK_user_id` (`user_id`),
  KEY `forum_reaction_FK` (`forum_id`),
  CONSTRAINT `forum_reaction_FK` FOREIGN KEY (`forum_id`) REFERENCES `forum` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `forum_reaction_FK_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `forum_reaction`
--

LOCK TABLES `forum_reaction` WRITE;
/*!40000 ALTER TABLE `forum_reaction` DISABLE KEYS */;
/*!40000 ALTER TABLE `forum_reaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `forum_replies`
--

DROP TABLE IF EXISTS `forum_replies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `forum_replies` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `forum_id` bigint NOT NULL,
  `author_user_id` bigint NOT NULL,
  `answer` varchar(10000) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `forum_replies_FK_forum_id` (`forum_id`),
  KEY `forum_replies_FK` (`author_user_id`),
  CONSTRAINT `forum_replies_FK` FOREIGN KEY (`author_user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `forum_replies_FK_forum_id` FOREIGN KEY (`forum_id`) REFERENCES `forum` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `forum_replies`
--

LOCK TABLES `forum_replies` WRITE;
/*!40000 ALTER TABLE `forum_replies` DISABLE KEYS */;
/*!40000 ALTER TABLE `forum_replies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `forum_replies_reactions`
--

DROP TABLE IF EXISTS `forum_replies_reactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `forum_replies_reactions` (
  `user_id` bigint NOT NULL,
  `forum_id` bigint NOT NULL,
  `forum_replies_id` bigint NOT NULL,
  `up_vote` tinyint(1) DEFAULT NULL,
  `down_vote` tinyint(1) DEFAULT NULL,
  `agree` tinyint(1) DEFAULT NULL,
  `skeptic` tinyint(1) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `forum_replies_reactions_FK_user_id` (`user_id`),
  KEY `forum_replies_reactions_FK_forum_id` (`forum_id`),
  KEY `forum_replies_reactions_FKforum_replies_id` (`forum_replies_id`),
  CONSTRAINT `forum_replies_reactions_FK_forum_id` FOREIGN KEY (`forum_id`) REFERENCES `forum` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `forum_replies_reactions_FK_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `forum_replies_reactions_FKforum_replies_id` FOREIGN KEY (`forum_replies_id`) REFERENCES `forum_replies` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `forum_replies_reactions`
--

LOCK TABLES `forum_replies_reactions` WRITE;
/*!40000 ALTER TABLE `forum_replies_reactions` DISABLE KEYS */;
/*!40000 ALTER TABLE `forum_replies_reactions` ENABLE KEYS */;
UNLOCK TABLES;

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
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `posts_FK_discussion_id` (`discussion_id`),
  KEY `posts_FK_author_user_id` (`author_user_id`),
  KEY `posts_FK_category_id` (`category_id`),
  CONSTRAINT `posts_FK_author_user_id` FOREIGN KEY (`author_user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `posts_FK_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `posts_FK_discussion_id` FOREIGN KEY (`discussion_id`) REFERENCES `discussion` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;
/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;

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
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `posts_reactions_FK_user_id` (`user_id`),
  KEY `posts_reactions_FK_posts_id` (`posts_id`),
  CONSTRAINT `posts_reactions_FK_posts_id` FOREIGN KEY (`posts_id`) REFERENCES `posts` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `posts_reactions_FK_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posts_reactions`
--

LOCK TABLES `posts_reactions` WRITE;
/*!40000 ALTER TABLE `posts_reactions` DISABLE KEYS */;
/*!40000 ALTER TABLE `posts_reactions` ENABLE KEYS */;
UNLOCK TABLES;

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
  CONSTRAINT `role_activities_FK` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `role_activities_FK_1` FOREIGN KEY (`activities_id`) REFERENCES `activities` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_activities`
--

LOCK TABLES `role_activities` WRITE;
/*!40000 ALTER TABLE `role_activities` DISABLE KEYS */;
/*!40000 ALTER TABLE `role_activities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `roles_name` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

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
  `email` varchar(255) NOT NULL,
  `username` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `status_id` tinyint DEFAULT NULL,
  `user_type_id` tinyint DEFAULT NULL,
  `user_entity_id` bigint DEFAULT NULL,
  `user_key` varchar(64) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `login_at` timestamp NULL DEFAULT NULL,
  `modified_by` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_auth_UN_email` (`email`),
  UNIQUE KEY `user_auth_UN_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_auth`
--

LOCK TABLES `user_auth` WRITE;
/*!40000 ALTER TABLE `user_auth` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_auth` ENABLE KEYS */;
UNLOCK TABLES;

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
  CONSTRAINT `user_roles_FK` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `user_roles_FK_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_roles`
--

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-01-05 19:41:48
