package user

import (
	"jun/dto"
)

func GetByUid(uid int64) (*dto.User, error) {
	p := new(dto.User)
	err := selectAllByUid(p, uid)
	return p, err
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

func Add(p *dto.User) (int64, error) {
	return add(p)
}
