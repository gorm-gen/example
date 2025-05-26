CREATE TABLE IF NOT EXISTS `order` (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录ID',
    `sharding` char(6) NOT NULL COMMENT '分表关键字段[月份:202503]',
    `uid` int UNSIGNED NOT NULL COMMENT '用户ID',
    `order_no` char(24) NOT NULL COMMENT '订单号',
    `status` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '订单状态',
    `amount` decimal(25, 2) UNSIGNED NOT NULL DEFAULT '0.00' COMMENT '订单金额[单位:元]',
    `created_at` datetime NOT NULL COMMENT '创建日期',
    `updated_at` datetime NOT NULL COMMENT '更新日期',
    `deleted_at` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间戳[0:未删除,非0:删除时间戳]',
    PRIMARY KEY (`id`),
    KEY `idx_uid` (`uid`),
    UNIQUE KEY `uk_order_no` (`order_no`),
    KEY `idx_status` (`status`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';