package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jun/controller/auth"
	"jun/dto"
	"jun/model/user"
	"net/http"
	"strings"
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

func authorization(c *gin.Context, lowRole dto.UserRole) (e error) {
	e = nil
	if lowRole <= dto.U_ROLE_VISITOR {
		return
	}

	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		e = errors.New("请求头授权字段为空")
		c.AbortWithError(http.StatusUnauthorized, e)
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		e = errors.New("请求头授权字段格式错误")
		c.AbortWithError(http.StatusBadRequest, e)
		return
	}

	if claims, err := auth.Check(parts[1]); err != nil {
		e = err
		c.AbortWithError(http.StatusUnauthorized, e)

	} else if role, err := user.GetRoleById(claims.UID); err != nil {
		e = errors.New("授权在获取 UID 时发生错误")
		c.AbortWithError(http.StatusUnauthorized, e)

	} else if *role < lowRole {
		e = errors.New("请求头授权字段权限不足")
		c.AbortWithError(http.StatusUnauthorized, e)

	}
	return
}

func init() {

}

func Dispatch2Vistor(g *gin.RouterGroup) {

}
