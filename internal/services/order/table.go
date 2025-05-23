package order

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
)

func (o *Order) Table(shardingKey string, path ...string) error {
	tableName := fmt.Sprintf("order_%s", shardingKey)
	_path := "resources/sql/order.sql"
	if len(path) > 0 && path[0] != "" {
		_path = path[0]
	}
	if !o.db.Migrator().HasTable(tableName) {
		bytes, err := os.ReadFile(_path)
		if err != nil {
			o.logger.Error("读取【order.sql】失败", zap.Error(err))
			return err
		}
		sql := string(bytes)
		sql = strings.ReplaceAll(sql, "CREATE TABLE IF NOT EXISTS `order` (", fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (", tableName))
		if err = o.db.Exec(sql).Error; err != nil {
			o.logger.Error(fmt.Sprintf("创建表【%s】失败", tableName), zap.Error(err))
			return err
		}
	}
	return nil
}
