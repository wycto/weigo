package weigo

/*
日志类，日志的处理
*/
import (
	"fmt"
	"os"
	"time"
)

type log struct {
}

func (l *log) FormatLogString(s string, t string, n string) string {
	typeInfo := "[" + t + " " + n + "]"
	string := "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + typeInfo + ":" + s
	return string
}

//写日志文件
func (l *log) Write(str string) {
	os.Mkdir(RootPath+DS+"log", os.ModePerm)
	file, err := os.OpenFile(RootPath+DS+"log"+DS+"error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	str = "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + str + "\n\n"

	file.WriteString(str)
}
