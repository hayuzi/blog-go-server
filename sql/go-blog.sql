/*
 Navicat Premium Data Transfer

 Source Server         : 106.15.91.102
 Source Server Type    : MySQL
 Source Server Version : 50637
 Source Host           : 106.15.91.102:3306
 Source Schema         : go-blog

 Target Server Type    : MySQL
 Target Server Version : 50637
 File Encoding         : 65001

 Date: 13/03/2019 14:11:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '文章标题',
  `sketch` varchar(255) NOT NULL DEFAULT '' COMMENT '文章简述',
  `content` text NOT NULL COMMENT '文章内容',
  `tag_id` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '标签ID',
  `weight` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '文章权重 0-100之间',
  `article_status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '文章状态 1=草稿 2=发布',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_atl_tagid` (`tag_id`),
  KEY `idx_atl_createdat` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='文章表';

-- ----------------------------
-- Records of blog_article
-- ----------------------------
BEGIN;
INSERT INTO `blog_article` VALUES (1, '示例标题', '示例简述', '示例内容', 1, 1, 2, '2019-01-01 00:00:00', '2019-01-01 00:00:00', NULL);
COMMIT;

-- ----------------------------
-- Table structure for blog_comment
-- ----------------------------
DROP TABLE IF EXISTS `blog_comment`;
CREATE TABLE `blog_comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `article_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文章ID',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '评论用户ID',
  `mention_user_id` int(11) NOT NULL COMMENT '被@的用户ID',
  `content` varchar(1024) NOT NULL COMMENT '评论内容',
  `comment_status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1=正常 2=隐藏',
  `created_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '为删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_c_atlid` (`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='文章评论表';

-- ----------------------------
-- Records of blog_comment
-- ----------------------------
BEGIN;
INSERT INTO `blog_comment` VALUES (1, 1, 1, 0, '评论示例', 1, '2019-02-26 17:36:30', '2019-02-26 17:36:35', NULL);
COMMIT;


-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tag_name` varchar(64) NOT NULL DEFAULT '' COMMENT '标签名称',
  `weight` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '排序权重 0-100之间',
  `tag_status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '标签业务状态 1=启用 2=禁用',
  `created_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_at_tagname` (`tag_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='文章标签表';

-- ----------------------------
-- Records of blog_tag
-- ----------------------------
BEGIN;
INSERT INTO `blog_tag` VALUES (1, '示例标签', 10, 1, '2019-01-01 01:16:47', '2019-01-01 01:16:47', NULL);
COMMIT;

-- ----------------------------
-- Table structure for blog_user
-- ----------------------------
DROP TABLE IF EXISTS `blog_user`;
CREATE TABLE `blog_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(64) NOT NULL DEFAULT '' COMMENT '邮箱',
  `pwd` varchar(64) NOT NULL DEFAULT '' COMMENT '用户密码',
  `nickname` varchar(64) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `introduction` varchar(256) NOT NULL DEFAULT '' COMMENT '自我介绍',
  `user_type` tinyint(3) unsigned NOT NULL DEFAULT '2' COMMENT '用户类型 1=管理员 2=前台用户',
  `user_status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '用户业务状态 1=正常 2=封禁',
  `created_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_mobile` (`mobile`),
  KEY `idx_user_uname` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of blog_user   该管理员的默认密码为123456， 部署之后请修改密码
-- ----------------------------
BEGIN;
INSERT INTO `blog_user` VALUES (1, 'admin', '', '', 'bcee8bb121c71ee1645d6d64ec0a15a6', '', '', 1, 1, '2019-03-13 06:09:41', '2019-03-13 06:09:41', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
