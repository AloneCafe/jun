package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jun/dto"
	"jun/model/user"
	"jun/util"
	"net/http"
	"strings"
)

func getBearerToken(c *gin.Context) (token *string, e error) {
	token = nil
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
	token = &parts[1]
	return
}

func authorization(c *gin.Context, lowRole dto.UserRole) (e error) {
	e = nil
	if lowRole <= dto.U_ROLE_VISITOR {
		return
	}

	token, err := getBearerToken(c)
	if err != nil {
		e = err
		c.AbortWithError(http.StatusUnauthorized, e)

	} else if claims, err := util.Check(*token); err != nil {
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

var (
	cLogout = func() func(*gin.Context) {
		return func(c *gin.Context) {
			token, err := getBearerToken(c)
			if err != nil {
				panic(err)
			} else {
				util.BanToken(*token)
				c.JSON(http.StatusOK,
					dto.NewResult(true, "用户已注销", nil))
			}
		}
	}

	cCheck = func() func(*gin.Context) {
		return func(c *gin.Context) {
			token, err := getBearerToken(c)
			if err != nil {
				panic(err)
			} else {
				_, err := util.Check(*token)
				if err != nil {
					c.JSON(http.StatusOK,
						dto.NewResult(false, "授权凭据无效", nil))
				} else {
					c.JSON(http.StatusOK,
						dto.NewResult(true, "授权凭据有效", nil))
				}
			}
		}
	}

	cLogin = func() func(*gin.Context) {
		return func(c *gin.Context) {
			username := c.PostForm("username")
			password := c.PostForm("password")
			ok, token, err := util.Login(username, password)
			if !ok || err != nil {
				panic(err)
			} else {
				c.JSON(http.StatusOK,
					dto.NewResult(true, "用户登录成功", map[string]string{
						"token": token,
					}))
			}
		}
	}

	cUpdate = func() func(*gin.Context) {
		return func(c *gin.Context) {
			token, err := getBearerToken(c)
			if err != nil {
				panic(err)
			} else {
				claims, err := util.Check(*token)
				if err != nil {
					c.JSON(http.StatusOK,
						dto.NewResult(false, "授权凭据过期或无效", nil))

				} else {
					// 尝试再次登录（用户名密码）
					ok, token, err := util.Login(claims.Uname, claims.Pwd)
					if !ok || err != nil {
						c.JSON(http.StatusOK,
							dto.NewResult(false, "授权凭据已变更，无法更新", nil))
					} else {
						c.JSON(http.StatusOK,
							dto.NewResult(true, "授权凭据更新成功", map[string]string{
								"token": token,
							}))
					}
				}
			}
		}
	}
)

func init() {

	setController("/auth", &ReqController{
		ConnectHandler: nil,
		DeleteHandler:  cLogout(),
		GetHandler:     cCheck(),
		HeadHandler:    nil,
		OptionsHandler: nil,
		PatchHandler:   nil,
		PostHandler:    cLogin(),
		PutHandler:     cUpdate(),
		TraceHandler:   nil,
	})

}
