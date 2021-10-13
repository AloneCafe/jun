package controller

import (
	"jun/controller/auth"
	"jun/controller/base"
	"jun/controller/posts"
	"jun/controller/users"
	"jun/dto"
)

func init() {
	base.SetBasicController("/auth", &auth.AuthController{})

	base.SetBasicController("/users/me", &users.MeController{})
	base.SetBasicController("/users", &users.RootController{LowestRole: dto.U_ROLE_ADMIN})
	base.SetBasicController("/users/:uid", &users.UidController{LowestRole: dto.U_ROLE_ADMIN})
	base.SetBasicController("/users/ban", &users.BanController{LowestRole: dto.U_ROLE_ADMIN})
	base.SetBasicController("/users/:uid/ban", &users.UidBanController{LowestRole: dto.U_ROLE_ADMIN})

	base.SetBasicController("/posts/match/:match/page/:pageIndex", &posts.MatchListController{
		LowestRole: dto.U_ROLE_VISITOR,
	})
	base.SetBasicController("/posts/match/:match/page/", &posts.MatchListController{
		LowestRole: dto.U_ROLE_VISITOR,
	})
	base.SetBasicController("/posts/page/:pageIndex", &posts.MatchListController{
		LowestRole: dto.U_ROLE_VISITOR,
	})
	base.SetBasicController("/posts/page/", &posts.MatchListController{
		LowestRole: dto.U_ROLE_VISITOR,
	})
	base.SetBasicController("/posts/:pid", &posts.PidController{
		GetLowestRole:    dto.U_ROLE_SUBCRIBER,
		PutLowestRole:    dto.U_ROLE_ADMIN,
		DeleteLowestRole: dto.U_ROLE_ADMIN,
	})

}
