package weigo

import (
	"crypto/md5"
	"encoding/hex"
)

/*
助手函数都写在这儿
*/

//MD5加密
func MD5(s string) string {
	m := md5.Sum([]byte(s))
	return hex.EncodeToString(m[:])
}
