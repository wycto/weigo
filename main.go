package main //定义包名

import "wycto/frame" //引入包frame，wycto是在go.mod定义的模块名称，相当于项目名称go-frame项目文件夹层

//主函数入口，启动，main包下的main函数是可以执行的
func main() {
	//调用框架启动函数
	frame.Run()
}
