package category

import (
	"jun/dao"
)

func existCategoryIDs(ids []int64) (bool, error) {
	var b bool
	var idss []interface{}
	sql := `select true `
	for _, id := range ids {
		sql += ` and (select 1 from category where cg_id = ?`
		idss = append(idss, id)
	}

	err := dao.Query1(&b, sql, idss...)
	return b, err
}

func existCategoryNames(names []string) (bool, error) {
	var b bool
	var namess []interface{}
	sql := `select true `
	for _, name := range names {
		sql += ` and (select 1 from category where cg_name = ?`
		namess = append(namess, name)
	}

	err := dao.Query1(&b, sql, namess...)
	return b, err
}
