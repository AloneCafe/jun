package conf

import (
	"jun/utils/binary"
	"os"
	"strings"
	"sync"
)

var (
	m        sync.Mutex
	ipBL     = make(map[string]bool)
	ipBLFile = GetGlobalConfig().OtherConfig.IPBLFile
)

func BanIP(ip string) {
	m.Lock()
	defer m.Unlock()
	ipBL[ip] = true
}

func AllowIP(ip string) {
	m.Lock()
	defer m.Unlock()
	ipBL[ip] = false
}

func IsIPBanned(ip string) bool {
	return ipBL[ip]
}

func GetBL() (string, error) {
	m.Lock()
	defer m.Unlock()
	err := sync2BLFile()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(ipBLFile)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func SetBL(ipbl string) error {
	m.Lock()
	defer m.Unlock()
	err := os.WriteFile(ipBLFile, []byte(ipbl), 0666)
	if err != nil {
		return err
	}

	err = loadFromBLFile()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	err := loadFromBLFile()
	if err != nil {
		panic(err)
	}
}

func loadFromBLFile() error {
	data, err := os.ReadFile(ipBLFile)
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

func sync2BLFile() error {
	var buf []byte
	for k, v := range ipBL {
		if v {
			buf = binary.BytesMerge(buf, []byte(k+"\n"))
		}
	}
	err := os.WriteFile(ipBLFile, buf, 0666)
	if err != nil {
		return err
	}

	return nil
}
