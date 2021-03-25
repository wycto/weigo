package log

/*
日志类，日志的处理
*/
import (
	"github.com/wycto/weigo"
	"time"
)

type Log struct {
}

func (l *Log) FormatLogString(s string, t string, n string) string {
	typeInfo := "[" + t + " " + n + "]"
	string := "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + typeInfo + ":" + s
	return string
}

//写日志文
func Write(str string) {
	weigo.Log.Write(str)
}
