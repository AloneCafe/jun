package dto

import "time"

type UserRole int32
type UserRoles int32 // UserRole ^ 2

const (
	U_ROLE_VISITOR = iota
	U_ROLE_SUBCRIBER
	U_ROLE_CONTRIBUTOR
	U_ROLE_AUTHOR
	U_ROLE_EDITOR
	U_ROLE_ADMIN
)

type UserSex int32

const (
	U_SEX_SECRET = iota
	U_SEX_MALE
	U_SEX_FEMALE
	U_SEX_BINARY
)

type User struct {
	ID           int64      `json:"u_id"            db:"u_id"`
	Email        *string    `json:"u_email"         db:"u_email"`
	Uname        *string    `json:"u_uname"         db:"u_uname"`
	PwdEncrypted *string    `json:"u_pwd_encrypted" db:"u_pwd_encrypted"`
	Desc         *string    `json:"u_desc"          db:"u_desc"`
	Thumbnails   *string    `json:"u_thumbnails"    db:"u_thumbnails"`
	Sex          UserSex    `json:"u_sex"           db:"u_sex"`
	Birth        *time.Time `json:"u_birth"         db:"u_birth"`
	Tel          *string    `json:"u_tel"           db:"u_tel"`
	Role         UserRole   `json:"u_role"          db:"u_role"`
	ActiveTime   *time.Time `json:"u_active_time"   db:"u_active_time"`
	CreateTime   *time.Time `json:"u_create_time"   db:"u_create_time"`
}

type UserInfoBasicUpdate struct {
	IDReadOnly int64      `json:"u_id"` // 仅用于校验
	Email      *string    `json:"u_email"`
	Uname      *string    `json:"u_uname"`
	Pwd        *string    `json:"u_pwd"`
	Desc       *string    `json:"u_desc"`
	Thumbnails *string    `json:"u_thumbnails"`
	Sex        UserSex    `json:"u_sex"`
	Birth      *time.Time `json:"u_birth"`
	Tel        *string    `json:"u_tel"`
}

type UserInfoAllUpdate struct {
	UserInfoBasicUpdate
	Role UserRole `json:"u_role"`
}

type UserBanned struct {
	Banned bool `json:"banned"`
}

func NewUser(id int64, email string, uname string, pwdEncrypted string,
	desc string, thumbnails string,
	sex UserSex, birth time.Time, tel string, role UserRole, activeTime time.Time, createTime time.Time) *User {
	return &User{id, &email, &uname, &pwdEncrypted,
		&desc, &thumbnails,
		sex, &birth, &tel, role, &activeTime, &createTime}
}
