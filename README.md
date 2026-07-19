# WeiGo 2.0

基于 Echo + GORM 的 Go Web 框架，保留 PHP 风格的开发体验。

## 特性

- **MVC 架构**：控制器、模型、视图分离
- **Echo HTTP 框架**：高性能、易用
- **GORM ORM**：安全的数据库操作、链式查询
- **反射路由**：自动匹配控制器方法（PHP 风格）
- **模板渲染**：支持 HTML 模板（Debug 模式不缓存，生产环境缓存）
- **WebSocket**：实时通信支持
- **中间件**：日志、跨域、恢复
- **优雅关闭**：信号处理、平滑停机
- **连接池**：数据库连接管理（SQLite 自动跳过）

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

**支持的数据库驱动**：`mysql`、`postgres`、`sqlite`

**CORS 配置**（可选）：

```json
{
    "App": {
        "Debug": true,
        "AllowOrigins": ["http://localhost:3000"],
        "AllowHeaders": ["Content-Type", "Authorization"]
    }
}
```

> ⚠️ **CORS 注意事项**：
> - `Debug: true` + 未配置 `AllowOrigins` → 允许所有来源
> - `Debug: true` + 配置了 `AllowOrigins` → **只允许**列出的来源
> - `Debug: false` + 未配置 `AllowOrigins` → 拒绝所有来源
> - `AllowOrigins` 配置 `["*"]` → 允许所有来源

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

// 默认动作（对应 /user/ 或 /user/index）
func (c *UserController) Index() error {
    return c.Success("success", nil)
}

// 获取 URL 参数：/user/info/id/42/name/john
func (c *UserController) Info() error {
    id := c.GetParam("id")      // 路径参数优先
    name := c.GetString("name", "") // 支持默认值
    page := c.GetInt("page", 1) // 整数，支持默认值
    price := c.GetFloat64("price", 0.0)
    active := c.GetBool("active", true)
    
    return c.Success("success", map[string]interface{}{
        "id": id, "name": name, "page": page,
    })
}

// 获取 JSON body（POST/PUT/PATCH 自动解析）
func (c *UserController) Create() error {
    name := c.GetString("name", "")
    age := c.GetInt("age", 0)
    
    if !c.HasAndNotEmpty("name") {
        return c.BadRequest("名称不能为空")
    }
    
    return c.Success("创建成功", nil)
}
```

**控制器方法签名要求**（反射路由）：
- 必须是 `func (c *XxxController) MethodName() error`
- 零参数（除了 receiver）
- 返回值必须是 `error`
- 不能在基类 `Controller` 或 `WebSocketController` 上定义
- 方法名匹配**不区分大小写**，连字符会被忽略

## 项目结构

```
weigo/
├── weigo.go              # 框架入口
├── config/               # 配置模块
│   ├── config.go
│   └── config_test.go
├── database/             # 数据库模块
│   ├── database.go
│   └── database_test.go
├── model/                # 模型模块
│   ├── base.go           # BaseModel 基类
│   ├── chain.go          # 链式查询构建器
│   └── chain_test.go
├── controller/           # 控制器模块
│   ├── controller.go     # 反射路由 + 控制器基类
│   ├── websocket.go      # WebSocket 控制器
│   └── controller_test.go
├── middleware/            # 中间件
│   ├── logger.go         # 请求日志
│   ├── cors.go           # 跨域处理
│   └── cors_test.go
├── response/             # 响应模块
│   ├── response.go       # 统一响应格式
│   ├── message.go        # 系统消息码
│   └── response_test.go
├── helper/               # 工具函数
│   ├── crypto.go         # MD5/SHA256
│   ├── network.go        # IP/主机名获取
│   ├── cors.go           # CORS 辅助函数
│   └── crypto_test.go
└── AGENTS.md             # Agent 指南
```

## 核心 API

### 控制器响应

```go
// 成功响应 (code: 0)
c.Success("操作成功", data)

// 失败响应 (code: 1, HTTP 200)
c.Error("操作失败", nil)

// 带 HTTP 状态码的错误响应
c.ErrorWithStatus(400, 200101, "参数错误", nil)
c.BadRequest("参数错误")       // 400
c.Unauthorized("未登录")       // 401
c.Forbidden("无权限")          // 403
c.NotFound("数据不存在")       // 404
c.ServerError("服务器错误")    // 500

// 系统消息码响应
c.ResponseErrorMessage(&response.ParamsMissing, nil)
c.ResponseString("hello")
```

**系统消息码**：

| 常量 | Code | 含义 |
|------|------|------|
| `ParamsMissing` | 200101 | 参数缺失 |
| `ParamsInvalid` | 200102 | 参数无效 |
| `DataNone` | 200201 | 数据不存在 |
| `DataExists` | 200202 | 数据已存在 |
| `SysError` | 200501 | 系统错误 |

### 参数获取

```go
// 统一获取（路径 → query → form）
value := c.GetParam("key")

// 分类获取
id := c.Get("id")              // 路径参数
name := c.Query("name")        // query 参数
post := c.Post("field")        // form 参数

// 类型化获取（支持默认值）
page := c.GetInt("page", 1)
price := c.GetFloat64("price", 0.0)
active := c.GetBool("active", true)
name := c.GetString("name", "")

// 参数检查
c.Has("id", "name")            // 参数存在
c.HasAndNotEmpty("name")       // 参数存在且不为空
c.NotHas("optional")           // 参数不存在

// 获取 JSON body（POST/PUT/PATCH）
body := c.GetBody()            // map[string]interface{}
```

### 数据库操作

```go
// 链式查询（推荐）
var users []User
model.NewChain(weigo.DB, "user").
    Fields("id,name,email").
    Where("status = ?", 1).
    Where("age > ?", 18).
    Order("id desc").
    Page(1, 10).
    Find(&users)

// 查询单条
var user User
model.NewChain(weigo.DB, "user").
    Where("id = ?", 1).
    FindOne(&user)

// 计数
var count int64
model.NewChain(weigo.DB, "user").
    Where("status = ?", 1).
    Count(&count)

// 新增
model.NewChain(weigo.DB, "user").Insert(&user)

// 更新
model.NewChain(weigo.DB, "user").
    Where("id = ?", 1).
    Update(map[string]interface{}{"name": "new name"})

// 删除
model.NewChain(weigo.DB, "user").
    Where("id = ?", 1).
    Delete()

// 直接使用 GORM（表前缀自动添加）
weigo.DB.Table("user").Where("id = ?", 1).First(&user)
weigo.DB.Table("user").Create(&user)
```

> 💡 **表前缀**：配置中的 `Prefix` 会自动添加到表名，无需手动拼接。

### 模板渲染

```go
// 赋值并渲染
c.Assign("title", "首页")
c.Assign("users", users)
return c.Display("")  // 空字符串使用默认视图路径

// 直接输出 HTML（不解析模板）
return c.ShowHtml("path/to/file.html")

// 清空模板缓存（生产环境修改模板后调用）
controller.ClearTemplateCache()
```

**默认视图路径**：`{View.RootPath}/{AppName}/{ControllerName}/{ActionName}.html`

### WebSocket

```go
type ChatController struct {
    controller.WebSocketController
}

func (c *ChatController) HandleMessage(msg string) string {
    return "echo: " + msg
}

// 路由注册
wsGroup := weigo.Echo.Group("/ws")
wsGroup.GET("/chat", &ChatController{}.Handle)
```

### WebSocket 管理器

```go
manager := controller.NewWebSocketManager()
go manager.Run()

// 注册/注销连接
manager.Register(conn)
manager.Unregister(conn)

// 广播消息
manager.Broadcast([]byte("hello"))

// 获取连接信息
conns := manager.GetConnections()
count := manager.GetConnectionCount()
```

## 测试

```bash
# 运行所有测试
go test ./...

# 运行单个包的测试
go test ./config/...
go test ./controller/...

# 运行特定测试
go test -run TestLoadValidConfig ./config/

# 显示测试覆盖
go test -cover ./...
```

**测试特点**：
- 白盒测试（同包访问私有字段）
- SQLite 内存数据库（`:memory:`）
- `httptest.NewRecorder()` + `echo.NewContext()` 测试控制器
- `t.TempDir()` 隔离文件系统测试

## 版本要求

- Go 1.26+
- Echo v4.11+
- GORM v1.25+

## License

MIT License
