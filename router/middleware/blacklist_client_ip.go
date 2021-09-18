package middleware

import (
	"github.com/gin-gonic/gin"
	"jun/dto"
	"jun/utils/conf"
	"net/http"
)

func BL2ClientIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		if conf.IsIPBanned(c.ClientIP()) {
			c.AbortWithStatusJSON(http.StatusForbidden,
				dto.NewResult(false, "审计规则生效，访问被拒绝", nil))
		} else {
			c.Next()
		}
	}
}
