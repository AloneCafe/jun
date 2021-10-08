package option

import (
	"jun/dao"
)

func getOption(key string) (*string, error) {
	val := new(string)
	sql := `select o_val from option where o_key = ?`
	err := dao.Query1(val, sql, key)
	return val, err
}
