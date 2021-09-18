package binary

import (
	"bytes"
	"encoding/gob"
)

// parameter no need pointer
func SerializeValue(i interface{}) []byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(i)
	if err != nil {
		return nil
	}
	return buff.Bytes()
}

// parameter need pointer
func DeserializeValue(b []byte, i interface{}) bool {
	buff := bytes.NewBuffer(b)
	enc := gob.NewDecoder(buff)
	err := enc.Decode(i)
	if err != nil {
		return false
	}
	return true
}

func BytesMerge(ss ...[]byte) []byte {
	var totalLen int
	for _, s := range ss {
		totalLen += len(s)
	}
	tmp := make([]byte, totalLen)
	var i int
	for _, s := range ss {
		i += copy(tmp[i:], s)
	}
	return tmp
}
