package users

import (
	"github.com/gin-gonic/gin"
	"jun/controller/base"
	"jun/dto"
	"jun/model/user"
	"net/http"
)

type UsersBanController struct {
	LowestRole dto.UserRole
}

func (p *UsersBanController) DeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := base.Authorization(c, p.LowestRole); err != nil {
			return
		}

		_, err := user.DeleteAllBanned()
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				dto.NewResult(false, "用户解封失败", nil))
		} else {
			c.JSON(http.StatusOK,
				dto.NewResult(true, "已解封全部的封禁用户", nil))
		}
	}
}

func (p *UsersBanController) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := base.Authorization(c, p.LowestRole); err != nil {
			return
		}

		// 获取全部封禁用户的信息
		users, err := user.GetAllBanned()
		if err != nil || users == nil {
			c.JSON(http.StatusInternalServerError,
				dto.NewResult(false, "获取全部封禁用户的信息出错", nil))
			return

		} else {
			c.JSON(http.StatusOK,
				dto.NewResult(true, "获取全部封禁用户的信息成功", *users))
			return
		}
	}
}

func (p *UsersBanController) PostHandler() gin.HandlerFunc {
	return nil
}

func (p *UsersBanController) PutHandler() gin.HandlerFunc {
	return nil
}
