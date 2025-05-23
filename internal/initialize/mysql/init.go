package mysql

import (
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/gorm-gen/plugin/logger"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"example/internal/global"
)

var once sync.Once

func Init(useSharding ...bool) {
	_useSharding := true
	if len(useSharding) > 0 {
		_useSharding = useSharding[0]
	}
	once.Do(func() {
		newMySql(_useSharding)
	})
}

func newMySql(useSharding bool) {
	mysqlConfig := global.Config.Mysql
	dsn := mysqlConfig.DSN()

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger: logger.New(
			logger.WithPath(global.Config.Log.Path),
			logger.WithMaxBackups(100),
			logger.WithMaxAge(7),
		).Logger(),
	})
	if err != nil {
		global.Logger.Error("MySQL init failed -1.", zap.Error(err))
		log.Fatal(err)
		return
	}

	var sqlDB *sql.DB
	if sqlDB, err = db.DB(); err != nil {
		global.Logger.Error("MySQL init failed -3.", zap.Error(err))
		log.Fatal(err)
		return
	}

	if err = sqlDB.Ping(); err != nil {
		global.Logger.Error("MySQL init failed -4.", zap.Error(err))
		log.Fatal(err)
		return
	}

	sqlDB.SetConnMaxIdleTime(time.Hour)
	sqlDB.SetConnMaxLifetime(6 * time.Hour)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(500)

	if mysqlConfig.Debug {
		global.DB = db.Debug()
		return
	}

	global.DB = db
}
