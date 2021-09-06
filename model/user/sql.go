package user

import (
	"jun/dao"
	"jun/dto"
	"time"
)

func selectById(p *dto.User, id *int64) error {
	return dao.Query1(p, "select * from user where u_id = ? limit 1", id)
}

func selectByUname(p *dto.User, uname *string) error {
	return dao.Query1(p, "select * from user where u_uname = ? limit 1", uname)
}

func selectByEmail(p *dto.User, email *string) error {
	return dao.Query1(p, "select * from user where u_email = ? limit 1", email)
}

func selectAll(pp *[]dto.User) error {
	return dao.QueryN(pp, "select * from user")
}

func countAll(cnt *uint64) error {
	return dao.Query1(cnt, "select count(u_id) from user")
}

func auth(b *bool, uname *string, pwd *string) error {
	sql :=
		`select count(u_id) from user
	where u_uname = ? and u_pwd_encrypted = sha1(concat(?, 'jun990527'))`

	return dao.Query1(b, sql, uname, pwd)
}

func add(email, uname, pwd, desc *string,
	thumbnails *string, sex *dto.UserSex,
	birth *time.Time, tel *string,
	role *dto.UserRole) (int64, error) {

	sql :=
		`insert into user (u_email, u_uname, u_pwd_encrypted, u_desc, 
           u_thumbnails, u_sex, u_birth, u_tel, u_role, u_active_time, u_create_time)
	values(?, ?, sha1(concat(?, 'jun990527')), 
	       ifnull(?, ''), 
	       ifnull(?, ''), 
	       ifnull(?, 0),
	       ifnull(?, from_unixtime(0)),
	       ifnull(?, ''), 
	       ifnull(?, 0),
	       now(), now())`

	return dao.Insert(sql, email, uname, pwd, desc, thumbnails, sex, birth, tel, role)
}
