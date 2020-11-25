package weigo

//定义全局变量
var (
	Config   = &config{}   //配置
	DataBase = &dataBase{} //数据库
)

type App struct {
}

//初始化方法，初始化其他数据
func init() {
	configInit()   //加载配置
	dataBaseInit() //连接数据库
}

//初始化配置
func configInit() {
	if Config.initStatus == false {
		Config.loadConfig() //加载配置
	}
}

//初始化数据库
func dataBaseInit() {
	if DataBase.initStatus == false {
		DataBase.getConnect() //连接数据库
	}
}
