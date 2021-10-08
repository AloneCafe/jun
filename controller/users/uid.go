package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"jun/controller/base"
	"jun/dto"
	"jun/model/user"
)

type UidController struct {
	LowestRole dto.UserRole
}

func (p *UidController) DeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var aid int64
		if wc, err := base.Authorization(c, p.LowestRole); err != nil {
			return
		} else {
			aid = wc.UID
		}

		if id, err := strconv.ParseInt(c.Param("uid"), 10, 64); err != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
		} else {
			var err error
			if id == aid {
				c.JSON(http.StatusBadRequest,
					dto.NewResult(false, "危险操作，无法删除管理员账户", nil))
				return
			} else {
				_, err = user.DeleteById(id)
			}

			if err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, fmt.Sprintf("用户删除失败，id = %d", id), nil))
			} else {
				_, err := user.GetById(id)
				if err != nil {
					c.JSON(http.StatusOK,
						dto.NewResult(true, "用户删除成功", nil))
				} else {
					c.JSON(http.StatusInternalServerError,
						dto.NewResult(false, fmt.Sprintf("用户删除失败，id = %d", id), nil))
				}
			}
		}
	}
}

func (p *UidController) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := base.Authorization(c, p.LowestRole); err != nil {
			return
		}

		if id, err := strconv.ParseInt(c.Param("uid"), 10, 64); err != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
			return

		} else {
			u, err := user.GetById(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, fmt.Sprintf("获取用户信息出错，id = %d", id), nil))
				return

			} else {
				c.JSON(http.StatusOK,
					dto.NewResult(true, "获取用户信息成功", u))
				return
			}
		}
	}
}

func (p *UidController) PostHandler() gin.HandlerFunc {
	return nil
}

func (p *UidController) PutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var aid int64
		if wc, err := base.Authorization(c, p.LowestRole); err != nil {
			return
		} else {
			aid = wc.UID
		}

		var u dto.UserInfoAllUpdate
		err := c.BindJSON(&u)
		var err2 error
		u.IDReadOnly, err2 = strconv.ParseInt(c.Param("uid"), 10, 64)
		if err != nil || err2 != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
		} else if u.Uname == nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "用户数据序列化失败", nil))
		} else {
			var err error
			if u.IDReadOnly == aid { // 此处需要判断，管理员是无法更改自己的角色的（不允许自我降职）
				_, err = user.UpdateBasicInfo(&dto.UserInfoBasicUpdate{
					IDReadOnly: u.IDReadOnly,
					Email:      u.Email,
					Uname:      u.Uname,
					Pwd:        u.Pwd,
					Desc:       u.Desc,
					Thumbnails: u.Thumbnails,
					Sex:        u.Sex,
					Birth:      u.Birth,
					Tel:        u.Tel,
				})
			} else {
				_, err = user.UpdateAllInfo(&u)
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, fmt.Sprintf("用户数据更新失败，id = %d", u.IDReadOnly), nil))
			} else {
				_, err := user.GetById(u.IDReadOnly)
				if err != nil {
					c.JSON(http.StatusInternalServerError,
						dto.NewResult(false, fmt.Sprintf("用户数据更新失败，id = %d", u.IDReadOnly), nil))
				} else {
					c.JSON(http.StatusOK,
						dto.NewResult(true, "用户数据更新成功", nil))
				}
			}
		}
	}
}
