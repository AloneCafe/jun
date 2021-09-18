package base

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jun/dto"
	"jun/model/user"
	"jun/utils/jwt"
	"net/http"
	"strings"
)

func Login(username string, password string, ipaddr string) (bool, string, error) {
	// 先通过 Model 层验证用户名密码
	ok, id, err := user.AuthGetUid(&username, &password)
	if err != nil {
		return false, "", err
	} else if !ok {
		return false, "", errors.New("用户验证错误")
	}

	if banned, err := user.IsUserBannedById(id); err != nil {
		return false, "", errors.New("用户验证错误")
	} else if banned {
		return false, "", errors.New("此用户已被禁止登录")
	}

	// 验证成功，生成 JWT
	token, err := jwt.NewJwtTokenByUid(id, username, password, ipaddr)
	if err != nil {
		return false, "", err
	}

	return true, token, nil
}

func Logout(token string) {
	// 直接加入黑名单
	jwt.BanToken(token)
}

func Check(token string) (*jwt.WebClaims, error) {
	// 验证 JWT 是否在 blacklist 中（是否已注销）
	if jwt.IsTokenBanned(token) {
		return nil, errors.New("授权凭据已注销，请重新登录")
	}

	// 解密并且验证 JWT 是否过期
	var claim *jwt.WebClaims
	tk, err := jwt.ParseJwtToken(token)
	if err == nil && tk != nil {
		if !tk.Valid {
			return nil, errors.New("授权凭据已过期，请重新登录")
		} else if c, ok := tk.Claims.(*jwt.WebClaims); ok {
			claim = c
		}
	} else {
		return nil, err
	}

	// 验证 JWT 中附带的用户名密码
	ok, id, err := user.AuthGetUid(&claim.Uname, &claim.Pwd)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New("授权凭据已更改，请重新登录")
	}

	// 验证 JWT 中附带的 ID
	if id != claim.UID {
		return nil, errors.New("授权凭据已更改，请重新登录")
	}

	return claim, nil
}

func GetBearerToken(c *gin.Context) (token *string, e error) {
	token = nil
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		e = errors.New("请求头授权字段为空")
		c.JSON(http.StatusUnauthorized,
			dto.NewResult(false, "请求头授权字段为空", nil))
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		e = errors.New("请求头授权字段格式错误")
		c.JSON(http.StatusBadRequest,
			dto.NewResult(false, "请求头授权字段格式错误", nil))
		return
	}
	token = &parts[1]
	return
}

func Authorization(c *gin.Context, lowRole dto.UserRole) (wc *jwt.WebClaims, e error) {
	wc = nil
	e = nil
	if lowRole <= dto.U_ROLE_VISITOR {
		return
	}

	token, err := GetBearerToken(c)
	if err != nil {
		return

	} else if claims, err := Check(*token); err != nil {
		e = err
		c.JSON(http.StatusUnauthorized,
			dto.NewResult(false, "操作被拦截，用户凭据过期或者授权已经注销，请重新授权", nil))

	} else if banned, err := user.IsUserBannedById(claims.UID); err != nil {
		e = errors.New("授权在检测用户是否被封禁时发生错误")
		c.JSON(http.StatusUnauthorized,
			dto.NewResult(false, "操作被拦截，用户相关的信息可能已经修改，请重新登录", nil))

	} else if banned {
		e = errors.New("用户已被封禁，请与管理员取得联系")
		c.JSON(http.StatusUnauthorized,
			dto.NewResult(false, "操作被拦截，用户已被封禁", nil))

	} else if role, err := user.GetRoleById(claims.UID); err != nil || role == nil {
		e = errors.New("授权在获取 UID 时发生错误")
		c.JSON(http.StatusUnauthorized,
			dto.NewResult(false, "操作被拦截，用户相关的信息可能已经修改，请重新登录", nil))

	} else if *role < lowRole {
		e = errors.New("请求头授权字段权限不足")
		c.JSON(http.StatusUnauthorized,
			dto.NewResult(false, "操作被拦截，执行该操作需要更高的用户权限", nil))

	} else {
		wc = claims
	}
	return
}
