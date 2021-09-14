package router

import (
	"github.com/gin-gonic/gin"
	"jun/controller"
)

func Setup() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，将 GIN_MODE 设置为 release
	// 默认 gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500
	r.Use(gin.Recovery())

	// 为每个路由添加中间件
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	// 认证路由组
	// authorized := r.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样:

	v0 := r.Group("/v0")
	//v1 := r.Group("/v1")
	//v2 := r.Group("/v2")
	//v3 := r.Group("/v3")
	//v4 := r.Group("/v4")
	//v5 := r.Group("/v5")
	//
	//Dispatch2Vistor(v0)
	//Dispatch2Subcriber(&v1)
	//Dispatch2Contributor(&v2)
	//Dispatch2Author(&v3)
	//Dispatch2Editor(&v4)
	//Dispatch2Admin(&v5)

	// 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的
	// AuthRequired() 中间件
	//a := v0.Use(AuthRequired())
	//{
	//	a.POST("/login", loginEndpoint)
	//	a.POST("/submit", submitEndpoint)
	//	a.POST("/read", readEndpoint)
	//	// 嵌套路由组
	//	testing := authorized.Group("testing")
	//	testing.GET("/analytics", analyticsEndpoint)
	//}

	controllers := controller.GetMap()
	for point, handler := range *controllers {
		if handler.GetHandler != nil {
			v0.GET(string(point), handler.GetHandler)
		} else if handler.PostHandler != nil {
			v0.POST(string(point), handler.PostHandler)
		} else if handler.PutHandler != nil {
			v0.PUT(string(point), handler.PutHandler)
		} else if handler.DeleteHandler != nil {
			v0.DELETE(string(point), handler.DeleteHandler)
		} else if handler.HeadHandler != nil {
			v0.HEAD(string(point), handler.HeadHandler)
		} else if handler.PatchHandler != nil {
			v0.PATCH(string(point), handler.PatchHandler)
		} else if handler.TraceHandler != nil {
			//v0.TRACE(string(point), handler.TraceHandler)
			panic("路由 \"" + point + "\" 不支持 TRACE 方法")
		} else if handler.ConnectHandler != nil {
			//v0.CONNECT(string(point), handler.ConnectHandler)
			panic("路由 \"" + point + "\" 不支持 CONNECT 方法")
		} else if handler.OptionsHandler != nil {
			v0.OPTIONS(string(point), handler.OptionsHandler)
		}
	}

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
