-- MySQL dump 10.13  Distrib 8.3.0, for Linux (aarch64)
--
-- Host: localhost    Database: scootin_aboot
-- ------------------------------------------------------
-- Server version	8.3.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `scooters`
--

DROP TABLE IF EXISTS `scooters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `scooters` (
  `id` binary(16) NOT NULL DEFAULT (uuid_to_bin(uuid())),
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `latitude` smallint unsigned NOT NULL,
  `longitude` smallint unsigned NOT NULL,
  `location_updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `scooters`
--

LOCK TABLES `scooters` WRITE;
/*!40000 ALTER TABLE `scooters` DISABLE KEYS */;
INSERT INTO `scooters` VALUES (_binary 'yL5\Ä	 \ïžuB¬\0','Scooter 1',2,4,'2024-05-03 07:40:55'),(_binary 'yL?ü	 \ïžuB¬\0','Scooter 2',3,7,'2024-05-03 07:40:55'),(_binary 'yLB[	 \ïžuB¬\0','Scooter 3',3,9,'2024-05-03 07:40:55'),(_binary 'yLB\Ô	 \ïžuB¬\0','Scooter 4',5,15,'2024-05-03 07:40:55'),(_binary 'yLC)	 \ïžuB¬\0','Scooter 5',6,5,'2024-05-03 07:40:55'),(_binary 'yLD{	 \ïžuB¬\0','Scooter 6',7,11,'2024-05-03 07:40:55'),(_binary 'yLD\é	 \ïžuB¬\0','Scooter 7',9,14,'2024-05-03 07:40:55'),(_binary 'yLE.	 \ïžuB¬\0','Scooter 8',11,11,'2024-05-03 07:40:55'),(_binary 'yLE\ð	 \ïžuB¬\0','Scooter 9',11,18,'2024-05-03 07:40:55'),(_binary 'yLF\Â	 \ïžuB¬\0','Scooter 10',13,2,'2024-05-03 07:40:55'),(_binary 'yLGM	 \ïžuB¬\0','Scooter 11',14,13,'2024-05-03 07:40:55'),(_binary 'yLG°	 \ïžuB¬\0','Scooter 12',17,17,'2024-05-03 07:40:55'),(_binary 'yLG\ó	 \ïžuB¬\0','Scooter 13',18,5,'2024-05-03 07:40:55'),(_binary 'yLH?	 \ïžuB¬\0','Scooter 14',18,9,'2024-05-03 07:40:55'),(_binary 'yLH}	 \ïžuB¬\0','Scooter 15',19,3,'2024-05-03 07:40:55');
/*!40000 ALTER TABLE `scooters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `scooters_occupations`
--

DROP TABLE IF EXISTS `scooters_occupations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `scooters_occupations` (
  `id` binary(16) NOT NULL DEFAULT (uuid_to_bin(uuid())),
  `scooter_id` binary(16) NOT NULL,
  `user_id` binary(16) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_scooter_id` (`scooter_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `scooters_occupations`
--

LOCK TABLES `scooters_occupations` WRITE;
/*!40000 ALTER TABLE `scooters_occupations` DISABLE KEYS */;
/*!40000 ALTER TABLE `scooters_occupations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` binary(16) NOT NULL DEFAULT (uuid_to_bin(uuid())),
  `first_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `api_key` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (_binary 'r\Íˆ	 \ïžuB¬\0','Bob','Martin','yLxCMVd*p9hZNnvYfGxyQezPBE3@8Lp9sgKhE9QwtQMM!v6y%4Q&UuqJjC'),(_binary 'r\Ím	 \ïžuB¬\0','Forrest','Gump','rcYPJRUGf3xv&w4ny%Qhs1^&LJN@&T@H83srvheyYP&XJXs^S@sU1QFQHv'),(_binary 'r\Ía	 \ïžuB¬\0','Whitney','Houston','3JFj84qXE4^18jfbh2vfBvu3Ev4DrsQ*xrXJg7N5dWEgT2XmTEx%EQ4D8U');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-05-03  7:41:28
