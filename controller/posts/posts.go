package posts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"jun/controller/base"
	"jun/dto"
	"jun/model/post"
	"jun/model/user"
)

type RootController struct {
	PostLowestRole dto.UserRole
}

func (p *RootController) DeleteHandler() gin.HandlerFunc {
	return nil
}

func (p *RootController) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (p *RootController) PostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := base.Authorization(c, p.PostLowestRole); err != nil {
			return
		}

		var u dto.Post
		err := c.BindJSON(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
			return
		} else {
			id, err := post.Add(u.Email, u.Uname, u.PwdEncrypted, u.Desc,
				u.Thumbnails, u.Sex, u.Birth, u.Tel, u.Role)
			if err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, "文章添加失败", nil))
				return
			} else {
				_, err := post.GetById(id)
				if err != nil {
					c.JSON(http.StatusInternalServerError,
						dto.NewResult(false, fmt.Sprintf("获取文章信息出错，id = %d", id), nil))
					return
				} else {
					c.JSON(http.StatusOK,
						dto.NewResult(true, "文章添加成功", nil))
					return
				}
			}
		}
	}
}

func (p *RootController) PutHandler() gin.HandlerFunc {
	return nil
}
