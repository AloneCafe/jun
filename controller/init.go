package controller

import (
	"jun/controller/auth"
	"jun/controller/base"
	"jun/controller/users"
	"jun/dto"
)

func init() {
	base.SetBasicController("/auth", &auth.AuthController{})
	base.SetBasicController("/users/me", &users.UsersMeController{})
	base.SetBasicController("/users", &users.UsersController{LowestRole: dto.U_ROLE_ADMIN})
	base.SetBasicController("/users/:uid", &users.UsersUidController{LowestRole: dto.U_ROLE_ADMIN})
	base.SetBasicController("/users/ban", &users.UsersBanController{LowestRole: dto.U_ROLE_ADMIN})
	base.SetBasicController("/users/:uid/ban", &users.UsersUidBanController{LowestRole: dto.U_ROLE_ADMIN})
}
