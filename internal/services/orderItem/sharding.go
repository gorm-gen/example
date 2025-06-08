package orderItem

import (
	"regexp"
)

// ShardingSuffix 获取分表后缀
func (o *OrderItem) ShardingSuffix() ([]string, error) {
	reg := regexp.MustCompile(`^order_item_(\d{6})$`)
	tables, err := o.db.Migrator().GetTables()
	if err != nil {
		return nil, err
	}
	var list []string
	for _, table := range tables {
		arr := reg.FindStringSubmatch(table)
		if len(arr) != 2 {
			continue
		}
		list = append(list, arr[1])
	}
	return list, nil
}
