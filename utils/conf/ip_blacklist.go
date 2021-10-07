package conf

import (
	"os"
	"strings"
	"sync"

	"jun/utils/binary"
)

var (
	mIP  sync.Mutex
	ipBL = make(map[string]bool)
)

func BanIP(ip string) {
	mIP.Lock()
	defer mIP.Unlock()
	ipBL[ip] = true
}

func AllowIP(ip string) {
	mIP.Lock()
	defer mIP.Unlock()
	ipBL[ip] = false
}

func IsIPBanned(ip string) bool {
	return ipBL[ip]
}

func GetIPBL() (string, error) {
	mIP.Lock()
	defer mIP.Unlock()
	err := sync2IPBLFile()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(GetGlobalConfig().OtherConfig.IPBLFile)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func SetIPBL(ipbl string) error {
	mIP.Lock()
	defer mIP.Unlock()
	err := os.WriteFile(GetGlobalConfig().OtherConfig.IPBLFile, []byte(ipbl), 0666)
	if err != nil {
		return err
	}

	err = loadFromIPBLFile()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	err := loadFromIPBLFile()
	if err != nil {
		panic(err)
	}
}

func loadFromIPBLFile() error {
	data, err := os.ReadFile(GetGlobalConfig().OtherConfig.IPBLFile)
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(data), "\n") {
		ip := strings.TrimSpace(line)
		if len(ip) > 0 && ip[0] != '#' {
			ipBL[ip] = true
		}
	}

	return nil
}

func sync2IPBLFile() error {
	var buf []byte
	for k, v := range ipBL {
		if v {
			buf = binary.BytesMerge(buf, []byte(k+"\n"))
		}
	}
	err := os.WriteFile(GetGlobalConfig().OtherConfig.IPBLFile, buf, 0666)
	if err != nil {
		return err
	}

	return nil
}
