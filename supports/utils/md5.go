package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
