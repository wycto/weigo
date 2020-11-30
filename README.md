# WeiGo 

V0.3.0版本（有数据库的增删改查）
-----------------
Go框架 - WeiGo【weigo】

|project #项目目录

|----app   #应用目录

|----config    #配置目录

|----route    #路由目录

|----weigo   #框架目录

|----go.mod    #项目go.mod文件

|----main.go    #项目入口启动文件

# **特点：**
MVC模式，M-模型（数据层）, V-视图（页面层）, C-控制器（请求响应）


# 使用：

### 配置

全部配置在一个json文件，json格式；

App：是项目配置

Log：日志配置

DB：数据库配置

`{

   "App": {
   
     "Debug": true
     
   },
   
   "Log": {
   
     "SqlInfo": "console"
     
   },
   
   "DB": {
   
     "Type": "mysql",
     
     "Hostname": "127.0.0.1",
     
     "Port": "3306",
     
     "Database": "weigo",
     
     "Charset": "utf8mb4",
     
     "Username": "root",
     
     "Password": "root",
     
     "Prefix": "wei_"
     
   }
 }`
 
 
### 路由：
在route目录下route.go里面注册好路由：

`
weigo.Router("/", &controller.IndexController{})
 `
 
 `weigo.Router("/user/", &controller.UserController{})`
 
 
_第一个参数：_
 访问路劲，
 
 user是控制器名称，对应app/controller/UserController.go文件里面的UserController结构体；
 
 该结构体会有很多方法，方法名称不需要定义路由，自动匹配：
 比如127.0.0.1:9099/user/userinfo 会访问UserController的UserInfo方法
 路由全小写匹配
 
 
 _**第二个参数：**_
 控制器类型的变量
 
 ### 控制器：
 控制器需要继承框架控制器：
 
 `type UserController struct {
  	weigo.Controller
  }`
 
 ### 视图：
 在控制器方法里面，也就是结构体方法里面：
 
 `c.Display("")`
 
 参数空，则使用app/view/控制器名称/方法名称.html文件，不为空就使用参数模板
 
 赋值变量：
 
 `c.Assign("name","唯一")`
 
 c是当前控制器
 
 模板文件html使用：
 
 `<p>{{.name}}</p>`
 
 
 ### 模型：


user模型：
路径：`app/model/user.go`
代码规则建议:
一个模型对应一个数据表，属于一个模块；

一个模型只能有一个结构体，就是此模型的结构体，命名一致

模型里面定义的都是改结构体的方法，无其他函数

模型方法如果要获取数据库数据，采用db获取：

`rows, errorStr := weigo.DataBase.Table("cto_controller").SetFields("name,id").GetAll()`

###数据库操作，支持链式操作

`rows, errorStr := weigo.DataBase.Table("cto_controller").SetFields("name,id").GetAll()`

链式操作之前，先获取数据库连接：
`weigo.DataBase`

链式操作：先指定操作的数据表，两种方式：
`weigo.DataBase.Table("wei_user")`

`weigo.DataBase.Name("user")`

Table：手动写全表名称

Name：不加前缀，前缀使用配置里面配置的前缀

SetFields：设置要查询的字段

Where：查询条件；可以传字符串；也可以传map

Where("name=\"weigo\" and status=1")

Where(["name"=>"[:string]weigo","status"=>1]);  //[:string]代表字符串，会加引号

Where(["name[like]"=>"weigo","status"=>1])；//[like]代表模糊查询

Group("age,status")：分组，传字符串

Having：传字符串

Order("uid desc,name asc")：排序，传字符串

Limit(10)：GetAll有效，GetOne无效

Page(2,10)：分页查询，第一个参数页面，第二个参数每页大小

GetOne:获取一条数据

GetAll：获取多条数据

连式操作：Table或者Name指定表，GetOne或者GetAll获取数据，其他的不分顺序

`weigo.DataBase.Table("cto_controller").SetFields("name,id").Where().Group().Order().GetAll()`


## 数据库增删查改：

```go
package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
	"wycto/weigo"
)

type APIController struct {
	weigo.Controller
}

/**
查询
*/
func (c *APIController) Index() {
	var ww map[string]interface{}
	ww = make(map[string]interface{})

	rows, errorStr := weigo.DataBase.Name("user").SetFields("email,`name`,`nickname`").Where(ww).GetAll()
	if errorStr != "" {
		fmt.Println(errorStr)
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

/**
删除全部
*/
func (c *APIController) DeleteAll() {
	var ww map[string]interface{}
	ww = make(map[string]interface{})
	ww["uid|<"] = 5
	ww["nickname"] = "唯一"

	rows, errorStr := weigo.DataBase.Name("user").DeleteAll()
	if errorStr != "" {
		fmt.Println(errorStr)
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

/**
删除
*/
func (c *APIController) Delete() {
	var ww map[string]interface{}
	ww = make(map[string]interface{})
	ww["nickname"] = "update后的唯一"

	rows, errorStr := weigo.DataBase.Name("user").Where(ww).Delete()
	if errorStr != "" {
		fmt.Println(errorStr)
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

/**
更新
*/
func (c *APIController) Update() {
	var ww map[string]interface{}
	ww = make(map[string]interface{})
	ww["nickname"] = "唯一"

	var dd map[string]interface{}
	dd = make(map[string]interface{})
	dd["nickname"] = "update后的唯一"

	rows, errorStr := weigo.DataBase.Name("user").Where(ww).Update(dd)
	if errorStr != "" {
		fmt.Println(errorStr)
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

/**
新增
*/
func (c *APIController) Add() {
	var ww map[string]interface{}
	ww = make(map[string]interface{})
	ww["uid"] = 3

	var dd map[string]interface{}
	dd = make(map[string]interface{})
	dd["nickname"] = "唯一"

	rows, errorStr := weigo.DataBase.Name("user").Insert(dd)
	if errorStr != "" {
		fmt.Println(errorStr)
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

/**
更新全部
*/
func (c *APIController) UpdateAll() {
	var ww map[string]interface{}
	ww = make(map[string]interface{})
	ww["uid"] = 3

	var dd map[string]interface{}
	dd = make(map[string]interface{})
	dd["nickname"] = "唯一333"

	rows, errorStr := weigo.DataBase.Name("user").Where(ww).UpdateAll(dd)
	if errorStr != "" {
		fmt.Println(errorStr)
	}
	json, _ := json.Marshal(rows)
	io.WriteString(c.Context.ResponseWriter, string(json))
}

```

 