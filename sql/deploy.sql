CREATE DATABASE IF NOT EXISTS `taurus` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
DROP TABLE IF EXISTS `award`, `lottery`, `sms`, `signature`;

CREATE TABLE `signature`
(
    `id`          BIGINT AUTO_INCREMENT Primary Key,
    `phone`       VARCHAR(36)  NOT NULL COMMENT '用户手机号',
    `street`      VARCHAR(255) NOT NULL NULL COMMENT '街道',
    `create_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted`     TINYINT(1)            DEFAULT 0 COMMENT '是否被删除',
    `delete_time` DATETIME              DEFAULT NULL COMMENT '删除时间',

    KEY `idx_phone` (phone) using BTREE,
    KEY `idx_street` (street) using BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `lottery`
(
    `id`               BIGINT AUTO_INCREMENT PRIMARY KEY,
    `phone`            VARCHAR(36)  NOT NULL COMMENT '用户手机号',
    `is_win_lottery`   TINYINT(1)   NOT NULL DEFAULT 0 COMMENT '抽奖结果',
    `award_type`       INT          NOT NULL COMMENT '中奖类型',
    `award_code`       VARCHAR(128) NOT NULL COMMENT '中奖序列码',

    `create_time`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted`          TINYINT(1)            DEFAULT 0 COMMENT '是否被删除',
    `delete_time`      DATETIME              DEFAULT NULL COMMENT '删除时间',

    KEY `idx_phone` (phone) using BTREE,
    KEY `idx_time_type` (create_time, award_type)  using BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE `award`
(
    `id`          BIGINT AUTO_INCREMENT PRIMARY KEY,
    `code`        VARCHAR(128) NOT NULL COMMENT '奖品序列码',
    `type`        int          NOT NULL COMMENT '奖品类型',
    `is_used`     TINYINT(1)   NOT NULL DEFAULT 0 COMMENT '是否被使用',

    `version`     INT          NOT NULL DEFAULT 0 COMMENT '版本号',
    `create_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted`     TINYINT(1)            DEFAULT 0 COMMENT '是否被删除',
    `delete_time` DATETIME              DEFAULT NULL COMMENT '删除时间',

    UNIQUE KEY `idx_code` (code) using BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `sms`
(
    `id`          BIGINT AUTO_INCREMENT Primary Key,
    `phone`       VARCHAR(36)  NOT NULL COMMENT '用户手机号',
    `award_type`  INT          NOT NULL COMMENT '奖品类型',
    `award_code`  VARCHAR(128) NOT NULL COMMENT '奖品编码',
    `serial_no`   VARCHAR(255) NOT NULL NULL COMMENT '短信发送序列号',
    `is_sms_sent` VARCHAR(36)  NOT NULL NULL COMMENT '是否发送短信成功',
    `create_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted`     TINYINT(1)            DEFAULT 0 COMMENT '是否被删除',
    `delete_time` DATETIME              DEFAULT NULL COMMENT '删除时间',
    UNIQUE KEY `idx_phone` (phone) using BTREE,
    UNIQUE KEY `idx_sms_serial_no` (serial_no) using BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;