package weigo

/*
日志类，日志的处理
*/
import "time"

type log struct {
}

func (receiver *log) FormatLogString(s string, t string, n string) string {
	typeInfo := "[" + t + " " + n + "]"
	string := "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + typeInfo + ":" + s
	return string
}
