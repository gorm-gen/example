CREATE TABLE IF NOT EXISTS `area` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '地区ID',
    `name` varchar(50) NOT NULL DEFAULT '' COMMENT '地区名',
    `pid` int UNSIGNED NOT NULL COMMENT '父级ID',
    `level` tinyint UNSIGNED NOT NULL COMMENT '地区等级[1:省,2:市,3:县/区]',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `deleted_at` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间戳[0:未删除,非0:删除时间戳]',
    PRIMARY KEY (`id`),
    KEY `idx_pid` (`pid`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='地区表';