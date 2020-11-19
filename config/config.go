package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DataBase DataBase
}

type DataBase struct {
	Type     string
	HostName string
	Port     string
	UserName string
	Password string
	Database string
	Charset  string
	Prefix   string
}

func (receiver *Config) Get() *Config {

	filePtr, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return nil
	}
	defer filePtr.Close()

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&receiver)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())
		return nil
	}
	fmt.Println("receiver: ", receiver)
	return receiver
}
