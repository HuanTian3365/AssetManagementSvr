# Task 1 / Step 1: 创建资产模块建表 SQL

> 来源: `../2026-06-24-personal-asset-management-api-plan.md`  
> 上级任务: [建表 SQL](README.md)

- [ ] **Step 1: 创建资产模块建表 SQL**

写入 `manifest/sql/asset_init.sql`：

```sql
CREATE TABLE IF NOT EXISTS asset_building (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  name VARCHAR(100) NOT NULL COMMENT '建筑名称',
  code VARCHAR(64) NOT NULL COMMENT '建筑编码',
  address VARCHAR(255) NOT NULL DEFAULT '' COMMENT '地址',
  remark VARCHAR(500) NOT NULL DEFAULT '' COMMENT '备注',
  created_at DATETIME NULL COMMENT '创建时间',
  updated_at DATETIME NULL COMMENT '更新时间',
  deleted_at DATETIME NULL COMMENT '删除时间',
  PRIMARY KEY (id),
  UNIQUE KEY uk_asset_building_code (code),
  KEY idx_asset_building_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资产建筑';

CREATE TABLE IF NOT EXISTS asset_floor (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  building_id BIGINT UNSIGNED NOT NULL COMMENT '建筑ID',
  name VARCHAR(100) NOT NULL COMMENT '楼层名称',
  code VARCHAR(64) NOT NULL COMMENT '楼层编码',
  floor_no INT NOT NULL DEFAULT 0 COMMENT '楼层序号',
  remark VARCHAR(500) NOT NULL DEFAULT '' COMMENT '备注',
  created_at DATETIME NULL COMMENT '创建时间',
  updated_at DATETIME NULL COMMENT '更新时间',
  deleted_at DATETIME NULL COMMENT '删除时间',
  PRIMARY KEY (id),
  UNIQUE KEY uk_asset_floor_code (code),
  KEY idx_asset_floor_building_id (building_id),
  KEY idx_asset_floor_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资产楼层';

CREATE TABLE IF NOT EXISTS asset_room (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  building_id BIGINT UNSIGNED NOT NULL COMMENT '建筑ID',
  floor_id BIGINT UNSIGNED NOT NULL COMMENT '楼层ID',
  name VARCHAR(100) NOT NULL COMMENT '房间名称',
  code VARCHAR(64) NOT NULL COMMENT '房间编码',
  room_no VARCHAR(64) NOT NULL DEFAULT '' COMMENT '房间号',
  remark VARCHAR(500) NOT NULL DEFAULT '' COMMENT '备注',
  created_at DATETIME NULL COMMENT '创建时间',
  updated_at DATETIME NULL COMMENT '更新时间',
  deleted_at DATETIME NULL COMMENT '删除时间',
  PRIMARY KEY (id),
  UNIQUE KEY uk_asset_room_code (code),
  KEY idx_asset_room_building_id (building_id),
  KEY idx_asset_room_floor_id (floor_id),
  KEY idx_asset_room_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资产房间';

CREATE TABLE IF NOT EXISTS asset_category (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  parent_id BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级ID',
  name VARCHAR(100) NOT NULL COMMENT '分类名称',
  code VARCHAR(64) NOT NULL COMMENT '分类编码',
  sort INT NOT NULL DEFAULT 0 COMMENT '排序',
  remark VARCHAR(500) NOT NULL DEFAULT '' COMMENT '备注',
  created_at DATETIME NULL COMMENT '创建时间',
  updated_at DATETIME NULL COMMENT '更新时间',
  deleted_at DATETIME NULL COMMENT '删除时间',
  PRIMARY KEY (id),
  UNIQUE KEY uk_asset_category_code (code),
  KEY idx_asset_category_parent_id (parent_id),
  KEY idx_asset_category_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资产分类';

CREATE TABLE IF NOT EXISTS asset_product (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  category_id BIGINT UNSIGNED NOT NULL COMMENT '分类ID',
  name VARCHAR(100) NOT NULL COMMENT '产品名称',
  code VARCHAR(64) NOT NULL COMMENT '产品编码',
  brand VARCHAR(100) NOT NULL DEFAULT '' COMMENT '品牌',
  model VARCHAR(100) NOT NULL DEFAULT '' COMMENT '型号',
  unit VARCHAR(32) NOT NULL DEFAULT '件' COMMENT '单位',
  remark VARCHAR(500) NOT NULL DEFAULT '' COMMENT '备注',
  created_at DATETIME NULL COMMENT '创建时间',
  updated_at DATETIME NULL COMMENT '更新时间',
  deleted_at DATETIME NULL COMMENT '删除时间',
  PRIMARY KEY (id),
  UNIQUE KEY uk_asset_product_code (code),
  KEY idx_asset_product_category_id (category_id),
  KEY idx_asset_product_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资产产品';

CREATE TABLE IF NOT EXISTS asset_item (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  product_id BIGINT UNSIGNED NOT NULL COMMENT '产品ID',
  asset_code VARCHAR(128) NOT NULL COMMENT '资产编码',
  name VARCHAR(100) NOT NULL COMMENT '资产名称',
  status TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1闲置 2在用 3维修 4报废',
  current_building_id BIGINT UNSIGNED NULL COMMENT '当前建筑ID',
  current_floor_id BIGINT UNSIGNED NULL COMMENT '当前楼层ID',
  current_room_id BIGINT UNSIGNED NULL COMMENT '当前房间ID',
  purchase_date DATE NULL COMMENT '购买日期',
  purchase_price DECIMAL(12,2) NOT NULL DEFAULT 0 COMMENT '购买价格',
  remark VARCHAR(500) NOT NULL DEFAULT '' COMMENT '备注',
  created_at DATETIME NULL COMMENT '创建时间',
  updated_at DATETIME NULL COMMENT '更新时间',
  deleted_at DATETIME NULL COMMENT '删除时间',
  PRIMARY KEY (id),
  UNIQUE KEY uk_asset_item_asset_code (asset_code),
  KEY idx_asset_item_product_id (product_id),
  KEY idx_asset_item_current_room_id (current_room_id),
  KEY idx_asset_item_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资产实例';

CREATE TABLE IF NOT EXISTS asset_location_record (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  asset_id BIGINT UNSIGNED NOT NULL COMMENT '资产ID',
  action_type VARCHAR(32) NOT NULL COMMENT '动作类型：inbound入库 transfer转移',
  from_building_id BIGINT UNSIGNED NULL COMMENT '来源建筑ID',
  from_floor_id BIGINT UNSIGNED NULL COMMENT '来源楼层ID',
  from_room_id BIGINT UNSIGNED NULL COMMENT '来源房间ID',
  to_building_id BIGINT UNSIGNED NOT NULL COMMENT '目标建筑ID',
  to_floor_id BIGINT UNSIGNED NOT NULL COMMENT '目标楼层ID',
  to_room_id BIGINT UNSIGNED NOT NULL COMMENT '目标房间ID',
  operated_at DATETIME NOT NULL COMMENT '操作时间',
  remark VARCHAR(500) NOT NULL DEFAULT '' COMMENT '备注',
  created_at DATETIME NULL COMMENT '创建时间',
  PRIMARY KEY (id),
  KEY idx_asset_location_record_asset_id (asset_id),
  KEY idx_asset_location_record_action_type (action_type),
  KEY idx_asset_location_record_to_room_id (to_room_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资产位置记录';
```
