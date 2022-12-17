-- MySQL dump 10.13  Distrib 5.7.29, for osx10.15 (x86_64)
--
-- Host: 127.0.0.1    Database: activities
-- ------------------------------------------------------
-- Server version	5.7.35

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
-- Table structure for table `activity`
--

DROP TABLE IF EXISTS `activity`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `activity` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '活动ID',
  `planner` int(11) unsigned NOT NULL COMMENT '组织者ID',
  `group_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '群组ID',
  `kind` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '活动分类:1羽毛球,2篮球,3足球,4聚餐...',
  `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '活动类型:1全局保护,2全局公开,3群组',
  `title` varchar(256) CHARACTER SET utf8mb4 NOT NULL COMMENT '活动标题',
  `remark` text CHARACTER SET utf8mb4 COMMENT '活动描述',
  `quota` smallint(5) unsigned NOT NULL DEFAULT '1' COMMENT '名额',
  `waiting` smallint(5) unsigned NOT NULL DEFAULT '1' COMMENT '候补数量限制',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '活动状态:1进行中,2正常结算完成,3手动终止',
  `fee_type` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '结算方式:1免费,2活动前,3活动后男女平均,4活动后男固定|女平摊,5活动后男平摊|女固定',
  `fee_male` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '男费用,单位:分',
  `fee_female` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '女费用,单位:分',
  `addr` varchar(256) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '活动地址',
  `ahead` tinyint(4) NOT NULL COMMENT '可提前取消时间(小时)',
  `begin_at` datetime NOT NULL COMMENT '开始时间',
  `end_at` datetime NOT NULL COMMENT '结束时间',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_begin_end` (`begin_at`,`end_at`),
  UNIQUE KEY `idx_group_kind_type_status` (`group_id`,`kind`,`type`,`status`),
  UNIQUE KEY `idx_planner_kind_type_status` (`planner`,`kind`,`type`,`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `activity`
--

LOCK TABLES `activity` WRITE;
/*!40000 ALTER TABLE `activity` DISABLE KEYS */;
/*!40000 ALTER TABLE `activity` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `activity_user`
--

DROP TABLE IF EXISTS `activity_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `activity_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '活动报名id',
  `activity_id` int(11) unsigned NOT NULL COMMENT '活动id',
  `user_id` int(11) unsigned NOT NULL COMMENT '报名用户id',
  `alias` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '报名昵称',
  `self` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否带朋友',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `delete_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_activity_user` (`activity_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动报名表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `activity_user`
--

LOCK TABLES `activity_user` WRITE;
/*!40000 ALTER TABLE `activity_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `activity_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group`
--

DROP TABLE IF EXISTS `group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '群组ID',
  `scores` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '积分',
  `level` tinyint(4) NOT NULL DEFAULT '1' COMMENT '群组等级',
  `name` varchar(128) CHARACTER SET utf8mb4 NOT NULL COMMENT '群组名称',
  `logo` varchar(256) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '群组LOGO',
  `notice` text CHARACTER SET utf8mb4 COMMENT '群组公告',
  `addr` varchar(256) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '群组总部地址',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`),
  KEY `idx_level` (`level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='群组表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group`
--

LOCK TABLES `group` WRITE;
/*!40000 ALTER TABLE `group` DISABLE KEYS */;
/*!40000 ALTER TABLE `group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_application`
--

DROP TABLE IF EXISTS `group_application`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_application` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '申请id',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `group_id` int(11) unsigned NOT NULL COMMENT '群组id',
  `deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `delete_at` datetime DEFAULT NULL COMMENT '删除id',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_group` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='群组申请表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_application`
--

LOCK TABLES `group_application` WRITE;
/*!40000 ALTER TABLE `group_application` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_application` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_member`
--

DROP TABLE IF EXISTS `group_member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_member` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '成员id',
  `group_id` int(11) unsigned NOT NULL COMMENT '群组id',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `scores` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '积分',
  `position` tinyint(4) NOT NULL DEFAULT '1' COMMENT '群组职位',
  `alias` varchar(50) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '群组中别名',
  `avatar` varchar(128) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '头像',
  `enter_at` datetime DEFAULT NULL COMMENT '进入群组时间',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_group_user` (`group_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='群组成员表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_member`
--

LOCK TABLES `group_member` WRITE;
/*!40000 ALTER TABLE `group_member` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(128) CHARACTER SET utf8mb4 NOT NULL COMMENT '用户名',
  `password` varchar(128) CHARACTER SET utf8mb4 NOT NULL DEFAULT '123456' COMMENT '密码',
  `scores` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '积分',
  `token` varchar(128) CHARACTER SET utf8mb4 NOT NULL COMMENT 'token',
  `wx_token` varchar(128) CHARACTER SET utf8mb4 NOT NULL COMMENT '微信session_key',
  `wx_nick` varchar(256) CHARACTER SET utf8mb4 NOT NULL COMMENT '微信昵称',
  `nick` varchar(256) CHARACTER SET utf8mb4 NOT NULL COMMENT '昵称',
  `sex` tinyint(2) NOT NULL DEFAULT '1' COMMENT '性别:1男,2女',
  `phone` varchar(11) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '手机号码',
  `email` varchar(128) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '邮箱',
  `ip` varchar(128) CHARACTER SET utf8mb4 DEFAULT '0' COMMENT 'ip地址',
  `login_at` datetime DEFAULT NULL COMMENT '登录时间',
  `offline_at` datetime DEFAULT NULL COMMENT '离开时间',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name_pass` (`username`,`password`),
  UNIQUE KEY `idx_name_token` (`username`,`token`),
  UNIQUE KEY `idx_wxname_wxtoken` (`wx_nick`,`wx_token`),
  UNIQUE KEY `idx_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-17 16:11:37
