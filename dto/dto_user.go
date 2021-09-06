package dto

type UserRole int32

const (
	U_ROLE_VISTOR = iota
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
	ID           *int64   `json:"u_id"            db:"u_id"`
	Email        *string   `json:"u_email"         db:"u_email"`
	Uname        *string   `json:"u_uname"         db:"u_uname"`
	PwdEncrypted *string   `json:"u_pwd_encrypted" db:"u_pwd_encrypted"`
	Desc         *string   `json:"u_desc"          db:"u_desc"`
	Thumbnails   *string   `json:"u_thumbnails"    db:"u_thumbnails"`
	Sex          *UserSex  `json:"u_sex"           db:"u_sex"`
	Tags         *string   `json:"u_tags"          db:"u_tags"`
	Role         *UserRole `json:"u_role"          db:"u_role"`
}
/*
func NewUser(id uint64, email *string, uname *string, pwdEncrypted *string,
	desc *string, thumbnails *string, sex *UserSex, tags *string, role *UserRole) *User {
	return &User{id, *email, *uname, *pwdEncrypted,
		*desc, *thumbnails, *sex, *tags, *role}
}
*/