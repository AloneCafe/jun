package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin",
			"*") // 可将 * 替换为指定的域名

		c.Header("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH, CONNECT")

		c.Header("Access-Control-Allow-Headers",
			"Content-Type, AccessToken, X-CSRF-Token, Token, Origin, X-Requested-With, Accept, Authorization")

		c.Header("Access-Control-Expose-Headers",
			"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")

		c.Header("Access-Control-Allow-Credentials",
			"true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

/*
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		//主要设置Access-Control-Allow-Origin
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "false")
		c.Set("content-type", "application/json")

		if method == "OPTIONS"{
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
*/
