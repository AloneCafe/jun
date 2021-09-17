package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jun/controller/base"
	"jun/dto"
	"jun/model/user"
	"net/http"
)

type UsersMeController struct {
	//base.IBasicController
}

func (*UsersMeController) DeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 至少是订阅者级别，才有个人信息呀
		var id int64
		if wc, err := base.Authorization(c, dto.U_ROLE_SUBCRIBER); err != nil {
			return
		} else {
			id = wc.UID
		}
		_, err := user.DeleteById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				dto.NewResult(false, fmt.Sprintf("当前用户注销失败，id = %d", id), nil))
		} else {
			_, err := user.GetById(id)
			if err != nil {
				c.JSON(http.StatusOK,
					dto.NewResult(true, "当前用户注销成功", nil))
			} else {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, fmt.Sprintf("当前用户注销失败，id = %d", id), nil))
			}
		}
	}
}
func (*UsersMeController) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 至少是订阅者级别，才有个人信息呀
		var id int64
		if wc, err := base.Authorization(c, dto.U_ROLE_SUBCRIBER); err != nil {
			return
		} else {
			id = wc.UID
		}
		u, err := user.GetById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				dto.NewResult(false, fmt.Sprintf("获取当前用户信息出错，id = %d", id), nil))
			return

		} else {
			c.JSON(http.StatusOK,
				dto.NewResult(true, "获取当前用户信息成功", u))
			return
		}
	}
}

func (*UsersMeController) PostHandler() gin.HandlerFunc {
	return nil
}

func (*UsersMeController) PutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 至少是订阅者级别，才有个人信息呀
		var id int64
		if wc, err := base.Authorization(c, dto.U_ROLE_SUBCRIBER); err != nil {
			return
		} else {
			id = wc.UID
		}
		var u dto.UserInfoBasicUpdate
		err := c.BindJSON(&u)
		u.IDReadOnly = id // 替换为 token 里面 payload 提供的 id，以免恶意篡改更改到其他用户的信息
		if err != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
		} else if u.Uname == nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "当前用户数据序列化失败", nil))
		} else {
			_, err := user.UpdateBasicInfo(&u) // 不能更改角色哦
			if err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, fmt.Sprintf("当前用户数据更新失败，id = %d", u.IDReadOnly), nil))
			} else {
				_, err := user.GetById(u.IDReadOnly)
				if err != nil {
					c.JSON(http.StatusInternalServerError,
						dto.NewResult(false, fmt.Sprintf("当前用户数据更新失败，id = %d", u.IDReadOnly), nil))
				} else {
					c.JSON(http.StatusOK,
						dto.NewResult(true, "用户数据更新成功", nil))
				}
			}
		}
	}
}
