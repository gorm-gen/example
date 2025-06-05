CREATE TABLE IF NOT EXISTS `identity_card` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录ID',
    `number` char(18) NOT NULL COMMENT '身份证号码',
    `user_id` int UNSIGNED NOT NULL COMMENT '用户ID',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间戳[0:未删除,非0:删除时间戳]',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_number_deleted_at` (`number`,`deleted_at`),
    UNIQUE KEY `uk_user_id_deleted_at` (`user_id`,`deleted_at`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='身份证';