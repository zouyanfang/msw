/*
Navicat MySQL Data Transfer

Source Server         : msw
Source Server Version : 50528
Source Host           : localhost:3306
Source Database       : msw

Target Server Type    : MYSQL
Target Server Version : 50528
File Encoding         : 65001

Date: 2018-05-07 22:32:55
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL COMMENT '用户id',
  `pwd` varchar(15) NOT NULL COMMENT '密码',
  `role` int(11) NOT NULL COMMENT '角色 0.超级管理员 1.普通管理员',
  `state` int(11) NOT NULL COMMENT '状态 0.禁用 1.启用',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8 COMMENT='系统用户表';

-- ----------------------------
-- Records of admin
-- ----------------------------

-- ----------------------------
-- Table structure for dish
-- ----------------------------
DROP TABLE IF EXISTS `dish`;
CREATE TABLE `dish` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL COMMENT '用户id',
  `dish_name` varchar(50) NOT NULL COMMENT '菜名',
  `dish_img` varchar(50) DEFAULT NULL COMMENT '菜成品图',
  `release_date` date DEFAULT NULL COMMENT '发布日期',
  `main_material` varchar(50) DEFAULT NULL COMMENT '主料',
  `second_material` varchar(50) DEFAULT NULL COMMENT '辅料',
  `release_role` tinyint(8) DEFAULT NULL COMMENT '发布角色 1.用户 2.官方',
  `tasty` varchar(8) DEFAULT NULL COMMENT '口味',
  `dish_system` varchar(30) DEFAULT NULL COMMENT '菜系',
  `dish_describe` text CHARACTER SET utf8mb4 COMMENT '菜品描述',
  `collect_count` int(11) DEFAULT NULL COMMENT '收藏人数',
  `popular_count` int(11) DEFAULT NULL COMMENT '人气',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8 COMMENT='菜谱信息';

-- ----------------------------
-- Records of dish
-- ----------------------------
INSERT INTO `dish` VALUES ('1001', '141842295', '大闸蟹·', 'static/img/dish_img/1.jpg', '2018-05-01', '大闸蟹 2个', '盐 10g / 葱花 适量', '2', '咸', '闽菜', '海鲜美味大闸蟹，色香味俱全', '12', '34');
INSERT INTO `dish` VALUES ('1002', '141842295', '日本刺身', 'static/img/dish_img/2.jpg', null, null, null, null, null, null, null, null, null);

-- ----------------------------
-- Table structure for dish_comment
-- ----------------------------
DROP TABLE IF EXISTS `dish_comment`;
CREATE TABLE `dish_comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL COMMENT '用户id',
  `dish_id` int(11) NOT NULL COMMENT '菜谱id',
  `content` varchar(50) DEFAULT NULL COMMENT '评论内容',
  `comment_date` datetime DEFAULT NULL COMMENT '评论时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8 COMMENT='菜谱评论表';

-- ----------------------------
-- Records of dish_comment
-- ----------------------------

-- ----------------------------
-- Table structure for dish_step
-- ----------------------------
DROP TABLE IF EXISTS `dish_step`;
CREATE TABLE `dish_step` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dish_id` int(11) NOT NULL COMMENT '菜谱id',
  `step` int(11) NOT NULL COMMENT '步骤',
  `step_img` varchar(50) DEFAULT NULL COMMENT '步骤图',
  `step_describe` varchar(100) DEFAULT NULL COMMENT '步骤描述',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8 COMMENT='菜谱步骤表';

-- ----------------------------
-- Records of dish_step
-- ----------------------------

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL COMMENT '用户id',
  `menu_name` varchar(50) NOT NULL COMMENT '菜单名',
  `menu_img` varchar(50) DEFAULT NULL COMMENT '菜单图',
  `create_date` date DEFAULT NULL COMMENT '创建日期',
  `menu_discribe` text COMMENT '菜单描述',
  `collect_count` int(11) DEFAULT NULL COMMENT '收藏人数',
  `popular_count` int(11) DEFAULT NULL COMMENT '人气',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8 COMMENT='菜单信息';

-- ----------------------------
-- Records of menu
-- ----------------------------

-- ----------------------------
-- Table structure for menu_dish
-- ----------------------------
DROP TABLE IF EXISTS `menu_dish`;
CREATE TABLE `menu_dish` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_id` int(11) NOT NULL COMMENT '菜单id',
  `dish_id` int(11) NOT NULL COMMENT '菜谱id',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8 COMMENT='菜单对应的菜谱表';

-- ----------------------------
-- Records of menu_dish
-- ----------------------------

-- ----------------------------
-- Table structure for notice_board
-- ----------------------------
DROP TABLE IF EXISTS `notice_board`;
CREATE TABLE `notice_board` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL COMMENT '公告标题',
  `content` text NOT NULL COMMENT '公告内容',
  `create_time` datetime DEFAULT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='公告栏表';

-- ----------------------------
-- Records of notice_board
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(50) NOT NULL COMMENT '账户',
  `pwd` varchar(50) NOT NULL COMMENT '密码',
  `name` varchar(50) DEFAULT NULL COMMENT '昵称',
  `user_img` varchar(50) DEFAULT NULL COMMENT '头像',
  `register_date` date DEFAULT NULL COMMENT '注册时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8 COMMENT='注册用户信息';

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', '141842295', '123456', 'lilyzou', './static/img/user_img/1.jpg', '2018-05-06');
INSERT INTO `users` VALUES ('2', '141842261', '111111', 'hxy', './static/img/user_img/2.jpg', '2018-05-05');
INSERT INTO `users` VALUES ('3', '141842268', '222222', '张三', './static/img/user_img/3.jpg', '2018-05-02');

-- ----------------------------
-- Table structure for user_collection
-- ----------------------------
DROP TABLE IF EXISTS `user_collection`;
CREATE TABLE `user_collection` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL COMMENT '用户id',
  `dish_id` int(11) NOT NULL COMMENT '菜谱id',
  `menu_id` int(11) NOT NULL COMMENT '菜单id',
  `collect_time` datetime DEFAULT NULL COMMENT '收藏时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8 COMMENT='用户收藏表';

-- ----------------------------
-- Records of user_collection
-- ----------------------------
