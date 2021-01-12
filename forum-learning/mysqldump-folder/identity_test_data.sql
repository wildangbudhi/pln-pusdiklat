USE `identity`;

LOCK TABLES `activities` WRITE;
/*!40000 ALTER TABLE `activities` DISABLE KEYS */;
INSERT INTO `activities` VALUES (1,'API','GET','/api/v1/roles/','/api/v1/roles/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),(2,'API','POST','/api/v2/roles/','/api/v2/roles/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),(3,'API 2','GET','/api/v3/roles/','/api/v3/roles/[^/?#]+','2021-01-06 08:03:48','2021-01-06 08:03:48');
/*!40000 ALTER TABLE `activities` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'Client','2021-01-06 06:27:16','2021-01-06 06:27:16'),(2,'Admin','2021-01-06 06:27:16','2021-01-06 06:27:16');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `user_auth` WRITE;
/*!40000 ALTER TABLE `user_auth` DISABLE KEYS */;
INSERT INTO `user_auth` VALUES (1,'Rangga Kusuma Dinata',NULL,'ranggakd@gmail.com','0511174000120','$2a$14$JC4.1C0npGNHT8E03/O54.Clq5a/pAthGEnI01wbYFWU8p7KPqaG2',NULL,NULL,NULL,NULL,'2021-01-06 10:36:35','2021-01-06 10:36:35',NULL,NULL);
/*!40000 ALTER TABLE `user_auth` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `role_activities` WRITE;
/*!40000 ALTER TABLE `role_activities` DISABLE KEYS */;
INSERT INTO `role_activities` VALUES (1,1),(1,2),(2,3);
/*!40000 ALTER TABLE `role_activities` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
INSERT INTO `user_roles` VALUES (1,1);
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;