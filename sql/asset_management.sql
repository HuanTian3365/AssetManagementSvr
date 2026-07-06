/*
 Navicat MySQL Dump SQL

 Source Server         : 本机
 Source Server Type    : MySQL
 Source Server Version : 90701 (9.7.1)
 Source Host           : localhost:3306
 Source Schema         : asset_management

 Target Server Type    : MySQL
 Target Server Version : 90701 (9.7.1)
 File Encoding         : 65001

 Date: 06/07/2026 23:06:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for asset_building
-- ----------------------------
DROP TABLE IF EXISTS `asset_building`;
CREATE TABLE `asset_building`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '建筑名称',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '建筑编码',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_asset_building_code`(`code` ASC) USING BTREE,
  INDEX `idx_asset_building_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '建筑表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of asset_building
-- ----------------------------
INSERT INTO `asset_building` VALUES (1, '新大楼', 'sz-xdl', '山东商业职业技术大学新大楼', '新大楼', '2026-07-06 14:09:00', '2026-07-06 14:09:00', NULL);
INSERT INTO `asset_building` VALUES (2, '教学楼A楼', 'sz-jxa', '山东商业职业技术大学教学楼A楼', '教学楼A楼', '2026-07-06 14:09:24', '2026-07-06 14:09:24', NULL);
INSERT INTO `asset_building` VALUES (3, '教学楼B楼', 'sz-jxb', '山东商业职业技术大学教学楼B楼', '教学楼B楼', '2026-07-06 14:09:32', '2026-07-06 14:09:32', NULL);
INSERT INTO `asset_building` VALUES (4, '教学楼C楼', 'sz-jxc', '山东商业职业技术大学教学楼C楼', '教学楼C楼', '2026-07-06 14:09:43', '2026-07-06 14:09:43', NULL);

-- ----------------------------
-- Table structure for asset_category
-- ----------------------------
DROP TABLE IF EXISTS `asset_category`;
CREATE TABLE `asset_category`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `parent_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '父分类ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类名称',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类编码',
  `sort` int NOT NULL DEFAULT 0 COMMENT '排序',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_asset_category_code`(`code` ASC) USING BTREE,
  INDEX `idx_asset_category_parent_id`(`parent_id` ASC) USING BTREE,
  INDEX `idx_asset_category_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of asset_category
-- ----------------------------

-- ----------------------------
-- Table structure for asset_floor
-- ----------------------------
DROP TABLE IF EXISTS `asset_floor`;
CREATE TABLE `asset_floor`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `building_id` bigint UNSIGNED NOT NULL COMMENT '建筑ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '楼层名称',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '楼层编码',
  `floor_no` int NOT NULL COMMENT '楼层序号',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_asset_floor_code`(`code` ASC) USING BTREE,
  INDEX `idx_asset_floor_building_id`(`building_id` ASC) USING BTREE,
  INDEX `idx_asset_floor_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '楼层表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of asset_floor
-- ----------------------------
INSERT INTO `asset_floor` VALUES (1, 1, '大厅', 'xdl-01', 1, '一楼大厅', '2026-07-06 14:10:42', '2026-07-06 14:10:42', NULL);
INSERT INTO `asset_floor` VALUES (2, 1, '二楼', 'xdl-02', 2, '二楼', '2026-07-06 14:11:02', '2026-07-06 14:11:02', NULL);
INSERT INTO `asset_floor` VALUES (3, 1, '三楼', 'xdl-03', 3, '三楼', '2026-07-06 14:11:17', '2026-07-06 14:11:17', NULL);
INSERT INTO `asset_floor` VALUES (4, 1, '四楼', 'xdl-04', 4, '四楼', '2026-07-06 14:11:28', '2026-07-06 14:11:28', NULL);
INSERT INTO `asset_floor` VALUES (5, 1, '五楼', 'xdl-05', 5, '五楼', '2026-07-06 14:11:37', '2026-07-06 14:11:37', NULL);
INSERT INTO `asset_floor` VALUES (6, 1, '六楼', 'xdl-06', 6, '六楼', '2026-07-06 14:11:50', '2026-07-06 14:11:50', NULL);
INSERT INTO `asset_floor` VALUES (7, 1, '七楼', 'xdl-07', 7, '七楼', '2026-07-06 14:11:59', '2026-07-06 14:11:59', NULL);
INSERT INTO `asset_floor` VALUES (8, 1, '八楼', 'xdl-08', 8, '八楼', '2026-07-06 14:12:12', '2026-07-06 14:12:12', NULL);

-- ----------------------------
-- Table structure for asset_item
-- ----------------------------
DROP TABLE IF EXISTS `asset_item`;
CREATE TABLE `asset_item`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `product_id` bigint UNSIGNED NOT NULL COMMENT '产品ID',
  `asset_code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '资产编码',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '资产名称',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '资产状态：1闲置 2在用 3维修 4报废',
  `current_building_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '建筑ID',
  `current_floor_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '楼层ID',
  `current_room_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '房间ID',
  `purchase_date` date NULL DEFAULT NULL COMMENT '购买日期',
  `purchase_price` decimal(12, 2) NOT NULL DEFAULT 0.00 COMMENT '购买价格',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_asset_item_asset_code`(`asset_code` ASC) USING BTREE,
  INDEX `idx_asset_item_product_id`(`product_id` ASC) USING BTREE,
  INDEX `idx_asset_item_current_room_id`(`current_room_id` ASC) USING BTREE,
  INDEX `idx_asset_item_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '资产表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of asset_item
-- ----------------------------

-- ----------------------------
-- Table structure for asset_location_record
-- ----------------------------
DROP TABLE IF EXISTS `asset_location_record`;
CREATE TABLE `asset_location_record`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `asset_id` bigint UNSIGNED NOT NULL COMMENT '资产ID',
  `action_type` tinyint NOT NULL COMMENT '动作类型: 0入库 1转移',
  `from_building_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '来源建筑ID',
  `from_floor_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '来源楼层ID',
  `from_room_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '来源房间ID',
  `to_building_id` bigint UNSIGNED NOT NULL COMMENT '目标建筑ID',
  `to_floor_id` bigint UNSIGNED NOT NULL COMMENT '目标楼层ID',
  `to_room_id` bigint UNSIGNED NOT NULL COMMENT '目标房间ID',
  `operated_at` datetime NOT NULL COMMENT '操作时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_asset_location_record_asset_id`(`asset_id` ASC) USING BTREE,
  INDEX `idx_asset_location_record_action_type`(`action_type` ASC) USING BTREE,
  INDEX `idx_asset_location_record_to_room_id`(`to_room_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '资产位置记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of asset_location_record
-- ----------------------------

-- ----------------------------
-- Table structure for asset_product
-- ----------------------------
DROP TABLE IF EXISTS `asset_product`;
CREATE TABLE `asset_product`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `category_id` bigint UNSIGNED NOT NULL COMMENT '分类ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '产品名称',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '产品编码',
  `brand` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '品牌',
  `model` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '型号',
  `unit` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '件' COMMENT '计量单位',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_asset_product_code`(`code` ASC) USING BTREE,
  INDEX `idx_asset_product_category_id`(`category_id` ASC) USING BTREE,
  INDEX `idx_asset_product_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of asset_product
-- ----------------------------

-- ----------------------------
-- Table structure for asset_room
-- ----------------------------
DROP TABLE IF EXISTS `asset_room`;
CREATE TABLE `asset_room`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `building_id` bigint UNSIGNED NOT NULL COMMENT '建筑ID',
  `floor_id` bigint UNSIGNED NOT NULL COMMENT '楼层ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '房间名称',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '房间编码',
  `room_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '房间号',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_asset_room_code`(`code` ASC) USING BTREE,
  INDEX `idx_asset_room_building_id`(`building_id` ASC) USING BTREE,
  INDEX `idx_asset_room_floor_id`(`floor_id` ASC) USING BTREE,
  INDEX `idx_asset_room_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '房间表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of asset_room
-- ----------------------------
INSERT INTO `asset_room` VALUES (1, 1, 8, '801', 'xdl-08-801', '801', '办公室', '2026-07-06 15:05:18', '2026-07-06 15:05:18', NULL);

SET FOREIGN_KEY_CHECKS = 1;
