# asset_management
# utf8mb4
# utf8mb4_0900_ai_ci

drop table if exists `asset_building`;
create table if not exists `asset_building`
(
    `id`         bigint unsigned not null auto_increment comment 'ID',
    `name`       varchar(100)    not null comment '建筑名称',
    `code`       varchar(64)     not null comment '建筑编码',
    `address`    varchar(255)    not null default '' comment '地址',
    `remark`     varchar(500)    not null default '' comment '备注',
    `created_at` datetime        not null default current_timestamp comment '创建时间',
    `updated_at` datetime        not null default current_timestamp on update current_timestamp comment '更新时间',
    `deleted_at` datetime comment '删除时间',
    primary key (`id`),
    unique key `uk_asset_building_code` (`code`),
    key `idx_asset_building_deleted_at` (`deleted_at`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci
    comment ='建筑表';


drop table if exists `asset_floor`;
create table if not exists `asset_floor`
(
    `id`          bigint unsigned not null auto_increment comment 'ID',
    `building_id` bigint unsigned not null comment '建筑ID',
    `name`        varchar(100)    not null comment '楼层名称',
    `code`        varchar(64)     not null comment '楼层编码',
    `floor_no`    int             not null comment '楼层序号',
    `remark`      varchar(500)    not null default '' comment '备注',
    `created_at`  datetime        not null default current_timestamp comment '创建时间',
    `updated_at`  datetime        not null default current_timestamp on update current_timestamp comment '更新时间',
    `deleted_at`  datetime comment '删除时间',
    primary key (`id`),
    unique key `uk_asset_floor_code` (`code`),
    key `idx_asset_floor_building_id` (`building_id`),
    key `idx_asset_floor_deleted_at` (`deleted_at`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci
    comment ='楼层表';



drop table if exists `asset_room`;
create table if not exists `asset_room`
(
    `id`          bigint unsigned not null auto_increment comment 'ID',
    `building_id` bigint unsigned not null comment '建筑ID',
    `floor_id`    bigint unsigned not null comment '楼层ID',
    `name`        varchar(100)    not null comment '房间名称',
    `code`        varchar(64)     not null comment '房间编码',
    `room_no`     varchar(64)     not null default '' comment '房间号',
    `remark`      varchar(500)    not null default '' comment '备注',
    `created_at`  datetime        not null default current_timestamp comment '创建时间',
    `updated_at`  datetime        not null default current_timestamp on update current_timestamp comment '更新时间',
    `deleted_at`  datetime comment '删除时间',
    primary key (`id`),
    unique key `uk_asset_room_code` (`code`),
    key `idx_asset_room_building_id` (`building_id`),
    key `idx_asset_room_floor_id` (`floor_id`),
    key `idx_asset_room_deleted_at` (`deleted_at`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci
    comment ='房间表';


drop table if exists `asset_category`;
create table if not exists `asset_category`
(
    `id`         bigint unsigned not null auto_increment comment 'ID',
    `parent_id`  bigint unsigned not null default 0 comment '父分类ID',
    `name`       varchar(100)    not null comment '分类名称',
    `code`       varchar(64)     not null comment '分类编码',
    `sort`       int             not null default 0 comment '排序',
    `remark`     varchar(500)    not null default '' comment '备注',
    `created_at` datetime        not null default current_timestamp comment '创建时间',
    `updated_at` datetime        not null default current_timestamp on update current_timestamp comment '更新时间',
    `deleted_at` datetime comment '删除时间',
    primary key (`id`),
    unique key `uk_asset_category_code` (`code`),
    key `idx_asset_category_parent_id` (`parent_id`),
    key `idx_asset_category_deleted_at` (`deleted_at`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci
    comment ='分类表';


drop table if exists `asset_product`;
create table if not exists `asset_product`
(
    `id`          bigint unsigned not null auto_increment comment 'ID',
    `category_id` bigint unsigned not null comment '分类ID',
    `name`        varchar(100)    not null comment '产品名称',
    `code`        varchar(64)     not null comment '产品编码',
    `brand`       varchar(100)    not null default '' comment '品牌',
    `model`       varchar(100)    not null default '' comment '型号',
    `unit`        varchar(32)     not null default '件' comment '计量单位',
    `remark`      varchar(500)    not null default '' comment '备注',
    `created_at`  datetime        not null default current_timestamp comment '创建时间',
    `updated_at`  datetime        not null default current_timestamp on update current_timestamp comment '更新时间',
    `deleted_at`  datetime comment '删除时间',
    primary key (`id`),
    unique key `uk_asset_product_code` (`code`),
    key `idx_asset_product_category_id` (`category_id`),
    key `idx_asset_product_deleted_at` (`deleted_at`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci
    comment ='产品表';

drop table if exists `asset_item`;
create table if not exists `asset_item`
(
    `id`                  bigint unsigned not null auto_increment comment 'ID',
    `product_id`          bigint unsigned not null comment '产品ID',
    `asset_code`          varchar(128)    not null comment '资产编码',
    `name`                varchar(100)    not null comment '资产名称',
    `status`              tinyint         not null default 1 comment '资产状态：1闲置 2在用 3维修 4报废',
    `current_building_id` bigint unsigned comment '建筑ID',
    `current_floor_id`    bigint unsigned comment '楼层ID',
    `current_room_id`     bigint unsigned comment '房间ID',
    `purchase_date`       date            null comment '购买日期',
    `purchase_price`      decimal(12, 2)  not null default 0 comment '购买价格',
    `remark`              varchar(500)    not null default '' comment '备注',
    `created_at`          datetime        not null default current_timestamp comment '创建时间',
    `updated_at`          datetime        not null default current_timestamp on update current_timestamp comment '更新时间',
    `deleted_at`          datetime comment '删除时间',
    primary key (`id`),
    unique key `uk_asset_item_asset_code` (`asset_code`),
    key `idx_asset_item_product_id` (`product_id`),
    key `idx_asset_item_current_room_id` (`current_room_id`),
    key `idx_asset_item_deleted_at` (`deleted_at`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci
    comment ='资产表';


drop table if exists `asset_location_record`;
create table if not exists `asset_location_record`
(
    `id`               bigint unsigned not null auto_increment comment 'ID',
    `asset_id`         bigint unsigned not null comment '资产ID',
    `action_type`      tinyint         not null comment '动作类型: 0入库 1转移',
    `from_building_id` bigint unsigned null comment '来源建筑ID',
    `from_floor_id`    bigint unsigned null comment '来源楼层ID',
    `from_room_id`     bigint unsigned null comment '来源房间ID',
    `to_building_id`   bigint unsigned not null comment '目标建筑ID',
    `to_floor_id`      bigint unsigned not null comment '目标楼层ID',
    `to_room_id`       bigint unsigned not null comment '目标房间ID',
    `operated_at`      datetime        not null comment '操作时间',
    `remark`           varchar(500)    not null default '' comment '备注',
    `created_at`       datetime        not null default current_timestamp comment '创建时间',
    primary key (`id`),
    key `idx_asset_location_record_asset_id` (`asset_id`),
    key `idx_asset_location_record_action_type` (`action_type`),
    key `idx_asset_location_record_to_room_id` (`to_room_id`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci
    comment ='资产位置记录';

