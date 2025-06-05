CREATE TABLE IF NOT EXISTS `language` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `name` varchar(50) NOT NULL DEFAULT '' COMMENT '语言名称',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间戳[0:未删除,非0:删除时间戳]',
    PRIMARY KEY (`id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='语言表';