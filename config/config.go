package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// RootPath 全局根目录
var RootPath, _ = os.Getwd()

// DS 目录分隔符
var DS = string(os.PathSeparator)

// Config 主配置
type Config struct {
	App  AppConfig  `json:"App"`
	DB   DBConfig   `json:"DB"`
	Log  LogConfig  `json:"Log"`
	View ViewConfig `json:"View"`
}

// AppConfig 应用配置
type AppConfig struct {
	Debug                 bool   `json:"Debug"`
	Port                  string `json:"Port"`
	MultiApp              bool   `json:"MultiApp"`
	DefaultAppName        string `json:"DefaultAppName"`
	DefaultControllerName string `json:"DefaultControllerName"`
	DefaultActionName     string `json:"DefaultActionName"`
	UseDB                 bool   `json:"UseDB"`
}

// DBConfig 数据库配置
type DBConfig struct {
	Driver   string `json:"Driver"`
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Database string `json:"Database"`
	Charset  string `json:"Charset"`
	Prefix   string `json:"Prefix"`
}

// ViewConfig 视图配置
type ViewConfig struct {
	RootPath string `json:"RootPath"`
}

// LogConfig 日志配置
type LogConfig struct {
	Console bool   `json:"Console"`
	File    bool   `json:"File"`
	Level   string `json:"Level"` // debug, info, warn, error
}

// GlobalConfig 全局配置实例
var GlobalConfig *Config

// Load 加载配置
func Load(configDir string) *Config {
	if configDir != "" {
		configDir = strings.Trim(configDir, DS) + DS
	}

	filePath := RootPath + DS + configDir + "config.json"
	file, err := os.Open(filePath)
	if err != nil {
		panic("配置文件打开失败: " + filePath + " " + err.Error())
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		panic("配置加载失败: " + err.Error())
	}

	// 设置默认值
	if cfg.App.DefaultControllerName == "" {
		cfg.App.DefaultControllerName = "index"
	}
	if cfg.App.DefaultActionName == "" {
		cfg.App.DefaultActionName = "index"
	}
	if cfg.App.DefaultAppName == "" {
		cfg.App.DefaultAppName = "index"
	}
	if cfg.View.RootPath == "" {
		cfg.View.RootPath = "view"
	}

	fmt.Println("配置加载完成")
	GlobalConfig = &cfg
	return &cfg
}
