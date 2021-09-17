package util

import (
	"github.com/dgrijalva/jwt-go"
	"jun/conf"
	"strconv"
	"time"
)

type WebClaims struct {
	jwt.StandardClaims
	UID   int64  `json:"u_id"`
	Uname string `json:"u_uname"`
	Pwd   string `json:"u_pwd"`
}

func NewJwtTokenByUid(id int64, uname string, pwd string) (string, error) {
	now := time.Now().Unix()
	idStr := strconv.FormatInt(id, 10)
	claims := WebClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  uname,
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(conf.GetGlobalConfig().Web.TokenExpiredMin)).Unix(),
			Id:        idStr,
			IssuedAt:  now,
			Issuer:    "Alone Cafe",
			NotBefore: now,
			Subject:   "login with user authorization",
		},
		UID:   id,
		Uname: uname,
		Pwd:   pwd,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(conf.GetGlobalConfig().Web.TokenSecretSalt)
	token, err := tokenClaims.SignedString(secret)
	return token, err
}

func ParseJwtToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &WebClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		secret := []byte(conf.GetGlobalConfig().Web.TokenSecretSalt)
		return secret, nil
	})
}
