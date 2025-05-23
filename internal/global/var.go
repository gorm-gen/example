package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Config 配置信息
var Config *Conf

// Logger 日志
var Logger *zap.Logger

// DB MySQL 数据库
var DB *gorm.DB
