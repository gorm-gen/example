package main

import (
	"flag"

	"example/internal/global"
	"example/internal/initialize/config"
	"example/internal/initialize/logger"
	"example/internal/initialize/mysql"
	"example/internal/initialize/table"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", global.ConfigFile, "local config file path")

	// 1、初始化本地配置文件
	config.Init(configPath)

	// 2、初始化日志
	logger.Init(false, logger.Hooks())

	// 3、初始化MySQL
	mysql.Init()

	// 4、初始化数据库表
	table.Init()
}

func main() {
	flag.Parse()
}
