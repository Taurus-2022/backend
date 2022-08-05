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

    UNIQUE KEY `idx_code` (code) using BTREE,
    KEY `idx_type_is_used` (`type`, `is_used`) using BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;