package router

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"jun/conf"
	_ "jun/controller"
	"jun/controller/base"
	"jun/dto"
	"jun/router/middleware"
	"jun/util"
	"net/http"
	"strconv"
	"time"
)

func Setup() {
	wc := conf.GetGlobalConfig().Web

	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，将 GIN_MODE 设置为 release
	// 默认 gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500
	r.Use(middleware.Recovery())

	r.Any("/ip", func(c *gin.Context) {
		c.String(http.StatusOK, c.ClientIP())
	})

	r.Any("/timestamp", func(c *gin.Context) {
		c.String(http.StatusOK, strconv.FormatInt(time.Now().Unix(), 10))
	})

	r.Any("/version", func(c *gin.Context) {
		c.String(http.StatusOK, util.GetApiVer())
	})

	rest := r.Group("/" + util.GetApiVer())

	rest.Use(middleware.ServerName())
	rest.Use(middleware.CORS())
	rest.Use(gzip.Gzip(gzip.DefaultCompression))

	bcs := base.GetBasicControllers()
	for _, c := range *bcs {
		fget := c.IBasicController.GetHandler()
		fpost := c.IBasicController.PostHandler()
		fput := c.IBasicController.PutHandler()
		fdelete := c.IBasicController.DeleteHandler()

		if fget != nil {
			rest.GET(string(c.ReqPoint), fget)
		}
		if fpost != nil {
			rest.POST(string(c.ReqPoint), fpost)
		}
		if fput != nil {
			rest.PUT(string(c.ReqPoint), fput)
		}
		if fdelete != nil {
			rest.DELETE(string(c.ReqPoint), fdelete)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, dto.NewResult(false, "请求的 API 路由不存在", nil))
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, dto.NewResult(false, "当前的 API 不接受该 HTTP 请求类型", nil))
	})

	uri := fmt.Sprintf("%s:%d", wc.BindAddr, wc.BindPort)
	if wc.HTTPS.Enabled {
		r.RunTLS(uri, wc.HTTPS.PemFile, wc.HTTPS.KeyFile)

	} else {
		r.Run(uri)
	}
}
