package user

import (
	"errors"
	"jun/dto"
	"time"
)

func GetById(id int64) (*dto.User, error) {
	p := new(dto.User)
	err := selectById(p, &id)
	return p, err
}

func GetByUname(uname string) (*dto.User, error) {
	p := new(dto.User)
	err := selectByUname(p, &uname)
	return p, err
}

func GetByEmail(email string) (*dto.User, error) {
	p := new(dto.User)
	err := selectByEmail(p, &email)
	return p, err
}

func GetRoleById(id int64) (*dto.UserRole, error) {
	p, err := GetById(id)
	if err != nil {
		return nil, err
	} else {
		return p.Role, nil
	}
}

func GetAll() (*[]dto.User, error) {
	pp := new([]dto.User)
	err := selectAll(pp)
	return pp, err
}

func CountAll() (*uint64, error) {
	p := new(uint64)
	err := countAll(p)
	return p, err
}

func Auth(uname *string, pwd *string) (*bool, error) {
	b := new(bool)
	err := auth(b, uname, pwd)
	return b, err
}

func Add(email, uname, pwd, desc *string,
	thumbnails *string, sex *dto.UserSex,
	birth *time.Time, tel *string,
	role *dto.UserRole) (int64, error) {

	if email == nil || *email == "" {
		return 0, errors.New("电子邮件不能为空")
	} else if uname == nil || *uname == "" {
		return 0, errors.New("用户名不能为空")
	} else if pwd == nil || *pwd == "" {
		return 0, errors.New("密码不能为空")
	} else {
		return add(email, uname, pwd, desc, thumbnails, sex, birth, tel, role)
	}
}
