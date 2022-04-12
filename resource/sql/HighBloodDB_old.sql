/*
 Navicat Premium Data Transfer

 Source Server         : BloodPressureMeasurementRecordApplet
 Source Server Type    : MariaDB
 Source Server Version : 50568
 Source Host           : 1.117.222.119:3306
 Source Schema         : HighBloodDB

 Target Server Type    : MariaDB
 Target Server Version : 50568
 File Encoding         : 65001

 Date: 10/04/2022 21:00:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE `admin_users`  (
  `admin_id` int(32) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '管理员Id',
  `admin_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '管理员名称',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '管理员密码',
  `tel` char(11) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '电话号码',
  `email` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '邮件',
  `permission` tinyint(2) NULL DEFAULT NULL COMMENT '权限',
  `last_time` datetime(0) NULL DEFAULT NULL COMMENT '上次登陆时间',
  `create_time` datetime(0) NOT NULL COMMENT '创建日期',
  `sex` enum('男','女','其他') CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '性别',
  `status` enum('开启','关闭') CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '开启' COMMENT '状态',
  PRIMARY KEY (`admin_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for base_users
-- ----------------------------
DROP TABLE IF EXISTS `base_users`;
CREATE TABLE `base_users`  (
  `user_id` int(32) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '用户在系统内的ID',
  `open_id` char(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '微信用户openid',
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '用户名称',
  `real_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '用户真实姓名',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '用户密码 - 预留了登录接口，但是不使用（给可能出现的Web版本预留）',
  `tel` char(11) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '用户电话号码',
  `email` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '用户预留的邮箱',
  `last_time` datetime(0) NULL DEFAULT NULL COMMENT '用户最近的登录时间',
  `create_time` datetime(0) NOT NULL COMMENT '该账户的注册时间',
  `birthday` date NULL DEFAULT NULL COMMENT '用户的出生年月日，用来计算用户的生日、年龄',
  `sex` enum('男','女','其他') CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '其他' COMMENT '注册用户性别，可选项为“男”，“女”，“其他”',
  `status` enum('开启','关闭') CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '开启' COMMENT '该账户使用状态，可选项为“开启”，“关闭”',
  PRIMARY KEY (`user_id`) USING BTREE,
  INDEX `userid`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1039 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for medical_reports
-- ----------------------------
DROP TABLE IF EXISTS `medical_reports`;
CREATE TABLE `medical_reports`  (
  `id` int(32) NOT NULL COMMENT '报告id',
  `user_id` int(32) UNSIGNED ZEROFILL NOT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_report_id`(`user_id`) USING BTREE,
  CONSTRAINT `user_report_id` FOREIGN KEY (`user_id`) REFERENCES `base_users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for medical_supplies
-- ----------------------------
DROP TABLE IF EXISTS `medical_supplies`;
CREATE TABLE `medical_supplies`  (
  `medical_id` int(32) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '药品使用id',
  `metering` int(32) NULL DEFAULT NULL COMMENT '药品使用计量',
  `measuring` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '药品计量单位，如克、千克、毫克、毫升、升等',
  PRIMARY KEY (`medical_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for patient_bp_records
-- ----------------------------
DROP TABLE IF EXISTS `patient_bp_records`;
CREATE TABLE `patient_bp_records`  (
  `record_id` int(32) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '血压记录id',
  `user_id` int(32) UNSIGNED ZEROFILL NOT NULL COMMENT '用户id-外键',
  `record_date` date NULL DEFAULT NULL COMMENT '用户记录日期',
  `record_time` time(0) NULL DEFAULT NULL COMMENT '用户记录时间',
  `low_pressure` int(32) NULL DEFAULT NULL COMMENT '测量患者血压-低压',
  `high_pressure` int(32) NULL DEFAULT NULL COMMENT '测量患者血压-高压',
  PRIMARY KEY (`record_id`) USING BTREE,
  INDEX `user_recordbp_id`(`user_id`) USING BTREE,
  CONSTRAINT `user_recordbp_id` FOREIGN KEY (`user_id`) REFERENCES `base_users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1038 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for treatment_plans
-- ----------------------------
DROP TABLE IF EXISTS `treatment_plans`;
CREATE TABLE `treatment_plans`  (
  `treatment_id` int(32) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '治疗方案id',
  `user_id` int(32) UNSIGNED ZEROFILL NOT NULL COMMENT '患者id(即用户id)',
  `plan` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '治疗计划（1天/一次、2天/一次、1天/三次……）',
  `note` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '治疗方案备注，如禁忌症',
  `create_date` date NULL DEFAULT NULL COMMENT '方案创建日期',
  `create_time` time(0) NULL DEFAULT NULL COMMENT '方案创建时间',
  `status` enum('生效','失效') CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '生效' COMMENT '该方案状态（“生效”，“失效”）',
  `end_date` date NULL DEFAULT NULL COMMENT '方案结束日期',
  `end_time` time(0) NULL DEFAULT NULL COMMENT '方案结束时间',
  PRIMARY KEY (`treatment_id`) USING BTREE,
  INDEX `user_treatment_id`(`user_id`) USING BTREE,
  CONSTRAINT `user_treatment_id` FOREIGN KEY (`user_id`) REFERENCES `base_users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for treatment_use_medicals
-- ----------------------------
DROP TABLE IF EXISTS `treatment_use_medicals`;
CREATE TABLE `treatment_use_medicals`  (
  `id` int(32) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '记录id',
  `treatment_id` int(32) UNSIGNED ZEROFILL NULL DEFAULT NULL COMMENT '治疗方案id',
  `medical_id` int(32) UNSIGNED ZEROFILL NULL DEFAULT NULL COMMENT '药品使用id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `treatment_id`(`treatment_id`) USING BTREE,
  INDEX `medical_id`(`medical_id`) USING BTREE,
  CONSTRAINT `medical_id` FOREIGN KEY (`medical_id`) REFERENCES `medical_supplies` (`medical_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `treatment_id` FOREIGN KEY (`treatment_id`) REFERENCES `treatment_plans` (`treatment_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
