package weigo

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	initStatus bool
	DB         dbConfig
	APP        appConfig
	Console    consoleConfig
}

type dbConfig struct {
	Type     string
	HostName string
	Port     string
	UserName string
	Password string
	Database string
	Charset  string
	Prefix   string
}

type appConfig struct {
}

type consoleConfig struct {
	InfoSqlLog bool
}

func (receiver *config) loadConfig() {

	filePtr, err := os.Open(RootPath + "/config/config.json")
	if err != nil {
		fmt.Println("Open file "+RootPath+"/config/config.json failed [Err:%s]", err.Error())
	}
	defer filePtr.Close()

	//创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&receiver)
	if err != nil {
		fmt.Println("config init failed", err.Error())
	} else {
		receiver.initStatus = true
		fmt.Println("config init ok")
	}
}
