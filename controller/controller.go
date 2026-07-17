package controller

import (
	"html/template"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/wycto/weigo/config"
)

// Controller 控制器基类
type Controller struct {
	C    echo.Context
	data map[string]interface{}
}

// Init 初始化控制器
func (c *Controller) Init(ctx echo.Context) {
	c.C = ctx
	c.data = make(map[string]interface{})
}

// GetAppName 获取应用名
func (c *Controller) GetAppName() string {
	if v := c.C.Get("appName"); v != nil {
		return v.(string)
	}
	return config.GlobalConfig.App.DefaultAppName
}

// Assign 模板赋值
func (c *Controller) Assign(key string, value interface{}) {
	c.data[key] = value
}

// Display 模板渲染
func (c *Controller) Display(viewName string) error {
	if viewName == "" {
		appName := c.GetAppName()
		controllerName := c.C.Get("controllerName").(string)
		actionName := c.C.Get("actionName").(string)
		viewName = filepath.Join(
			config.RootPath,
			config.GlobalConfig.View.RootPath,
			appName,
			strings.ToLower(controllerName),
			strings.ToLower(actionName)+".html",
		)
	}
	t, err := template.ParseFiles(viewName)
	if err != nil {
		return c.C.String(500, err.Error())
	}
	return t.Execute(c.C.Response().Writer, c.data)
}

// ShowHtml 直接输出 HTML（不解析模板）
func (c *Controller) ShowHtml(viewName string) error {
	if viewName == "" {
		appName := c.GetAppName()
		controllerName := c.C.Get("controllerName").(string)
		actionName := c.C.Get("actionName").(string)
		viewName = filepath.Join(
			config.RootPath,
			config.GlobalConfig.View.RootPath,
			appName,
			strings.ToLower(controllerName),
			strings.ToLower(actionName)+".html",
		)
	}
	content, err := os.ReadFile(viewName)
	if err != nil {
		return c.C.String(500, err.Error())
	}
	return c.C.HTML(200, string(content))
}

// Success 成功响应
func (c *Controller) Success(msg string, data interface{}) error {
	return c.C.JSON(200, map[string]interface{}{
		"code": 0,
		"msg":  msg,
		"data": data,
	})
}

// Error 失败响应
func (c *Controller) Error(msg string, data interface{}) error {
	return c.C.JSON(200, map[string]interface{}{
		"code": 1,
		"msg":  msg,
		"data": data,
	})
}

// ResponseErrorMessage 使用系统消息码响应
func (c *Controller) ResponseErrorMessage(msg *Message, data interface{}) error {
	return c.C.JSON(200, map[string]interface{}{
		"code": msg.Code,
		"msg":  msg.Msg,
		"data": data,
	})
}

// ResponseString 输出字符串
func (c *Controller) ResponseString(msg string) error {
	return c.C.String(200, msg)
}

// MethodNotAllowed 方法不允许
func (c *Controller) MethodNotAllowed() error {
	return c.C.String(405, "Method Not Allowed")
}

// --- 反射路由分发（保持 PHP 风格） ---

// Message 系统消息（用于 ResponseErrorMessage）
type Message struct {
	Code int
	Msg  string
}

// Router 自动路由（反射调用控制器方法）
func Router(e *echo.Echo, path string, appName string, controller interface{}) {
	// 存储控制器类型信息用于反射
	typeInfo := reflect.TypeOf(controller)

	// 构建 action 映射（小写 -> 原名）
	actionMap := make(map[string]string)
	for i := 0; i < typeInfo.NumMethod(); i++ {
		m := typeInfo.Method(i)
		name := m.Name
		// 跳过 Init 方法
		if name == "Init" {
			continue
		}
		actionMap[strings.ToLower(name)] = name
	}

	e.GET(path+"*", func(ctx echo.Context) error {
		// 创建新的控制器实例（避免并发问题）
		newController := reflect.New(typeInfo.Elem()).Interface()

		// 初始化控制器
		initMethod := reflect.ValueOf(newController).MethodByName("Init")
		if initMethod.IsValid() {
			initMethod.Call([]reflect.Value{reflect.ValueOf(ctx)})
		}

		// 解析 URL
		urlPath := ctx.Request().URL.Path
		mvcPath := strings.TrimPrefix(urlPath, path)
		parts := strings.FieldsFunc(mvcPath, func(r rune) bool {
			return r == '/'
		})

		// 设置上下文
		ctx.Set("appName", appName)
		ctx.Set("controllerName", typeInfo.Elem().Name())
		actionName := config.GlobalConfig.App.DefaultActionName
		if len(parts) > 0 && parts[0] != "" {
			actionName = parts[0]
		}

		// 处理 URL 参数（/key/value/key/value 格式）
		for i := 1; i < len(parts); i += 2 {
			if i+1 < len(parts) {
				ctx.Set(parts[i], parts[i+1])
			} else {
				ctx.Set(parts[i], "")
			}
		}

		ctx.Set("actionName", actionName)

		// 查找 action 方法
		if methodName, ok := actionMap[strings.ToLower(actionName)]; ok {
			method := reflect.ValueOf(newController).MethodByName(methodName)
			if method.IsValid() {
				result := method.Call(nil)
				if len(result) > 0 && result[0].Interface() != nil {
					return result[0].Interface().(error)
				}
				return nil
			}
		}

		return ctx.String(404, "404 page not found: "+urlPath)
	})
}

// RouterStatic 静态资源路由
func RouterStatic(e *echo.Echo, url string, path string) {
	if path == "" {
		path = url
	}
	e.Static(url, filepath.Join(config.RootPath, path))
}

// GetInt 获取整数参数
func (c *Controller) GetInt(key string, defaultVal int) int {
	v := c.C.Param(key)
	if v == "" {
		v = c.C.QueryParam(key)
	}
	if v == "" {
		return defaultVal
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return defaultVal
	}
	return i
}

// GetString 获取字符串参数
func (c *Controller) GetString(key string, defaultVal string) string {
	v := c.C.Param(key)
	if v == "" {
		v = c.C.QueryParam(key)
	}
	if v == "" {
		return defaultVal
	}
	return v
}
