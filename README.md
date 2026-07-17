# WeiGo 2.0

基于 Echo + GORM 的 Go Web 框架，保留 PHP 风格的开发体验。

## 特性

- **MVC 架构**：控制器、模型、视图分离
- **Echo HTTP 框架**：高性能、易用
- **GORM ORM**：安全的数据库操作、链式查询
- **反射路由**：自动匹配控制器方法（PHP 风格）
- **模板渲染**：支持 HTML 模板
- **WebSocket**：实时通信支持
- **中间件**：日志、跨域、恢复
- **优雅关闭**：信号处理、平滑停机
- **连接池**：数据库连接管理

## 安装

```bash
go get github.com/wycto/weigo
```

## 快速开始

### 1. 创建项目

```bash
mkdir myproject && cd myproject
go mod init myproject
go get github.com/wycto/weigo
```

### 2. 配置文件

创建 `config/config.json`：

```json
{
    "App": {
        "Debug": true,
        "Port": "9099",
        "UseDB": true
    },
    "DB": {
        "Driver": "mysql",
        "Host": "127.0.0.1",
        "Port": "3306",
        "Database": "mydb",
        "Username": "root",
        "Password": "root",
        "Charset": "utf8mb4",
        "Prefix": "wei_"
    },
    "View": {
        "RootPath": "view"
    }
}
```

### 3. 入口文件

```go
package main

import (
    "github.com/wycto/weigo"
    _ "myproject/route"
)

func main() {
    weigo.New()
    weigo.ConfigInit("config")
    weigo.Run()
}
```

### 4. 路由定义

```go
package route

import (
    "github.com/wycto/weigo"
    "github.com/wycto/weigo/controller"
    "myproject/controller/index"
)

func init() {
    controller.RouterStatic(weigo.Echo, "/static", "static")
    controller.Router(weigo.Echo, "/user/", "index", &index.UserController{})
}
```

### 5. 控制器

```go
package index

import (
    "github.com/wycto/weigo/controller"
)

type UserController struct {
    controller.Controller
}

func (c *UserController) Index() error {
    return c.Success("success", nil)
}

func (c *UserController) Info() error {
    id := c.C.Param("id")
    return c.Success("success", map[string]interface{}{"id": id})
}
```

## 项目结构

```
weigo/
├── weigo.go              # 框架入口
├── config/               # 配置模块
│   └── config.go
├── ctx/                  # 请求上下文
│   └── context.go
├── database/             # 数据库模块
│   └── database.go
├── model/                # 模型模块
│   ├── base.go
│   └── chain.go
├── controller/           # 控制器模块
│   ├── controller.go
│   └── websocket.go
├── middleware/            # 中间件
│   ├── logger.go
│   └── cors.go
├── response/             # 响应模块
│   ├── response.go
│   └── message.go
├── helper/               # 工具函数
│   ├── crypto.go
│   └── network.go
└── log/                  # 日志模块
    └── log.go
```

## 核心 API

### 控制器

```go
// 成功响应
c.Success("操作成功", data)

// 失败响应
c.Error("操作失败", nil)

// 系统消息响应
c.ResponseErrorMessage(&controller.Message{Code: 200101, Msg: "参数缺失"}, nil)

// 模板渲染
c.Assign("name", "value")
c.Display("")

// 字符串响应
c.ResponseString("hello")

// 获取参数
id := c.C.Param("id")
name := c.C.QueryParam("name")
page := c.GetInt("page", 1)
```

### 数据库

```go
// 链式查询
var users []User
model.NewChain(weigo.DB, "user").
    Fields("id,name,email").
    Where("status = ?", 1).
    Order("id desc").
    Page(1, 10).
    Find(&users)

// 直接查询
weigo.DB.Table("user").Where("id = ?", 1).First(&user)

// 新增
weigo.DB.Table("user").Create(&user)

// 更新
weigo.DB.Table("user").Where("id = ?", 1).Updates(data)

// 删除
weigo.DB.Table("user").Where("id = ?", 1).Delete(nil)
```

### WebSocket

```go
type ChatController struct {
    controller.WebSocketController
}

func (c *ChatController) HandleMessage(msg string) string {
    return "echo: " + msg
}

// 路由
wsGroup := weigo.Echo.Group("/ws")
wsGroup.GET("/chat", &ChatController{}.Handle)
```

## 版本要求

- Go 1.21+
- Echo v4.16+
- GORM v2.5+

## License

MIT License
