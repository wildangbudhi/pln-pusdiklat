
LOCK TABLES `activities` WRITE;
/*!40000 ALTER TABLE `activities` DISABLE KEYS */;
INSERT INTO `activities` VALUES 
(1,'Identity','GET','/api/v1/identity/user/:id','/api/v1/identity/user/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(2,'Identity','POST','/api/v1/identity/user/update/:id','/api/v1/identity/user/update/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(3,'Forum','GET','/api/v1/forum/category','/api/v1/forum/category','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(4,'Forum','POST','/api/v1/forum/create','/api/v1/forum/create','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(5,'Forum','POST','/api/v1/forum/update/:forum_id','/api/v1/forum/update/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(6,'Forum','GET','/api/v1/forum/close/:forum_id','/api/v1/forum/close/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(7,'Forum','GET','/api/v1/forum/delete/:forum_id','/api/v1/forum/delete/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(8,'Forum','GET','/api/v1/forum/fetch','/api/v1/forum/fetch','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(9,'Forum','GET','/api/v1/forum/get/:forum_id','/api/v1/forum/get/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(10,'Forum','GET','/api/v1/forum/author/:author_id','/api/v1/forum/author/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(11,'Forum','GET','/api/v1/forum/search','/api/v1/forum/search','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(12,'Forum','GET','/api/v1/forum/react/:forum_id','/api/v1/forum/react/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(13,'Forum','POST','/api/v1/forum/reply/create/:forum_id','/api/v1/forum/reply/create/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(14,'Forum','POST','/api/v1/forum/reply/update/:forum_reply_id','/api/v1/forum/reply/update/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(15,'Forum','GET','/api/v1/forum/reply/delete/:forum_reply_id','/api/v1/forum/reply/delete/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(16,'Forum','GET','/api/v1/forum/reply/react/:forum_reply_id','/api/v1/forum/reply/react/[^/?#]+','2021-01-06 07:55:46','2021-01-06 07:55:46'),
(17,'Forum','GET','/api/v1/forum/get/:forum_id/replies','/api/v1/forum/get/[^/?#]+/replies','2021-01-06 07:55:46','2021-01-06 07:55:46');

/*!40000 ALTER TABLE `activities` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'Client','2021-01-06 06:27:16','2021-01-06 06:27:16'),(2,'Admin','2021-01-06 06:27:16','2021-01-06 06:27:16');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `user_auth` WRITE;
/*!40000 ALTER TABLE `user_auth` DISABLE KEYS */;
INSERT INTO `user_auth` VALUES 
(1,'Rangga Kusuma Dinata',NULL,'0511174000120','$2a$14$JC4.1C0npGNHT8E03/O54.Clq5a/pAthGEnI01wbYFWU8p7KPqaG2','2021-01-06 10:36:35','2021-01-06 10:36:35',NULL,0), 
(2,'Wildan G Budhi',NULL,'0511174000184','$2a$14$JC4.1C0npGNHT8E03/O54.Clq5a/pAthGEnI01wbYFWU8p7KPqaG2','2021-01-06 10:36:35','2021-01-06 10:36:35',NULL,0);
/*!40000 ALTER TABLE `user_auth` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `role_activities` WRITE;
/*!40000 ALTER TABLE `role_activities` DISABLE KEYS */;
INSERT INTO `role_activities` VALUES
(1, 1),
(1, 2),
(1, 3), (2, 3),
(1, 4), (2, 4),
(1, 5), (2, 5),
(1, 6), (2, 6),
(1, 7), (2, 7),
(1, 8), (2, 8),
(1, 9), (2, 9),
(1, 10), (2, 10),
(1, 11), (2, 11),
(1, 12), (2, 12),
(1, 13), (2, 13),
(1, 14), (2, 14),
(1, 15), (2, 15),
(1, 16), (2, 16),
(1, 17), (2, 17);

/*!40000 ALTER TABLE `role_activities` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
INSERT INTO `user_roles` VALUES (1,1),(2,1);
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES 
(1,'Generation'),
(2,'Transmission'),
(3,'Distribution'),
(4,'Commerce & Customer Management'),
(5,'Electricity Equipment Production'),
(6,'Electric Safety, OHS, Security & Environment'),
(7,'Project Management, Engineering & Construction'),
(8,'Research & Development'),
(9,'Learning'),
(10,'Certification'),
(11,'Supply Chain Management'),
(12,'Regulatory and Compliance'),
(13,'Information Technology'),
(14,'SDM'),
(15,'Finance'),
(16,'Communication, CSR & Office Management'),
(17,'Company Management'),
(18,'Miscellaneous');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `forum` WRITE;
/*!40000 ALTER TABLE `forum` DISABLE KEYS */;
INSERT INTO `forum` VALUES ('41c2a806-a2d0-4d6f-9b04-3dfd98dbd441','Bagaimana Cara Membuka Kaleng','Saya sudah beli Sarden Kalengan, Tapi saya tidak bisa membukan nya',2,18,'TERBUKA','2021-01-14 04:02:09','2021-01-14 04:02:09'),('a8948c1a-b47f-4bda-872b-bfd449145379','Bagaimana Cara Memasang Baut','Saya mengalami masalah dalam memperbaiki sepeda saya, ketika ingin memasang baut pada lampu sepeda bautnya tidak bisa di pasang padahal saya sudah memutarnya berlawanan arah jarum jam.',1,18,'TERBUKA','2021-01-14 03:10:22','2021-01-14 03:10:22');
/*!40000 ALTER TABLE `forum` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `forum_reaction` WRITE;
/*!40000 ALTER TABLE `forum_reaction` DISABLE KEYS */;
INSERT INTO `forum_reaction` VALUES (1,'a8948c1a-b47f-4bda-872b-bfd449145379',1,0,'2021-01-14 03:11:05','2021-01-14 03:11:05'),(1,'a8948c1a-b47f-4bda-872b-bfd449145379',1,0,'2021-01-14 03:11:38','2021-01-14 03:11:38'),(2,'a8948c1a-b47f-4bda-872b-bfd449145379',0,1,'2021-01-14 03:10:22','2021-01-14 03:10:22');
/*!40000 ALTER TABLE `forum_reaction` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `forum_replies` WRITE;
/*!40000 ALTER TABLE `forum_replies` DISABLE KEYS */;
INSERT INTO `forum_replies` VALUES ('6463b313-fd26-4b93-bbb0-e7c45f8dcffe','a8948c1a-b47f-4bda-872b-bfd449145379',2,'Loh kan seharusnya di putar searah jarum jam','2021-01-14 03:13:57','2021-01-14 03:13:57'),('be4c0f07-9e55-43cc-bb3a-aa16aafbd8e4','a8948c1a-b47f-4bda-872b-bfd449145379',2,'Saya rasa seharusnya di tarik menggunakan tang','2021-01-14 03:13:57','2021-01-14 03:13:57'),('fe297f2b-9580-422d-ae26-1dba542f45ea','a8948c1a-b47f-4bda-872b-bfd449145379',2,'Saya rasa anda salah dalam cara memasang baut, bukan di putar melain kan di pukul menggunakan palu','2021-01-14 03:13:57','2021-01-14 03:13:57');
/*!40000 ALTER TABLE `forum_replies` ENABLE KEYS */;
UNLOCK TABLES;