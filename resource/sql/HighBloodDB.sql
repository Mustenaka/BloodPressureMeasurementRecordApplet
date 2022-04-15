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

 Date: 15/04/2022 18:54:57
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
  `real_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '真实姓名',
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
  `heart_rate` int(32) NULL DEFAULT NULL COMMENT '测量患者心率',
  PRIMARY KEY (`record_id`) USING BTREE,
  INDEX `user_recordbp_id`(`user_id`) USING BTREE,
  CONSTRAINT `user_recordbp_id` FOREIGN KEY (`user_id`) REFERENCES `base_users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1040 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for patient_infos
-- ----------------------------
DROP TABLE IF EXISTS `patient_infos`;
CREATE TABLE `patient_infos`  (
  `patient_id` int(32) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '病历信息id',
  `user_id` int(32) UNSIGNED ZEROFILL NOT NULL COMMENT '对应用户',
  `real_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '真实姓名',
  `tel` char(11) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '用户电话号码',
  `birthday` date NULL DEFAULT NULL COMMENT '生日',
  `sex` enum('男','女','其他') CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '性别',
  `is_married` tinyint(1) NULL DEFAULT NULL COMMENT '0-未婚、1-已婚',
  `hbp_years` int(32) NULL DEFAULT NULL COMMENT '高血压患病时间（年）',
  `anamnesis` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '既往病史(对应表格1~12)',
  `is_smoking` tinyint(1) NULL DEFAULT NULL COMMENT '是否吸烟',
  `smoking_history` int(32) NULL DEFAULT NULL COMMENT '吸烟史（年）',
  `smoking_daily` int(32) NULL DEFAULT NULL COMMENT '日吸烟数',
  `is_drink` tinyint(1) NULL DEFAULT NULL COMMENT '是否饮酒',
  `drink_history` int(32) NULL DEFAULT NULL COMMENT '饮酒史（年）',
  `drink_daily` int(32) NULL DEFAULT NULL COMMENT '每日饮酒量',
  `patient_height` int(32) NULL DEFAULT NULL COMMENT '身高',
  `patient_weight` int(32) NULL DEFAULT NULL COMMENT '体重',
  `patient_waist_circumference` int(32) NULL DEFAULT NULL COMMENT '腰围',
  `patient_chest_circumference` int(32) NULL DEFAULT NULL COMMENT '胸围',
  `patient_hip_circumference` int(32) NULL DEFAULT NULL COMMENT '臀围',
  `is_take_chinese_medicine` tinyint(1) NULL DEFAULT NULL COMMENT '是否服用中药',
  `antihypertensive_plan` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '降压方案',
  `is_nondrug_control_plan` tinyint(1) NULL DEFAULT NULL COMMENT '是否非药物控制手段',
  `nondrug_control_plan` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '非药物控制手段',
  PRIMARY KEY (`patient_id`) USING BTREE,
  INDEX `user_info_id`(`user_id`) USING BTREE,
  CONSTRAINT `user_info_id` FOREIGN KEY (`user_id`) REFERENCES `base_users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for treatment_plans
-- ----------------------------
DROP TABLE IF EXISTS `treatment_plans`;
CREATE TABLE `treatment_plans`  (
  `treatment_id` int(32) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '治疗方案id',
  `user_id` int(32) UNSIGNED ZEROFILL NOT NULL COMMENT '患者id(即用户id)',
  `plan` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '治疗方案内容',
  `note` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '治疗方案备注，如禁忌症',
  `create_datetime` datetime(0) NULL DEFAULT NULL COMMENT '方案创建日期时间',
  `status` enum('生效','失效') CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '生效' COMMENT '该方案状态（“生效”，“失效”）',
  PRIMARY KEY (`treatment_id`) USING BTREE,
  INDEX `user_treatment_id`(`user_id`) USING BTREE,
  CONSTRAINT `user_treatment_id` FOREIGN KEY (`user_id`) REFERENCES `base_users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
