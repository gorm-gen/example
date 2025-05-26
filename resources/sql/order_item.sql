CREATE TABLE IF NOT EXISTS `order_item` (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录ID',
    `sharding` char(6) NOT NULL COMMENT '分表关键字段[月份:202503]',
    `uid` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户ID',
    `order_no` char(24) NOT NULL DEFAULT '' COMMENT '订单号',
    `comment` varchar(255) NOT NULL DEFAULT '' COMMENT '订单评论',
    `express_number` varchar(50) NOT NULL DEFAULT '' COMMENT '快递单号',
    `created_at` datetime NOT NULL COMMENT '创建日期',
    `updated_at` datetime NOT NULL COMMENT '更新日期',
    `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间[null:未删除,非null:删除时间]',
    PRIMARY KEY (`id`),
    KEY `idx_uid` (`uid`),
    UNIQUE KEY `uk_order_no` (`order_no`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单条目表';