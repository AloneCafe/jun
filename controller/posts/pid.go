package posts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"jun/controller/base"
	"jun/dto"
	"jun/model/post"
	"jun/model/user"
)

type PidController struct {
	DeleteLowestRole dto.UserRole
	GetLowestRole    dto.UserRole
	PutLowestRole    dto.UserRole
}

func (p *PidController) DeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := base.Authorization(c, p.DeleteLowestRole); err != nil {
			return
		}

		if id, err := strconv.ParseInt(c.Param("pid"), 10, 64); err != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
		} else {
			if _, err := post.DeleteByID(id); err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, fmt.Sprintf("删除文章失败，id = %d", id), nil))
			} else {
				_, err := post.GetByID(id)
				if err == nil {
					c.JSON(http.StatusOK,
						dto.NewResult(true, "删除文章成功", nil))
				} else {
					c.JSON(http.StatusInternalServerError,
						dto.NewResult(false, fmt.Sprintf("删除文章失败，id = %d", id), nil))
				}
			}
		}
	}
}

func (p *PidController) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := base.Authorization(c, p.GetLowestRole); err != nil {
			return
		}

		if id, err := strconv.ParseInt(c.Param("pid"), 10, 64); err != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
		} else {
			if _, err := post.DeleteByID(id); err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, fmt.Sprintf("删除文章失败，id = %d", id), nil))
			} else {
				u, err := post.GetByID(id)
				if err != nil {
					c.JSON(http.StatusOK,
						dto.NewResult(true, "获取用户信息成功", u))
					return
				} else {
					c.JSON(http.StatusInternalServerError,
						dto.NewResult(false, fmt.Sprintf("获取用户信息出错，id = %d", id), nil))
					return
				}
			}
		}
	}
}

func (p *PidController) PostHandler() gin.HandlerFunc {
	return nil
}

func (p *PidController) PutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var aid int64
		if wc, err := base.Authorization(c, p.PutLowestRole); err != nil {
			return
		} else {
			aid = wc.UID
		}

		var u dto.PostInfoUpdate
		err := c.BindJSON(&u)
		var err2 error
		u.PIDReadOnly, err2 = strconv.ParseInt(c.Param("pid"), 10, 64) // 覆盖 ID 值
		if err != nil || err2 != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
		} else if u.Title == nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "文章数据序列化失败", nil))
		} else {
			var err error
			var thisRole dto.UserRole
			if role, err := user.GetRoleById(aid); err != nil {
				thisRole = *role
			} else {
				c.JSON(http.StatusBadRequest,
					dto.NewResult(false, "参数不正确", nil))
				return
			}

			if u.AuthorID == aid || thisRole == dto.U_ROLE_ADMIN { // 此处需要判断，文章的作者才有自我修改权，或者当前用户是管理员
				_, err = post.UpdateInfo(&u)
			} else {
				c.JSON(http.StatusUnauthorized,
					dto.NewResult(false, "当前用户没有修改权限", nil))
				return
			}

			if err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, fmt.Sprintf("文章数据更新失败，id = %d", u.PIDReadOnly), nil))
			} else {
				_, err := user.GetById(u.PIDReadOnly)
				if err != nil {
					c.JSON(http.StatusInternalServerError,
						dto.NewResult(false, fmt.Sprintf("文章数据更新失败，id = %d", u.PIDReadOnly), nil))
				} else {
					c.JSON(http.StatusOK,
						dto.NewResult(true, "文章数据更新成功", nil))
				}
			}
		}
	}
}
