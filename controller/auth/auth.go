package auth

import (
	"github.com/gin-gonic/gin"
	"jun/controller/base"
	"jun/dto"
	"jun/util"
	"log"
	"net/http"
)

var (
	cLogout = func() gin.HandlerFunc {
		return func(c *gin.Context) {
			token, err := base.GetBearerToken(c)
			if err != nil {
				return

			} else {
				util.BanToken(*token)
				c.JSON(http.StatusOK,
					dto.NewResult(true, "授权已注销", nil))
			}
		}
	}

	cCheck = func() gin.HandlerFunc {
		return func(c *gin.Context) {
			token, err := base.GetBearerToken(c)
			if err != nil {
				return

			} else {
				_, err := base.Check(*token)
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

	cLogin = func() gin.HandlerFunc {
		return func(c *gin.Context) {
			//username := c.PostForm("u_uname")
			//password := c.PostForm("u_pwd")
			json := make(map[string]string)
			err := c.BindJSON(&json)
			if err != nil {
				c.JSON(http.StatusBadRequest,
					dto.NewResult(false, "错误的请求格式", nil))
				return
			}
			username := json["u_uname"]
			password := json["u_pwd"]
			log.Println(username, password)

			ok, token, err := base.Login(username, password)
			if !ok || err != nil {
				c.JSON(http.StatusUnauthorized,
					dto.NewResult(false, "授权失败，用户名或者密码错误", map[string]string{
						"token": token,
					}))
			} else {
				c.JSON(http.StatusOK,
					dto.NewResult(true, "授权成功", map[string]string{
						"token": token,
					}))
			}
		}
	}

	cUpdate = func() gin.HandlerFunc {
		return func(c *gin.Context) {
			token, err := base.GetBearerToken(c)
			if err != nil {
				return

			} else {
				claims, err := base.Check(*token)
				if err != nil {
					c.JSON(http.StatusOK,
						dto.NewResult(false, "授权凭据过期或无效", nil))

				} else {
					// 尝试再次登录（用户名密码）
					ok, token, err := base.Login(claims.Uname, claims.Pwd)
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

type AuthController struct {
	base.IBasicController
}

func (*AuthController) DeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := base.GetBearerToken(c)
		if err != nil {
			return

		} else {
			util.BanToken(*token)
			c.JSON(http.StatusOK,
				dto.NewResult(true, "授权已注销", nil))
		}
	}
}

func (*AuthController) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := base.GetBearerToken(c)
		if err != nil {
			return

		} else {
			_, err := base.Check(*token)
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

func (*AuthController) PostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//username := c.PostForm("u_uname")
		//password := c.PostForm("u_pwd")
		json := make(map[string]string)
		err := c.BindJSON(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				dto.NewResult(false, "错误的请求格式", nil))
			return
		}
		username := json["u_uname"]
		password := json["u_pwd"]
		log.Println(username, password)

		ok, token, err := base.Login(username, password)
		if !ok || err != nil {
			c.JSON(http.StatusUnauthorized,
				dto.NewResult(false, "授权失败，用户名或者密码错误", map[string]string{
					"token": token,
				}))
		} else {
			c.JSON(http.StatusOK,
				dto.NewResult(true, "授权成功", map[string]string{
					"token": token,
				}))
		}
	}
}

func (*AuthController) PutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := base.GetBearerToken(c)
		if err != nil {
			return

		} else {
			claims, err := base.Check(*token)
			if err != nil {
				c.JSON(http.StatusOK,
					dto.NewResult(false, "授权凭据过期或无效", nil))

			} else {
				// 尝试再次登录（用户名密码）
				ok, token, err := base.Login(claims.Uname, claims.Pwd)
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
