package dao

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/gob"
	"github.com/gomodule/redigo/redis"
	"jun/conf"
	"jun/dao/cache"
	"jun/dao/db"
	"log"
	"testing"
)


func key512(args ...interface{}) [64]byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(args)
	if err != nil {
		return [64]byte{}
	}
	return sha512.Sum512(buff.Bytes())
}


func key1(args ...interface{}) [20]byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(args)
	if err != nil {
		return [20]byte{}
	}
	return sha1.Sum(buff.Bytes())
}

var (
	genKey = key1
)

// parameter no need pointer
func serializeValue(i interface{}) []byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(i)
	if err != nil {
		return nil
	}
	return buff.Bytes()
}

// parameter need pointer
func deserializeValue(b []byte, i interface{}) bool {
	buff := bytes.NewBuffer(b)
	enc := gob.NewDecoder(buff)
	err := enc.Decode(i)
	if err != nil {
		return false
	}
	return true
}

func TestRedis(arg1 string, arg2 string, t *testing.T) (interface{}, error) {
	type KV struct {
		K string
		V string
	}
	p := KV{ arg1, arg2 }
	myRds := cache.GetCache()
	defer myRds.Close()

	key := genKey(serializeValue(p))

	exists, err := redis.Bool(myRds.Do("EXISTS", key))
	if err == nil {
		Store := func() error {
			_, err := myRds.Do("SET", key, serializeValue(p), "PX", conf.GetGlobalConfig().Cache.CacheLifeMs)
			t.Log("Set sample data,", "hashKey:", key, "value:", p)
			return err
		}
		if exists { // get from cache
			res, err := redis.Bytes(myRds.Do("GET", key))
			if err != nil { // deserialization failed
				return nil, err
			} else if !deserializeValue(res, &p) {
				t.Error("Deserializing failed,", "hashKey:", key)
			} else {
				t.Log("Get sample data:", "hashKey:", key, "value:", p)
				return p, err
			}
		} else { // store
			return nil, Store()
		}

	}
	return nil, err
}

// Query1 p points to a struct
func Query1(p interface{}, sql string, args ...interface{}) error {
	//m, _ := json.Marshal(args)
	//sqlargs := sql + " @<" + string(m) + "> "
	//log.Printf("Query1: %s", sqlargs)
	log.Printf("Query1: %s", sql)

	myRds := cache.GetCache()
	defer myRds.Close()

	myDb := db.GetDB()
	//defer myDb.Close()

	sqlargs := append(args, sql)
	key := genKey(serializeValue(sqlargs))

	exists, err := redis.Bool(myRds.Do("EXISTS", key))
	if err == nil {
		getFromDB := func() error {
			if err := myDb.Get(p, sql, args...); err != nil {
				return err
			}
			_, err := myRds.Do("SET", key, serializeValue(p), "PX", conf.GetGlobalConfig().Cache.CacheLifeMs)
			log.Println("No cache found, get from DB...")
			return err
		}
		if exists { // get from cache
			res, err := redis.Bytes(myRds.Do("GET", key))
			if err != nil { // deserialization failed, get from DB
				return getFromDB()
			} else if !deserializeValue(res, p) {
				log.Println("Deserializing failed,", "hashKey:", key)
				return err
			} else {
				log.Println("Cache founded, get from Cache...")
				return err
			}
		} else { // get from DB, then cache it
			return getFromDB()
		}

	}
	log.Println()
	return err
}

// QueryN p points to a struct array
func QueryN(pp interface{}, sql string, args ...interface{}) error {
	//m, _ := json.Marshal(args)
	//sqlargs := sql + "@<" + string(m) + "> "
	//log.Printf("QueryN: %s", sqlargs)
	log.Printf("QueryN: %s", sql)

	myRds := cache.GetCache()
	defer myRds.Close()

	myDb := db.GetDB()
	//defer myDb.Close()

	sqlargs := append(args, sql)
	key := genKey(serializeValue(sqlargs))

	exists, err := redis.Bool(myRds.Do("EXISTS", key))
	if err == nil {
		getFromDB := func() error {
			if err := myDb.Select(pp, sql, args...); err != nil {
				return err
			}
			_, err := myRds.Do("SET", key, serializeValue(pp), "PX", conf.GetGlobalConfig().Cache.CacheLifeMs)
			log.Println("No cache found, get from DB...")
			return err
		}
		if exists { // get from cache
			res, err := redis.Bytes(myRds.Do("GET", key))
			if err != nil { // deserialization failed, get from DB
				return getFromDB()
			} else if !deserializeValue(res, pp) {
				log.Println("Deserializing failed,", "hashKey:", key)
				return err
			} else {
				log.Println("Cache founded, get from Cache...")
				return err
			}
		} else { // get from DB, then cache it
			return getFromDB()
		}
	}
	log.Println()
	return err
}

func Insert(sql string, args ...interface{}) (int64, error) {
	//m, _ := json.Marshal(args)
	//log.Println("Insert:", sql, string(m))
	log.Printf("Insert: %s", sql)

	myDb := db.GetDB()
	//defer myDb.Close()

	result, err := myDb.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	insId, err := result.LastInsertId()
	if err != nil {
		return insId, err
	}
	log.Printf("Insert succeed, last id: %d\n", insId)
	return insId, nil
}

func Update(sql string, args ...interface{}) (int64, error) {
	//m, _ := json.Marshal(args)
	//log.Println("Update:", sql, string(m))
	log.Printf("Update: %s", sql)

	myDb := db.GetDB()
	//defer myDb.Close()

	result, err := myDb.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return rows, err
	}
	log.Printf("Update succeed, affected rows: %d\n", rows)
	return rows, nil
}

func Delete(sql string, args ...interface{}) (int64, error) {
	//m, _ := json.Marshal(args)
	//log.Println("Delete:", sql, string(m))
	log.Printf("Delete: %s", sql)

	myDb := db.GetDB()
	//defer myDb.Close()

	result, err := myDb.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return rows, err
	}
	log.Printf("Delete succeed, affected rows: %d\n", rows)
	return rows, nil
}

