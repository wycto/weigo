package weigo

import "os"

/*
系统应用程序全局向光，定义常量、全局变量等
*/

//定义全局变量-常量
var (
	RootPath, _ = os.Getwd()               //根目录
	DS          = string(os.PathSeparator) //目录分隔符
)

//定义全局变量-系统类
var (
	Config = &config{}   //配置
	DB     = &dataBase{} //数据库
	Log    = &log{}      //日志
)

//APP类
type App struct {
}

//初始化方法
func init() {
	configInit()   //初始化配置
	dataBaseInit() //初始化数据库
}

//初始化配置
func configInit() {
	if Config.initStatus == false {
		Config.loadConfig() //加载配置
	}
}

//初始化数据库
func dataBaseInit() {
	if DB.initStatus == false {
		DB.getConnect() //连接数据库
	}
}
