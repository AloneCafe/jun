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
		return &p.Role, nil
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

func AuthGetUid(uname *string, pwd *string) (bool, int64, error) {
	id := new(int64)
	ok, err := authGetUid(id, uname, pwd)
	return ok, *id, err
}

func Add(email, uname, pwd, desc *string,
	thumbnails *string, sex dto.UserSex,
	birth *time.Time, tel *string,
	role dto.UserRole) (int64, error) {

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

func DeleteById(id int64) (int64, error) {
	return deleteById(id)
}

// UpdateBasicInfo 更新除了 Role 和 UID 以外的字段
func UpdateBasicInfo(p *dto.UserInfoBasicUpdate) (int64, error) {
	return updateBasicInfo(p)
}

// UpdateAllInfo 更新除了 UID 以外的字段
func UpdateAllInfo(p *dto.UserInfoAllUpdate) (int64, error) {
	return updateAllInfo(p)
}

func IsUserBannedById(id int64) (bool, error) {
	return isUserBannedById(id)
}

func BanUserById(id int64) (int64, error) {
	return banUserById(id)
}

func UnbanUserById(id int64) (int64, error) {
	return unbanUserById(id)
}

func DeleteAllBanned() (int64, error) {
	return deleteAllBanned()
}

func GetAllBanned() (*[]dto.User, error) {
	pp := new([]dto.User)
	err := selectAllBanned(pp)
	return pp, err
}
