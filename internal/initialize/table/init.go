package table

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"

	"example/internal/global"
)

var once sync.Once

func Init() {
	once.Do(func() {
		tableName := fmt.Sprintf("order_%s", time.Now().Format("200601"))
		if !global.DB.Migrator().HasTable(tableName) {
			bytes, err := os.ReadFile("resources/sql/order.sql")
			if err != nil {
				global.Logger.Error("读取【order.sql】失败", zap.Error(err))
				log.Fatal(err)
				return
			}
			sql := string(bytes)
			sql = strings.ReplaceAll(sql, "CREATE TABLE IF NOT EXISTS `order` (", fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (", tableName))
			if err = global.DB.Exec(sql).Error; err != nil {
				global.Logger.Error(fmt.Sprintf("创建表【%s】失败", tableName), zap.Error(err))
				log.Fatal(err)
				return
			}
		}
	})
}
