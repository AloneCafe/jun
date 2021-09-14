package controller

import (
	"github.com/gin-gonic/gin"
)

type ReqController struct {
	ConnectHandler gin.HandlerFunc
	DeleteHandler  gin.HandlerFunc
	GetHandler     gin.HandlerFunc
	HeadHandler    gin.HandlerFunc
	OptionsHandler gin.HandlerFunc
	PatchHandler   gin.HandlerFunc
	PostHandler    gin.HandlerFunc
	PutHandler     gin.HandlerFunc
	TraceHandler   gin.HandlerFunc
}

type ReqPoint string

var (
	controllers map[ReqPoint]*ReqController
)

func setController(point ReqPoint, controller *ReqController) {
	controllers[point] = controller
}

func getController(point ReqPoint) (controller *ReqController) {
	return controllers[point]
}

func GetMap() *map[ReqPoint]*ReqController {
	return &controllers
}

func init() {

}

func Dispatch2Vistor(g *gin.RouterGroup) {

}
