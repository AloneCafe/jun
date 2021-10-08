package tests

import (
	"encoding/json"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"jun/dao"
	"jun/model/option"
	"jun/model/post"
	"jun/utils/conf"
)

func TestGlobalConfig(t *testing.T) {
	t.Log(*conf.GetGlobalConfig())
}

func TestDbConn(t *testing.T) {
	func() {
		database, err := sqlx.Open("mysql", conf.GetGlobalConfig().Db.DSN)
		if err != nil {
			t.Error("Failed to initialize MySQL driver, ", err)
			return
		} else {
			t.Log("Database driver available")
		}
		Db := database
		err = Db.Ping()
		if err != nil {
			t.Error("Failed to connect MySQL server, ", err)
			return
		} else {
			t.Log("Database connection available")
		}
	}()
}

func TestRedisConn(t *testing.T) {
	func() {
		const siz = 1
		var args1, args2 [siz]string
		for i := 0; i < siz; i++ {
			arg1, arg2 := "testK"+strconv.Itoa(i), "testV"+strconv.Itoa(i)
			args1[i] = arg1
			args2[i] = arg2
			if _, err := dao.TestRedis(arg1, arg2, t); err != nil {
				t.Error("Something failed on Redis connect/set,", err)

			}
		}

		for i := 0; i < siz; i++ {
			if _, err := dao.TestRedis(args1[i], args2[i], t); err != nil {
				t.Error("Something failed on Redis connect/get,", err)
			}
		}

	}()
}

func TestFindPost(t *testing.T) {
	func() {
		pp, err := post.FindPost("7", ".*", ".*", 1, 0)
		if err != nil {
			t.Error(err)
		} else {
			s, err := json.Marshal(pp)
			if err != nil {
				t.Error(err)
			}
			t.Log(string(s))
		}
	}()
}

func TestGetOption(t *testing.T) {
	func() {
		val, err := option.GetPostCountPerPage()
		if err != nil {
			t.Error(err)
		} else {
			t.Log(val)
		}
	}()
}

/*
func TestGetUser(t *testing.T) {
	func() {
		const uid = 0
		p, err := user.GetById(uid)
		if err != nil {
			t.Error("Failed with uid =", uid, err)
		} else {
			if b, err := json.Marshal(*p); err != nil {
				t.Log("Get user failed:", err)
			} else {
				t.Log("User object:", string(b))
			}
		}
	}()
}

func TestGetAllUser(t *testing.T) {
	func() {
		pCnt, err := user.CountAll()
		pUsers, err := user.GetAll()
		if err != nil {
			t.Error("Failed,", err)
		} else {
			if b, err := json.Marshal(*pUsers); err != nil {
				t.Log("Get users failed:", err)
			} else {
				t.Log("Total", *pCnt, "Users object:", string(b))
			}
		}
		uname := "admin"
		pwd := "admin"
		pFlag, err := user.Auth(&uname, &pwd)
		if err != nil {
			t.Error("Auth failed,", err)
		} else {
			if *pFlag {
				t.Log("Auth successful")
			} else {
				t.Error("Wrong username or password")
			}
		}
	}()
}

func TestGetAllUser2(t *testing.T) {
	func() {
		pCnt, err := user.CountAll()
		pUsers, err := user.GetAll()
		if err != nil {
			t.Error("Failed,", err)
		} else {
			if b, err := json.Marshal(*pUsers); err != nil {
				t.Log("Get users failed:", err)
			} else {
				t.Log("Total", *pCnt, "Users object:", string(b))
			}
		}
		uname := "admin"
		pwd := "admin"
		pFlag, err := user.Auth(&uname, &pwd)
		if err != nil {
			t.Error("Auth failed,", err)
		} else {
			if *pFlag {
				t.Log("Auth successful")
			} else {
				t.Error("Wrong username or password")
			}
		}
	}()
}

func TestInsert(t *testing.T) {
	func() {
		email := "admin@example.com"
		uname := "admin"
		pwd := "admin"
		insId, err := user.Add(&email, &uname, &pwd, nil, nil, 0, nil, nil, 0)
		if err != nil {
			t.Error("Add user failed", err)
		} else {
			t.Log("Add user successful, last insert id:", insId)
		}
	}()
}

func TestTokenGen(t *testing.T) {
	func() {
		token, err := jwt.NewJwtTokenByUid(1, "admin", "admin", "127.0.0.1")
		if err != nil {
			t.Error(err)
		}

		jwtToken, err := jwt.ParseJwtToken(token)
		if err != nil {
			t.Error("Token generation test failed,", err)
			return
		} else if !jwtToken.Valid {
			t.Error("Token generation test failed,", "token is invalid")
			return
		}

	}()
}
*/
