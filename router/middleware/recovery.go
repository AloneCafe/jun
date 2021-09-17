package middleware

import (
	"github.com/gin-gonic/gin"
	"jun/dto"
	"log"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Print("Server painc! ")
				log.Println(err)
				c.JSON(http.StatusInternalServerError,
					dto.NewResult(false, "服务器程序发生未知的 painc 故障，当前无法正确处理请求", nil))
				return
			}
		}()
		c.Next()
	}
}
