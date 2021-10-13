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
	GetLowestRole  dto.UserRole
	PostLowestRole dto.UserRole
}

func (p *RootController) DeleteHandler() gin.HandlerFunc {
	return nil
}

func (p *RootController) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前用户的文章
		var authorID int64
		if wc, err := base.Authorization(c, p.GetLowestRole); err != nil {
			return
		} else {
			authorID = wc.UID
		}

		if role, err := user.GetRoleById(authorID); err != nil {
			c.JSON(http.StatusInternalServerError,
				dto.NewResult(false, "获取当前用户的文章信息出错", nil))
			return
		} else if *role <= dto.U_ROLE_SUBCRIBER {
			c.JSON(http.StatusUnauthorized,
				dto.NewResult(false, "没有获取所属文章的权限", nil))
			return
		} else {
			if posts, err := post.GetAllNoBodyByUID(authorID); err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, "获取当前用户的文章信息出错", nil))
			} else {
				c.JSON(http.StatusOK,
					dto.NewResult(true, "获取当前用户的文章信息成功", posts))
			}
			return
		}
	}
}

func (p *RootController) PostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authorID int64
		if wc, err := base.Authorization(c, p.PostLowestRole); err != nil {
			return
		} else {
			authorID = wc.UID
		}

		var p dto.Post
		err := c.BindJSON(&p)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
			return
		} else {
			id, err := post.Add(p.Title, p.Desc, p.Body, authorID, p.Keywords,
				dto.DetachTagsIDs(p.Tags), dto.DetachCategoriesIDs(p.Categories), p.Type, p.Thumbnails)
			if err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, "文章添加失败", nil))
				return
			} else {
				_, err := post.GetByID(id)
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
