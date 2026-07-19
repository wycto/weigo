package weigo

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/wycto/weigo/config"
	"github.com/wycto/weigo/database"
	"github.com/wycto/weigo/helper"
	weigoMiddleware "github.com/wycto/weigo/middleware"
)

var (
	RootPath, _ = os.Getwd()
	DS          = string(os.PathSeparator)
)

var (
	Echo *echo.Echo
	DB   *database.DB
)

// Config 全局配置（对外暴露）
var Config *config.Config

// New 创建框架实例
func New() *echo.Echo {
	Echo = echo.New()

	// 默认中间件
	Echo.Use(echoMiddleware.Recover())
	Echo.Use(echoMiddleware.Logger())

	// 自定义错误处理（JSON 格式）
	Echo.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}
		code := 500
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		c.JSON(code, map[string]interface{}{
			"code": code,
			"msg":  err.Error(),
			"data": nil,
		})
	}

	return Echo
}

// ConfigInit 初始化配置
func ConfigInit(configDir string) {
	Config = config.Load(configDir)
	if Config.App.UseDB {
		var err error
		DB, err = database.Connect(database.DBConfig{
			Driver:   Config.DB.Driver,
			Host:     Config.DB.Host,
			Port:     Config.DB.Port,
			Username: Config.DB.Username,
			Password: Config.DB.Password,
			Database: Config.DB.Database,
			Charset:  Config.DB.Charset,
			Prefix:   Config.DB.Prefix,
		})
		if err != nil {
			panic("数据库连接失败: " + err.Error())
		}
	}
}

// Run 启动服务
func Run() {
	if Echo == nil {
		New()
	}

	// 添加默认中间件
	Echo.Use(weigoMiddleware.CORS())

	// 优雅关闭
	go func() {
		ip := helper.GetIP()
		fmt.Printf("服务已启动，地址：%s:%s OR 127.0.0.1:%s\n", ip, Config.App.Port, Config.App.Port)
		if err := Echo.Start(":" + Config.App.Port); err != nil {
			Echo.Logger.Info("服务已停止")
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := Echo.Shutdown(ctx); err != nil {
		Echo.Logger.Fatal(err)
	}
}
