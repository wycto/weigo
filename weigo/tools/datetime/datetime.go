package datetime

import "time"

func DateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Date() string {
	return time.Now().Format("2006-01-02")
}

func Time() string {
	return time.Now().Format("15:04:05")
}
