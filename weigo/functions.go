package weigo

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
)

/**
获取IP地址
*/
func GetIP() string {
	var ip = ""
	address, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range address {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}

		}
	}

	return ip
}

func ParseWhereKeyAndRegexp(fieldReg string) (string, string) {
	reg, err := regexp.Compile(`\|(.*)`)
	if err != nil {
		fmt.Println("regexp err:", err.Error())
		return fieldReg, "="
	}

	result := reg.FindAllString(fieldReg, -1)
	if len(result) == 0 {
		return fieldReg, "="
	}

	fieldRegIndexArr := reg.FindAllStringIndex(fieldReg, -1)
	position := fieldRegIndexArr[0][0]
	field := fieldReg[:position]
	regexpStr := fieldReg[position+1:]
	regexpStr = strings.Trim(regexpStr, " ")
	if regexpStr == "" {
		regexpStr = "="
	}
	return field, regexpStr
}
