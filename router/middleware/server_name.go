package middleware

import (
	"github.com/gin-gonic/gin"
	"jun/utils/conf"
)

func ServerName() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("server", conf.GetGlobalConfig().Web.ServerName)
		c.Next()
	}
}
