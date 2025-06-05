CREATE TABLE IF NOT EXISTS `user_language` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录ID',
    `user_id` int UNSIGNED NOT NULL COMMENT '用户ID',
    `language_id` int UNSIGNED NOT NULL COMMENT '语言ID',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间戳[0:未删除,非0:删除时间戳]',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_language_id` (`language_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户语言表';