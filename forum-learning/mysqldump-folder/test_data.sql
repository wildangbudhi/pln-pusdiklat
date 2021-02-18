
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
INSERT INTO `user_auth` VALUES (1,'Rangga Kusuma Dinata',NULL,'ranggakd@gmail.com','0511174000120','$2a$14$JC4.1C0npGNHT8E03/O54.Clq5a/pAthGEnI01wbYFWU8p7KPqaG2',NULL,NULL,NULL,NULL,'2021-01-06 10:36:35','2021-01-06 10:36:35',NULL,NULL), (2,'Wildan G Budhi',NULL,'wildangbudhi@gmail.com','0511174000184','$2a$14$JC4.1C0npGNHT8E03/O54.Clq5a/pAthGEnI01wbYFWU8p7KPqaG2',NULL,NULL,NULL,NULL,'2021-01-06 10:36:35','2021-01-06 10:36:35',NULL,NULL);
/*!40000 ALTER TABLE `user_auth` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `role_activities` WRITE;
/*!40000 ALTER TABLE `role_activities` DISABLE KEYS */;
INSERT INTO `role_activities` VALUES (1,1),(1,2),(2,3);
/*!40000 ALTER TABLE `role_activities` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
INSERT INTO `user_roles` VALUES (1,1),(2,1);
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1,'Pembangkitan'),(2,'Transmisi'),(3,'Distribusi'),(4,'Niaga dan Manajemen Pelanggan'),(5,'Produksi Peralatan Ketenagalistrikan'),(6,'K2, K3, Keamanan dan Lingkungan'),(7,'Manajemen Proyek, Enjiniring (Engineering) dan Konstruksi'),(8,'Penelitian dan Pengembangan'),(9,'Pembelajaran'),(10,'Sertifikasi'),(11,'Supply Chain Management'),(12,'Regulatory and Compliance'),(13,'Teknologi Informasi'),(14,'SDM'),(15,'Keuangan'),(16,'Komunikasi, CSR dan Pengelolaan Kantor'),(17,'Manajemen Perusahaan'),(18,'Bebas');
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