package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jun/controller/base"
	"jun/dto"
	"jun/model/user"
	"net/http"
)

type UsersController struct {
	LowestRole dto.UserRole
	//base.IBasicController
}

func (p *UsersController) DeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := base.Authorization(c, p.LowestRole); err != nil {
			return
		}

		c.JSON(http.StatusBadRequest,
			dto.NewResult(false, "危险操作，无法删除全部用户", nil))
	}
}

func (p *UsersController) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := base.Authorization(c, p.LowestRole); err != nil {
			return
		}

		// 获取全部用户信息
		users, err := user.GetAll()
		if err != nil || users == nil {
			c.JSON(http.StatusInternalServerError,
				dto.NewResult(false, "获取全部用户信息出错", nil))
			return

		} else {
			c.JSON(http.StatusOK,
				dto.NewResult(true, "获取全部用户信息成功", *users))
			return
		}
	}
}

func (p *UsersController) PostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := base.Authorization(c, p.LowestRole); err != nil {
			return
		}

		var u dto.User
		err := c.BindJSON(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
			return
		} else {
			id, err := user.Add(u.Email, u.Uname, u.PwdEncrypted, u.Desc,
				u.Thumbnails, u.Sex, u.Birth, u.Tel, u.Role)
			if err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, "用户添加失败", nil))
				return
			} else {
				_, err := user.GetById(id)
				if err != nil {
					c.JSON(http.StatusInternalServerError,
						dto.NewResult(false, fmt.Sprintf("获取用户信息出错，id = %d", id), nil))
					return
				} else {
					c.JSON(http.StatusOK,
						dto.NewResult(true, "用户添加成功", nil))
					return
				}
			}
		}
	}
}

func (p *UsersController) PutHandler() gin.HandlerFunc {
	return nil
}
