package weigo

/*
配置处理，加载配置
*/
import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//配置
type config struct {
	initStatus bool      //是否已经初始化
	DB         dbConfig  //数据库配置
	App        appConfig //系统应用配置
	Log        logConfig //日志配置
	Sql        sqlConfig
	View       viewConfig
}

//数据库配置
type dbConfig struct {
	Type     string //数据库类型
	HostName string //数据库连接地址
	Port     string //端口
	UserName string //账号
	Password string //密码
	Database string //数据库名称
	Charset  string //字符集
	Prefix   string //表前缀
}

//系统应用配置
type appConfig struct {
	ServerPort            string
	Debug                 bool
	MultiApp              bool
	DefaultAppName        string
	DefaultControllerName string
	DefaultActionName     string
}

//日志配置
type logConfig struct {
	SqlInfo string
}

//sql语句记录
type sqlConfig struct {
	Console bool
	File    bool
}

//视图配置
type viewConfig struct {
	RootPath string
}

//加载配置
func (receiver *config) loadConfig(configDir string) {

	filePtr, err := os.Open(RootPath + DS + strings.Trim(configDir,DS) + DS + "config.json")
	if err != nil {
		fmt.Println("配置文件打开失败", err.Error())
		return ;
	}
	defer filePtr.Close()

	//创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&receiver)
	if err != nil {
		fmt.Println("配置加载失败", err.Error())
	} else {
		receiver.initStatus = true
		fmt.Println("配置加载完成")
	}
}
