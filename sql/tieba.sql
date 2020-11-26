/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : localhost:5664
 Source Schema         : tieba

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 24/11/2020 17:09:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_admin
-- ----------------------------
DROP TABLE IF EXISTS `tb_admin`;
CREATE TABLE `tb_admin`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'uuid',
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `pass` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  `operation_at` datetime(0) NULL DEFAULT NULL COMMENT '最后一次操作时间',
  `operation` int(2) NULL DEFAULT NULL COMMENT '最后一次的操作',
  `operation_text` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '最后一次的操作说明',
  PRIMARY KEY (`id`, `uuid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tb_article
-- ----------------------------
DROP TABLE IF EXISTS `tb_article`;
CREATE TABLE `tb_article`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'uuid',
  `club_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '贴吧的uuid',
  `user_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建者的uuid',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '内容（存储富文本， 理论上不限制大小）',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '发布时间',
  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  `creater` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `updater` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '修改人',
  `open_num` int(11) NULL DEFAULT NULL COMMENT '打开次数',
  `description` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '描述',
  `is_good` int(2) NULL DEFAULT 0 COMMENT '是否精品 1是， 0否',
  `flag` int(2) NULL DEFAULT 1 COMMENT '是否展示 1是 0 否',
  `is_money` int(2) NULL DEFAULT 0 COMMENT '是否需要付费 1 是 0 否',
  `money` int(11) NULL DEFAULT NULL COMMENT '观看所需金币数',
  PRIMARY KEY (`id`, `uuid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '帖子表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tb_bulletin
-- ----------------------------
DROP TABLE IF EXISTS `tb_bulletin`;
CREATE TABLE `tb_bulletin`  (
  `id` int(11) NOT NULL COMMENT '主键 ',
  `title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '内容',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `create_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建的管理员uuid',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '电子公告栏表， 比较独立' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tb_club
-- ----------------------------
DROP TABLE IF EXISTS `tb_club`;
CREATE TABLE `tb_club`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'uuid',
  `master_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人用户名',
  `master_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人的用户id',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '贴吧名称',
  `head_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '贴吧头像',
  `description` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '介绍',
  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '类型， 用逗号分隔开',
  `type_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '类型名称， 也是用逗号分隔开',
  `attention_num` int(11) NULL DEFAULT NULL COMMENT '关注数',
  `article_num` int(11) NULL DEFAULT NULL COMMENT '帖子数',
  `member_num` int(11) NULL DEFAULT NULL COMMENT '成员数',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  `creater` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `updater` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '修改人',
  PRIMARY KEY (`id`, `uuid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '贴吧' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tb_comment
-- ----------------------------
DROP TABLE IF EXISTS `tb_comment`;
CREATE TABLE `tb_comment`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'uuid',
  `p_id` int(11) NULL DEFAULT NULL COMMENT '父级id， 指向父级评论，当自己就是一级评论的时候等于自身id',
  `p_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '父级uuid， 可以为null',
  `type` int(2) NULL DEFAULT 1 COMMENT '1就是一级评论， 2就是评论， 默认一级评论',
  `user_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户id',
  `article_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '帖子id',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '内容',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '评论时间',
  `state` int(2) NULL DEFAULT 0 COMMENT '举报状态， 0 没有被举报 1 已被举报 2举报审核通过 ',
  PRIMARY KEY (`id`, `uuid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '主业务-评论表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tb_dict
-- ----------------------------
DROP TABLE IF EXISTS `tb_dict`;
CREATE TABLE `tb_dict`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type` int(11) NULL DEFAULT NULL COMMENT '类型',
  `type_name` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '类型名称',
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '名称',
  `pid` int(11) NULL DEFAULT NULL COMMENT '父级id',
  `level` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '等级',
  `order` int(11) NULL DEFAULT NULL COMMENT '排序规则',
  `other` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '未定',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '字典表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tb_gold
-- ----------------------------
DROP TABLE IF EXISTS `tb_gold`;
CREATE TABLE `tb_gold`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'uuid',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `user_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户的uuid',
  `glod_num` int(11) NULL DEFAULT NULL COMMENT '获得的金币数',
  `gold_type` int(2) NULL DEFAULT NULL COMMENT '获取金币的类型',
  PRIMARY KEY (`id`, `uuid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tb_like
-- ----------------------------
DROP TABLE IF EXISTS `tb_like`;
CREATE TABLE `tb_like`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'uuid',
  `user_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户的uuid',
  `article_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '收藏的帖子uuid',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`, `uuid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '收藏表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tb_top
-- ----------------------------
DROP TABLE IF EXISTS `tb_top`;
CREATE TABLE `tb_top`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `description` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '热榜说明',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `num` int(11) NULL DEFAULT NULL COMMENT '期数',
  `article_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '帖子的uuid',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '历史热榜表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'uuid',
  `p_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '父级uuid',
  `user_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `pass` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  `sex` int(2) NULL DEFAULT NULL COMMENT '性别 1男 0女',
  `nick` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `phone` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `experience` int(11) NULL DEFAULT NULL COMMENT '经验值',
  `level` int(11) NULL DEFAULT NULL COMMENT '等级',
  `gold` int(11) NULL DEFAULT NULL COMMENT '金币数',
  `flag` int(2) NULL DEFAULT 1 COMMENT '代表是否有效1是0否',
  `description` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '个人签名',
  `head_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
  `birthday` date NULL DEFAULT NULL COMMENT '生日',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地址',
  `type` int(2) NULL DEFAULT NULL COMMENT '用户类型',
  `modify` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '修改人',
  `modify_at` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  `creater` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `age` int(11) NULL DEFAULT NULL COMMENT '吧龄',
  `post_num` int(11) NULL DEFAULT NULL COMMENT '发帖数',
  PRIMARY KEY (`id`, `uuid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tb_user_club
-- ----------------------------
DROP TABLE IF EXISTS `tb_user_club`;
CREATE TABLE `tb_user_club`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'uuid',
  `club_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '帖子的uuid',
  `user_uuid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户的uuid',
  `identity` int(11) NULL DEFAULT 0 COMMENT '代表身份， 1吧主  0成员， 默认成员',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '贴吧成员表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
