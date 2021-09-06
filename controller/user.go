package controller

import (
	"github.com/gin-gonic/gin"
	"jun/dto"
	"jun/model/user"
	"net/http"
	"strconv"
)

var (
	cGetUser = func(lowRole dto.UserRole) func(*gin.Context) {
		return func(c *gin.Context) {
			authorization(c, lowRole)

			idStr := c.Query("id")
			if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
				panic(err)
			} else {
				u, err := user.GetById(id)
				if err != nil {
					panic(err)
				} else {
					c.JSON(http.StatusOK, u)
				}
			}
		}
	}

	cAddUser = func(lowRole dto.UserRole) func(*gin.Context) {
		return func(c *gin.Context) {
			authorization(c, lowRole)

			var u dto.User
			err := c.BindJSON(&u)
			if err != nil {
				panic(err)
			} else {
				id, err := user.Add(u.Email, u.Uname, u.PwdEncrypted, u.Desc,
					u.Thumbnails, u.Sex, u.Birth, u.Tel, u.Role)
				if err != nil {
					panic(err)
				} else {
					_, err := user.GetById(id)
					if err != nil {
						panic(err)
					} else {
						c.JSON(http.StatusOK,
							dto.NewResult(true, "用户添加成功", nil))
					}
				}
			}
		}
	}

	cDeleteUser = func(lowRole dto.UserRole) func(*gin.Context) {
		return func(c *gin.Context) {
			authorization(c, lowRole)

			idStr := c.Query("id")
			if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
				panic(err)
			} else {
				id, err := user.DeleteById(id)
				if err != nil {
					panic(err)
				} else {
					_, err := user.GetById(id)
					if err != nil {
						panic(err)
					} else {
						c.JSON(http.StatusOK,
							dto.NewResult(true, "用户删除成功", nil))
					}
				}
			}
		}
	}

	cUpdateUser = func(lowRole dto.UserRole) func(*gin.Context) {
		return func(c *gin.Context) {
			authorization(c, lowRole)

			var u dto.User
			err := c.BindJSON(&u)
			if err != nil {
				panic(err)
			} else {
				_, err := user.UpdateAllInfo(&u)
				if err != nil {
					panic(err)
				} else {
					_, err := user.GetById(*u.ID)
					if err != nil {
						panic(err)
					} else {
						c.JSON(http.StatusOK,
							dto.NewResult(true, "用户数据更新成功", nil))
					}
				}
			}
		}
	}
)

func init() {
	setController("/users/current", &ReqController{
		ConnectHandler: nil,
		DeleteHandler:  cDeleteUserCurr(),
		GetHandler:     cGetUserCurr(),
		HeadHandler:    nil,
		OptionsHandler: nil,
		PatchHandler:   nil,
		PostHandler:    nil,
		PutHandler:     cUpdateUserCurr(),
		TraceHandler:   nil,
	})

	setController("/users", &ReqController{
		ConnectHandler: nil,
		DeleteHandler:  cDeleteUser(dto.U_ROLE_ADMIN),
		GetHandler:     cGetUser(dto.U_ROLE_ADMIN),
		HeadHandler:    nil,
		OptionsHandler: nil,
		PatchHandler:   nil,
		PostHandler:    cAddUser(dto.U_ROLE_ADMIN),
		PutHandler:     cUpdateUser(dto.U_ROLE_ADMIN),
		TraceHandler:   nil,
	})

}
