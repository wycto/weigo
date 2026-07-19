package helper

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// MD5 生成 MD5 摘要（适用于缓存键、校验和，不适用于密码加密）
func MD5(s string) string {
	m := md5.Sum([]byte(s))
	return hex.EncodeToString(m[:])
}

// SHA256 生成 SHA-256 摘要
func SHA256(s string) string {
	m := sha256.Sum256([]byte(s))
	return hex.EncodeToString(m[:])
}
