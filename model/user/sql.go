package user

import (
	"jun/dao"
	"jun/dto"
)

func selectAllByUid(p *dto.User, uid int64) error {
	return dao.Query1(p, "select * from user where u_id = ?", uid)
}

func selectAll(pp *[]dto.User) error {
	return dao.QueryN(pp, "select * from user")
}

func countAll(cnt *uint64) error {
	return dao.Query1(cnt, "select count(u_id) from user")
}

func auth(b *bool, uname *string, pwd *string) error {
	return dao.Query1(b, "select count(u_id) from user " +
		"where u_uname = ? and u_pwd_encrypted = sha1(concat(?, 'jun990527'))", uname, pwd)
}

func add(p *dto.User) (int64, error) {
	return dao.Insert("insert into user values(?)", p)
}