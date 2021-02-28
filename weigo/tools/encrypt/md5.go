package encrypt

import (
	"crypto/md5"
	"fmt"
)

//MD5加密
func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
