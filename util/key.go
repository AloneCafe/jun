package util

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/gob"
)

func Key512(args ...interface{}) [64]byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(args)
	if err != nil {
		return [64]byte{}
	}
	return sha512.Sum512(buff.Bytes())
}

func Key1(args ...interface{}) [20]byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(args)
	if err != nil {
		return [20]byte{}
	}
	return sha1.Sum(buff.Bytes())
}
