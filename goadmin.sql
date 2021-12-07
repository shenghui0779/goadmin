-- MySQL dump 10.13  Distrib 5.7.34, for osx10.16 (x86_64)
--
-- Host: 121.199.68.249    Database: goadmin
-- ------------------------------------------------------
-- Server version	8.0.27

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
-- Table structure for table `attack_email`
--

DROP TABLE IF EXISTS `attack_email`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `attack_email` (
  `create_at` varchar(32) DEFAULT '未知',
  `send` varchar(32) DEFAULT NULL,
  `token` varchar(32) DEFAULT '未知',
  `send_to` varchar(32) DEFAULT '未知',
  `count` varchar(32) DEFAULT '未知',
  `subject` varchar(32) DEFAULT '未知',
  `body` varchar(32) DEFAULT '未知'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `attack_email`
--

LOCK TABLES `attack_email` WRITE;
/*!40000 ALTER TABLE `attack_email` DISABLE KEYS */;
INSERT INTO `attack_email` VALUES ('2021-11-26 17:48:00','2862972664@qq.com','owqjngqteoeqddbh','806459794@qq.com','1','fire','xxxx'),('2021-11-26 17:48:01','2862972664@qq.com','owqjngqteoeqddbh','806459794@qq.com','1','fire','xxxx'),('2021-11-26 17:48:02','2862972664@qq.com','owqjngqteoeqddbh','806459794@qq.com','1','fire','xxxx'),('2021-11-26 17:48:04','2862972664@qq.com','owqjngqteoeqddbh','806459794@qq.com','1','fire','xxxx'),('2021-11-26 18:43:45','2862972664@qq.com','owqjngqteoeqddbh','806459794@qq.com','5','123456','xxxxxxxx'),('2021-11-26 18:45:32','2862972664@qq.com','owqjngqteoeqddbh','806459794@qq.com','3','ignore','你xx'),('2021-11-26 18:47:18','2862972664@qq.com','owqjngqteoeqddbh','806459794@qq.com','2','123','ni xxx'),('2021-11-26 21:26:17','2862972664@qq.com','owqjngqteoeqddbh','806459794@qq.com','2','123','ni xxx'),('2021-11-26 21:51:33','2862972664@qq.com','owqjngqteoeqddbh','2395514919@qq.com','5','xxx','你xxxx'),('2021-11-26 22:07:17','2862972664@qq.com','owqjngqteoeqddbh','2395514919@qq.com','1','你好','哈哈哈哈哈'),('2021-11-26 22:49:46','2862972664@qq.com','owqjngqteoeqddbh','2395514919@qq.com','1','你好','哈哈哈哈哈'),('2021-11-26 22:50:52','2862972664@qq.com','owqjngqteoeqddbh','2395514919@qq.com','1','你好','哈哈哈哈哈'),('2021-11-26 22:51:28','	 2862972664@qq.com','owqjngqteoeqddbh','2395514919@qq.com','1','asdaf','asfdasf'),('2021-11-26 22:51:44','	 2862972664@qq.com','owqjngqteoeqddbh','2395514919@qq.com','1','asdaf','asfdasf'),('2021-11-26 22:52:16','2862972664@qq.com','owqjngqteoeqddbh','2395514919@qq.com','1','asdaf','asfdasf'),('2021-11-26 22:52:43','2862972664@qq.com','owqjngqteoeqddbh','2395514919@qq.com','1','asdaf','asfdasf'),('2021-11-26 22:52:43','2862972664@qq.com','owqjngqteoeqddbh','2395514919@qq.com','1','asdaf','asfdasf'),('2021-11-26 22:53:34','2862972664@qq.com','owqjngqteoeqddbh','2395514919@qq.com','1','asdaf','asfdasf'),('2021-11-30 18:43:13','2862972664@qq.com','owqjngqteoeqddbh','714239774@qq.com','5','茹茹','我好想你啊'),('2021-11-30 18:50:08','2862972664@qq.com','owqjngqteoeqddbh','714239774@qq.com','5','茹茹','我好想你啊');
/*!40000 ALTER TABLE `attack_email` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `db_send_to_her`
--

DROP TABLE IF EXISTS `db_send_to_her`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `db_send_to_her` (
  `id` varchar(191) COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` longtext COLLATE utf8mb4_general_ci,
  `modify_time` longtext COLLATE utf8mb4_general_ci,
  `delete_time` longtext COLLATE utf8mb4_general_ci,
  `code` bigint DEFAULT NULL,
  `message` longtext COLLATE utf8mb4_general_ci,
  `content` longtext COLLATE utf8mb4_general_ci,
  `cover` longtext COLLATE utf8mb4_general_ci,
  `title` longtext COLLATE utf8mb4_general_ci,
  `subtitle` longtext COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `db_send_to_her`
--

LOCK TABLES `db_send_to_her` WRITE;
/*!40000 ALTER TABLE `db_send_to_her` DISABLE KEYS */;
INSERT INTO `db_send_to_her` VALUES ('3c908a04-c35d-4e91-92af-0d532b6d26cd','2021-12-07 17:27:29','','',200,'success','我希望，大家无论通过什么方法，都能挣到足够的钱，去旅行，去闲着，去思考世界的过去和未来，去看书做梦，去街角闲逛，让思绪的钓线深深沉入街流之中。 from 伍尔夫《一间只属于自己的房间》','http://image.wufazhuce.com/Fk3SGcyXWKzHcaiz-PSzezwaROS3','VOL.3344','摄影＆Sarbajit Sen 作品'),('ac70bcd0-cb81-426c-b0e0-aee9b2d6e707','2021-12-07 16:44:58','','',200,'success','我希望，大家无论通过什么方法，都能挣到足够的钱，去旅行，去闲着，去思考世界的过去和未来，去看书做梦，去街角闲逛，让思绪的钓线深深沉入街流之中。 from 伍尔夫《一间只属于自己的房间》','http://image.wufazhuce.com/Fk3SGcyXWKzHcaiz-PSzezwaROS3','VOL.3344','摄影＆Sarbajit Sen 作品'),('e03cc126-c416-4580-8ec8-a3a1a3028b1a','2021-12-07 17:41:03','','',200,'success','我希望，大家无论通过什么方法，都能挣到足够的钱，去旅行，去闲着，去思考世界的过去和未来，去看书做梦，去街角闲逛，让思绪的钓线深深沉入街流之中。 from 伍尔夫《一间只属于自己的房间》','http://image.wufazhuce.com/Fk3SGcyXWKzHcaiz-PSzezwaROS3','VOL.3344','摄影＆Sarbajit Sen 作品'),('ef06eac2-33da-45e1-a9c6-75cd84ec7a56','2021-12-07 17:38:58','','',200,'success','我希望，大家无论通过什么方法，都能挣到足够的钱，去旅行，去闲着，去思考世界的过去和未来，去看书做梦，去街角闲逛，让思绪的钓线深深沉入街流之中。 from 伍尔夫《一间只属于自己的房间》','http://image.wufazhuce.com/Fk3SGcyXWKzHcaiz-PSzezwaROS3','VOL.3344','摄影＆Sarbajit Sen 作品');
/*!40000 ALTER TABLE `db_send_to_her` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `go_user`
--

DROP TABLE IF EXISTS `go_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `go_user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(16) NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(32) NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(256) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(16) NOT NULL DEFAULT '' COMMENT '加密盐',
  `role` tinyint NOT NULL DEFAULT '0' COMMENT '角色，1 - 超级管理员；2 - 高级管理员；3 - 普通管理员',
  `last_login_ip` varchar(20) NOT NULL DEFAULT '' COMMENT '最近登录IP',
  `last_login_time` int unsigned NOT NULL DEFAULT '0' COMMENT '最近登录时间',
  `created_at` int unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `updated_at` int unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3 COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `go_user`
--

LOCK TABLES `go_user` WRITE;
/*!40000 ALTER TABLE `go_user` DISABLE KEYS */;
INSERT INTO `go_user` VALUES (1,'admin','admin@qq.com','wang970425','LCV8xdTcqmkhA$ze',1,'106.75.220.2',1638849437,1509156160,1638846976),(2,'xx','xx@qq.com','123456','LCV8xdTcqmkhA$ze',2,'106.75.220.2',1638785066,1,1638859011);
/*!40000 ALTER TABLE `go_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `weibo`
--

DROP TABLE IF EXISTS `weibo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `weibo` (
  `id` int DEFAULT '0',
  `name` varchar(16) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `nickname` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `uid` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `fans` int DEFAULT '0',
  `follows` int DEFAULT '0',
  `watch` varchar(15) COLLATE utf8mb4_general_ci DEFAULT '是',
  `weibo_count` int DEFAULT '0',
  `description` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `location` varchar(15) COLLATE utf8mb4_general_ci DEFAULT '未知',
  UNIQUE KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `weibo`
--

LOCK TABLES `weibo` WRITE;
/*!40000 ALTER TABLE `weibo` DISABLE KEYS */;
INSERT INTO `weibo` VALUES (1,'狗哥','想不出来名字的小王','5391182376',36,26,'否',18,'土狗不土了','其他'),(0,'谢玉茹','薯片有点困','5731134200',42,33,'是',4,'萬事勝意','上海 杨浦区'),(0,'魏紫怡','dori1122','5091719938',35,46,'是',62,'我一定会爱上遥不可及的你','其他'),(0,'郝美琪','郝多好多鱼','3213576013',206,199,'是',616,'live happy/有事找他@DZ-Mr','陕西'),(0,'杨晨英','逃离半人马','7351105267',37,65,'是',11,'卿卿误我 我误卿卿','四川 成都'),(0,'杨瑞','衰北北','5093156141',35,196,'是',35,'未知','陕西 铜川'),(0,'梁钰铉','努力的Carrie_LM','6305501868',52,222,'是',129,'一起开万人演唱会','陕西 西安'),(0,'汤静','搬砖发财致富','6065121270',146,820,'是',570,'拥有钱钱','其他'),(0,'杨蔓','Villainessss','7394157357',263,72,'是',107,'成长就在一瞬间?','陕西 西安'),(0,'程西亚','Yaya_嘻嘻','3841946319',216,85,'是',728,'勝負何必多慮，人生尤在快意。','江苏');
/*!40000 ALTER TABLE `weibo` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-12-07 17:50:16
