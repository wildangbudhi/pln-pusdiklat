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
-- Dumping data for table `activities`
--

LOCK TABLES `activities` WRITE;
/*!40000 ALTER TABLE `activities` DISABLE KEYS */;
INSERT INTO `activities` VALUES (1,'API','GET','/api/v1/roles/','/api/v1/roles/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),(2,'API','POST','/api/v2/roles/','/api/v2/roles/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),(3,'API 2','GET','/api/v3/roles/','/api/v3/roles/[^/?#]+','2021-01-06 08:03:48','2021-01-06 08:03:48');
/*!40000 ALTER TABLE `activities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `discussion`
--

LOCK TABLES `discussion` WRITE;
/*!40000 ALTER TABLE `discussion` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `discussion_expertises`
--

LOCK TABLES `discussion_expertises` WRITE;
/*!40000 ALTER TABLE `discussion_expertises` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_expertises` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `discussion_participants`
--

LOCK TABLES `discussion_participants` WRITE;
/*!40000 ALTER TABLE `discussion_participants` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_participants` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `discussion_participants_waiting_list`
--

LOCK TABLES `discussion_participants_waiting_list` WRITE;
/*!40000 ALTER TABLE `discussion_participants_waiting_list` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_participants_waiting_list` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `discussion_reaction`
--

LOCK TABLES `discussion_reaction` WRITE;
/*!40000 ALTER TABLE `discussion_reaction` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_reaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `discussion_request_expertise_sugestions`
--

LOCK TABLES `discussion_request_expertise_sugestions` WRITE;
/*!40000 ALTER TABLE `discussion_request_expertise_sugestions` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_request_expertise_sugestions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `discussion_requests`
--

LOCK TABLES `discussion_requests` WRITE;
/*!40000 ALTER TABLE `discussion_requests` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_requests` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `discussion_requests_start_datetime_sugestions`
--

LOCK TABLES `discussion_requests_start_datetime_sugestions` WRITE;
/*!40000 ALTER TABLE `discussion_requests_start_datetime_sugestions` DISABLE KEYS */;
/*!40000 ALTER TABLE `discussion_requests_start_datetime_sugestions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `expertises`
--

LOCK TABLES `expertises` WRITE;
/*!40000 ALTER TABLE `expertises` DISABLE KEYS */;
/*!40000 ALTER TABLE `expertises` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `forum`
--

LOCK TABLES `forum` WRITE;
/*!40000 ALTER TABLE `forum` DISABLE KEYS */;
/*!40000 ALTER TABLE `forum` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `forum_reaction`
--

LOCK TABLES `forum_reaction` WRITE;
/*!40000 ALTER TABLE `forum_reaction` DISABLE KEYS */;
/*!40000 ALTER TABLE `forum_reaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `forum_replies`
--

LOCK TABLES `forum_replies` WRITE;
/*!40000 ALTER TABLE `forum_replies` DISABLE KEYS */;
/*!40000 ALTER TABLE `forum_replies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `forum_replies_reactions`
--

LOCK TABLES `forum_replies_reactions` WRITE;
/*!40000 ALTER TABLE `forum_replies_reactions` DISABLE KEYS */;
/*!40000 ALTER TABLE `forum_replies_reactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;
/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `posts_reactions`
--

LOCK TABLES `posts_reactions` WRITE;
/*!40000 ALTER TABLE `posts_reactions` DISABLE KEYS */;
/*!40000 ALTER TABLE `posts_reactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `role_activities`
--

LOCK TABLES `role_activities` WRITE;
/*!40000 ALTER TABLE `role_activities` DISABLE KEYS */;
INSERT INTO `role_activities` VALUES (1,1),(1,2),(2,3);
/*!40000 ALTER TABLE `role_activities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'Client','2021-01-06 06:27:16','2021-01-06 06:27:16'),(2,'Admin','2021-01-06 06:27:16','2021-01-06 06:27:16');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `user_auth`
--

LOCK TABLES `user_auth` WRITE;
/*!40000 ALTER TABLE `user_auth` DISABLE KEYS */;
INSERT INTO `user_auth` VALUES (1,'Rangga Kusuma Dinata',NULL,'ranggakd@gmail.com','0511174000120','$2a$14$JC4.1C0npGNHT8E03/O54.Clq5a/pAthGEnI01wbYFWU8p7KPqaG2',NULL,NULL,NULL,NULL,'2021-01-06 10:36:35','2021-01-06 10:36:35',NULL,NULL);
/*!40000 ALTER TABLE `user_auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `user_roles`
--

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
INSERT INTO `user_roles` VALUES (1,1);
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

-- Dump completed on 2021-01-07  8:28:21
