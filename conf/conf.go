package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type GlobalHttpsConfig struct {
	Enabled bool   `json:"Enabled"`
	PemFile string `json:"PemFile"`
	KeyFile string `json:"KeyFile"`
}

type GlobalWebConfig struct {
	BindAddr        string            `json:"BindAddr"`
	BindPort        int               `json:"BindPort"`
	ServerName      string            `json:"ServerName"`
	HTTPS           GlobalHttpsConfig `json:"HTTPS"`
	TokenExpiredMin int64             `json:"TokenExpiredMin"`
	TokenSecretSalt string            `json:"TokenSecretSalt"`
}

type GlobalDbConfig struct {
	DSN         string `json:"DSN"`
	MaxOpenConn int    `json:"MaxOpenConn"`
	MaxIdleConn int    `json:"MaxIdleConn"`
}

type GlobalCacheConfig struct {
	Network     string `json:"Network"`
	RedisServer string `json:"RedisServer"`
	Auth        string `json:"Auth"`
	SelectNo    int    `json:"SelectNo"`
	MaxActive   int    `json:"MaxActive"`
	MaxIdle     int    `json:"MaxIdle"`
	CacheLifeMs int    `json:"CacheLifeMs"`
}

type GlobalConfig struct {
	Web   GlobalWebConfig   `json:"Web"`
	Db    GlobalDbConfig    `json:"Db"`
	Cache GlobalCacheConfig `json:"Cache"`
}

var (
	file = "config.json"
	g    GlobalConfig
)

func GetGlobalConfig() *GlobalConfig {
	return &g
}

func json2GlobalConfig(b []byte, gc *GlobalConfig) (string, bool) {
	err := json.Unmarshal(b, &gc)
	if err != nil {
		return "Unmarshal json to global config object failed", false
	}
	return "Successfully read global config", true
}

func init() {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panicln("Cannot read configuration file '" + file + "'")
		return
	}

	msg, ok := json2GlobalConfig(b, &g)

	if !ok {
		log.Panicln(msg)
	} else {
		log.Println(msg)
	}
}
