CREATE TABLE `lottery`
(
    `id`               BIGINT AUTO_INCREMENT PRIMARY KEY,
    `phone`            VARCHAR(36)  NOT NULL COMMENT '用户手机号',
    `win_lottery_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '抽奖时间',
    `is_win_lottery`   TINYINT(1)   NOT NULL DEFAULT 0 COMMENT '抽奖结果',
    `award_type`       INT          NOT NULL COMMENT '中奖类型',
    `award_code`       VARCHAR(128) NOT NULL COMMENT '中奖序列码',

    `create_time`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted`          TINYINT(1)            DEFAULT 0 COMMENT '是否被删除',
    `delete_time`      DATETIME              DEFAULT NULL COMMENT '删除时间',

    KEY `idx_phone` (phone) using BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
