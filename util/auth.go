package util

import (
	"errors"
	"jun/model/user"
)

func Login(username string, password string) (bool, string, error) {
	// 先通过 Model 层验证用户名密码
	ok, id, err := user.AuthGetUid(&username, &password)
	if err != nil {
		return false, "", err
	} else if !ok {
		return false, "", errors.New("用户验证错误")
	}

	// 验证成功，生成 JWT
	token, err := NewJwtTokenByUid(id, username, password)
	if err != nil {
		return false, "", err
	}

	return true, token, nil
}

func Logout(token string) {
	// 直接加入黑名单
	BanToken(token)
}

func Check(token string) (*WebClaims, error) {
	// 验证 JWT 是否在 blacklist 中（是否已注销）
	if IsTokenBanned(token) {
		return nil, errors.New("授权凭据已注销，请重新登录")
	}

	// 解密并且验证 JWT 是否过期
	var claim *WebClaims
	jwt, err := ParseJwtToken(token)
	if err == nil && jwt != nil {
		if !jwt.Valid {
			return nil, errors.New("授权凭据已过期，请重新登录")
		} else if c, ok := jwt.Claims.(*WebClaims); ok {
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
