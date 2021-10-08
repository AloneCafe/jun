package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jun/controller/base"
	"jun/dto"
	"jun/model/user"
	"net/http"
	"strconv"
)

type UidBanController struct {
	LowestRole dto.UserRole
}

func (p *UidBanController) DeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := base.Authorization(c, p.LowestRole); err != nil {
			return
		}

		id, err := strconv.ParseInt(c.Param("uid"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
		} else {
			_, err := user.UnbanUserById(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, fmt.Sprintf("用户解封失败，id = %d", id), nil))
			} else {
				/*
					banned, err := user.IsUserBannedById(id)
					if err != nil {
						c.JSON(http.StatusInternalServerError,
							dto.NewResult(false, fmt.Sprintf("获取用户封禁状态失败，id = %d", id), nil))
					} else if banned {
						c.JSON(http.StatusInternalServerError,
							dto.NewResult(false, fmt.Sprintf("用户解封失败，id = %d", id), nil))
					} else {
						c.JSON(http.StatusOK,
							dto.NewResult(true, "用户解封成功", nil))
					}
				*/
				c.JSON(http.StatusOK,
					dto.NewResult(true, "用户解封成功", nil))
			}
		}
	}
}

func (p *UidBanController) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := base.Authorization(c, p.LowestRole); err != nil {
			return
		}

		id, err := strconv.ParseInt(c.Param("uid"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
		} else {
			banned, err := user.IsUserBannedById(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, fmt.Sprintf("获取用户封禁状态失败，id = %d", id), nil))
			} else if !banned {
				c.JSON(http.StatusOK,
					dto.NewResult(true, fmt.Sprintf("获取用户封禁状态成功，该用户未被封禁，id = %d", id),
						dto.UserBanned{Banned: false}))
			} else {
				c.JSON(http.StatusOK,
					dto.NewResult(true, fmt.Sprintf("获取用户封禁状态成功，该用户已被封禁，id = %d", id),
						dto.UserBanned{Banned: true}))
			}
		}
	}
}

func (p *UidBanController) PostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var aid int64
		if wc, err := base.Authorization(c, p.LowestRole); err != nil {
			return
		} else {
			aid = wc.UID
		}

		id, err := strconv.ParseInt(c.Param("uid"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "参数不正确", nil))
		} else {
			if aid == id {
				c.JSON(http.StatusBadRequest,
					dto.NewResult(false, fmt.Sprintf("无法封禁自身（管理员用户），id = %d", id), nil))
				return
			}

			_, err := user.BanUserById(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, fmt.Sprintf("用户封禁失败，id = %d", id), nil))
			} else {
				/*
					banned, err := user.IsUserBannedById(id)
					if err != nil {
						c.JSON(http.StatusInternalServerError,
							dto.NewResult(false, fmt.Sprintf("获取用户封禁状态失败，id = %d", id), nil))
					} else if !banned {
						c.JSON(http.StatusInternalServerError,
							dto.NewResult(false, fmt.Sprintf("用户封禁失败，id = %d", id), nil))
					} else {
						c.JSON(http.StatusOK,
							dto.NewResult(true, "用户封禁成功", nil))
					}
				*/
				c.JSON(http.StatusOK,
					dto.NewResult(true, "用户封禁成功", nil))
			}
		}
	}
}

func (p *UidBanController) PutHandler() gin.HandlerFunc {
	return nil
}
