package tag

import (
	"jun/dao"
)

func existTagIDs(ids []int64) (bool, error) {
	var b bool
	var idss []interface{}
	sql := `select true `
	for _, id := range ids {
		sql += ` and (select 1 from tag where t_id = ?`
		idss = append(idss, id)
	}

	err := dao.Query1(&b, sql, idss...)
	return b, err
}

func existTagNames(names []string) (bool, error) {
	var b bool
	var namess []interface{}
	sql := `select true `
	for _, name := range names {
		sql += ` and (select 1 from tag where t_name = ?`
		namess = append(namess, name)
	}

	err := dao.Query1(&b, sql, namess...)
	return b, err
}
