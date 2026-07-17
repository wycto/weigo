package helper

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 加密
func MD5(s string) string {
	m := md5.Sum([]byte(s))
	return hex.EncodeToString(m[:])
}
