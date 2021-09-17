package base

import (
	"github.com/gin-gonic/gin"
)

/*
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
*/

type IBasicController interface {
	DeleteHandler() gin.HandlerFunc
	GetHandler() gin.HandlerFunc
	PostHandler() gin.HandlerFunc
	PutHandler() gin.HandlerFunc
}

type IAdvancedController interface {
	IBasicController
	ConnectHandler() gin.HandlerFunc
	HeadHandler() gin.HandlerFunc
	OptionsHandler() gin.HandlerFunc
	PatchHandler() gin.HandlerFunc
	TraceHandler() gin.HandlerFunc
}

type ReqPoint string

type BasicControllerEntry struct {
	ReqPoint
	IBasicController
}

var (
	//controllers = make(map[ReqPoint]*ReqController)
	controllers []BasicControllerEntry
)

/*
func setController(point ReqPoint, controller *ReqController) {
	controllers[point] = controller
}

func getController(point ReqPoint) (controller *ReqController) {
	return controllers[point]
}
*/

func SetBasicController(point ReqPoint, controller IBasicController) {
	controllers = append(controllers, BasicControllerEntry{
		ReqPoint:         point,
		IBasicController: controller,
	})
}

/*
func GetController(point ReqPoint) (controller IBasicController) {
	return controllers[point]
}
*/

func GetBasicControllers() *[]BasicControllerEntry {
	return &controllers
}
