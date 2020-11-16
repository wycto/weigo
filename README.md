# WeiGo 
#### _1.0.0版本（无数据库操作）_
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
 待续
 