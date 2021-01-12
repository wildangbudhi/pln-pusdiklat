USE `forum`;

LOCK TABLES `user_auth` WRITE;
/*!40000 ALTER TABLE `user_auth` DISABLE KEYS */;
INSERT INTO `user_auth` VALUES (1,'Rangga Kusuma Dinata',NULL,'ranggakd@gmail.com','0511174000120','2021-01-06 10:36:35','2021-01-06 10:36:35');
/*!40000 ALTER TABLE `user_auth` ENABLE KEYS */;
UNLOCK TABLES;