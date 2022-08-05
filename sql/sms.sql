CREATE TABLE `sms`
(
    `id`              BIGINT AUTO_INCREMENT Primary Key,
    `phone`           VARCHAR(36)  NOT NULL COMMENT '用户手机号',
    `sms_serial_no`   VARCHAR(255) NOT NULL NULL COMMENT '短信发送序列号',
    `sms_send_status` VARCHAR(36) NOT NULL NULL COMMENT '短信发送状态',
    `create_time`     DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`     DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted`         TINYINT(1)            DEFAULT 0 COMMENT '是否被删除',
    `delete_time`     DATETIME              DEFAULT NULL COMMENT '删除时间',
    UNIQUE KEY `idx_phone` (phone) using BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;