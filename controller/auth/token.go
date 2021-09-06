package auth

var (
	blacklist map[string]int
)

func banToken(token string) {
	blacklist[token] = 1
}

func allowToken(token string) {
	delete(blacklist, token)
}

func isTokenBanned(token string) bool {
	return blacklist[token] == 1
}

//func verifyToken(token string) (*util.WebClaims, error) {
//	jwtToken, err := util.ParseJwtToken(token)
//	if err == nil && jwtToken != nil {
//		if !jwtToken.Valid {
//			return nil, errors.New("凭据已过期，请重新登录")
//		} else if claim, ok := jwtToken.Claims.(*util.WebClaims); ok {
//			return claim, nil
//		}
//	}
//	return nil, err
//}
//
//func getUidByToken(token string) (int64, error) {
//	claims, err := verifyToken(token)
//	if err != nil {
//		return 0, err
//	} else {
//		return strconv.ParseInt(claims.Id, 10, 64)
//	}
//}
