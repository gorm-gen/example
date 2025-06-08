package order_test

import (
	"log"
	"time"

	"example/internal/global"
	"example/internal/initialize/config"
	"example/internal/initialize/logger"
	"example/internal/initialize/mysql"
	"example/internal/services/order"
)

var orderSvc *order.Order

var shardingList []string

func init() {
	// 1、初始化本地配置文件
	config.InitAuto(global.ConfigFile)

	// 2、初始化日志
	logger.Init(false, logger.Hooks())

	// 3、初始化MySQL
	mysql.Init()

	// 4、实例化order服务
	orderSvc = order.New()

	var err error

	// 5、初始化数据库表
	if err = orderSvc.Table(time.Now().Format("200601"), "../../../resources/sql/order.sql"); err != nil {
		log.Fatal(err)
		return
	}

	// 6、获取分表后缀
	if shardingList, err = orderSvc.ShardingSuffix(); err != nil {
		log.Fatal(err)
		return
	}
}
