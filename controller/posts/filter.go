package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"jun/dto"
	"jun/model/option"
	"jun/model/post"
)

type MatchListController struct {
	LowestRole dto.UserRole
}

func (p *MatchListController) DeleteHandler() gin.HandlerFunc {
	return nil
}

func (p *MatchListController) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		match := c.Param("match")
		pageIndex := c.Param("pageIndex")
		if match == "" { //
			match = ".*"
		}
		if pageIndex == "" {
			pageIndex = "0"
		}
		sizeOfPage, err := option.GetPostCountPerPage()
		if err != nil {
			sizeOfPage = 10
		}
		pi, err := strconv.ParseInt(pageIndex, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
			return
		}

		findPost, err := post.FindPost(match, match, match, sizeOfPage, pi)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				dto.NewResult(false, "获取文章信息出错", nil))
		} else {
			c.JSON(http.StatusOK,
				dto.NewResult(true, "获取文章信息成功", findPost))
		}

	}
}

func (p *MatchListController) PostHandler() gin.HandlerFunc {
	return nil
}

func (p *MatchListController) PutHandler() gin.HandlerFunc {
	return nil
}
