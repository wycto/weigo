/*
 Navicat MySQL Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 127.0.0.1:3306
 Source Schema         : weigo

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 27/11/2020 09:16:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for wei_article
-- ----------------------------
DROP TABLE IF EXISTS `wei_article`;
CREATE TABLE `wei_article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `content` text,
  `status` tinyint(2) unsigned DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of wei_article
-- ----------------------------
BEGIN;
INSERT INTO `wei_article` VALUES (1, 'weigo框架的使用', 'weigo框架是程序爱好者weiyi打造，提供学习使用和工作的简单易学的框架，实用性很强', 1);
COMMIT;

-- ----------------------------
-- Table structure for wei_user
-- ----------------------------
DROP TABLE IF EXISTS `wei_user`;
CREATE TABLE `wei_user` (
  `uid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `qq` varchar(255) DEFAULT NULL,
  `weixin` varchar(255) DEFAULT NULL,
  `status` tinyint(2) unsigned DEFAULT '1',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of wei_user
-- ----------------------------
BEGIN;
INSERT INTO `wei_user` VALUES (1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', 'weiyi', '唯一', '294287600@qq.com', '294287600', '294287600', 1);
INSERT INTO `wei_user` VALUES (2, 'weigo', 'e10adc3949ba59abbe56e057f20f883e', 'weigo', 'weigo', '5665156156@qq.com', '41561561', '1651651', 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
