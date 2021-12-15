#
# SQL Export
# Created by Querious (301010)
# Created: December 12, 2021 at 01:14:53 GMT+8
# Encoding: Unicode (UTF-8)
#


SET @ORIG_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_CHECKS = 0;

SET @ORIG_UNIQUE_CHECKS = @@UNIQUE_CHECKS;
SET UNIQUE_CHECKS = 0;

SET @ORIG_TIME_ZONE = @@TIME_ZONE;
SET TIME_ZONE = '+00:00';

SET @ORIG_SQL_MODE = @@SQL_MODE;
SET SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO';



DROP DATABASE IF EXISTS `demo`;
CREATE DATABASE `demo` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
USE `demo`;




DROP TABLE IF EXISTS `user`;


CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `email` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `password` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `salt` varchar(16) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `role` tinyint unsigned NOT NULL DEFAULT '0',
  `registed_at` bigint NOT NULL DEFAULT '0',
  `last_login_at` bigint NOT NULL DEFAULT '0',
  `created_at` bigint NOT NULL DEFAULT '0',
  `updated_at` bigint NOT NULL DEFAULT '0',
  `deleted_at` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;




LOCK TABLES `user` WRITE;
TRUNCATE `user`;
INSERT INTO `user` (`id`, `name`, `email`, `password`, `salt`, `role`, `registed_at`, `last_login_at`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1,'admin','admin@qq.com','e03dcdf34a257041b36bd77132130fdc','LCV8xdTcqmkhA$ze',1,1509156160,1639242790,1509156160,1522049299,0),
	(2,'goadmin','goadmin@qq.com','076f1b2ec2ebfce609805381f2b278b0','oTQKR^BwPgRGM8Zj',2,1509156160,0,1509156160,1509156160,0),
	(3,'test','test@qq.com','9a71c5d96f767e173dfe064ae3120084','n@EdcITV#fQ1d&a@',3,1509156160,0,1509156160,1509156160,0);
UNLOCK TABLES;






SET FOREIGN_KEY_CHECKS = @ORIG_FOREIGN_KEY_CHECKS;

SET UNIQUE_CHECKS = @ORIG_UNIQUE_CHECKS;

SET @ORIG_TIME_ZONE = @@TIME_ZONE;
SET TIME_ZONE = @ORIG_TIME_ZONE;

SET SQL_MODE = @ORIG_SQL_MODE;



# Export Finished: December 12, 2021 at 01:14:53 GMT+8

