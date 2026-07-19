package controller

import (
	"encoding/json"
	"html/template"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/wycto/weigo/config"
	"github.com/wycto/weigo/response"
)

// 模板缓存
var templateCache sync.Map

// Controller 控制器基类
type Controller struct {
	C              echo.Context
	AppName        string
	ControllerName string
	ActionName     string
	data           map[string]interface{}
	body           map[string]interface{}
}

// Init 初始化控制器
func (c *Controller) Init(ctx echo.Context) {
	c.C = ctx
	c.data = make(map[string]interface{})
	c.body = make(map[string]interface{})

	// 仅在有 body 的请求方法中解析 JSON body
	method := ctx.Request().Method
	if method == "POST" || method == "PUT" || method == "PATCH" {
		if ctx.Request().Body != nil {
			json.NewDecoder(ctx.Request().Body).Decode(&c.body)
		}
	}
}

// --- 参数检查（兼容 PHP 风格） ---

// Has 参数是否存在（检查路径参数、query参数、form参数）
func (c *Controller) Has(keys ...string) bool {
	for _, key := range keys {
		if c.C.Param(key) != "" {
			continue
		}
		if c.C.QueryParam(key) != "" {
			continue
		}
		if c.C.FormValue(key) != "" {
			continue
		}
		return false
	}
	return true
}

// HasAndNotEmpty 参数存在且不为空
func (c *Controller) HasAndNotEmpty(keys ...string) bool {
	for _, key := range keys {
		v := c.GetParam(key)
		if v == "" || v == "false" || v == "null" || v == "0" {
			return false
		}
	}
	return true
}

// NotHas 参数不存在
func (c *Controller) NotHas(keys ...string) bool {
	return !c.Has(keys...)
}

// NotHasOrEmpty 不存在或者为空
func (c *Controller) NotHasOrEmpty(keys ...string) bool {
	return !c.HasAndNotEmpty(keys...)
}

// GetParam 获取参数（统一入口：路径 → query → form）
func (c *Controller) GetParam(key string) string {
	if v := c.C.Param(key); v != "" {
		return v
	}
	if v := c.C.QueryParam(key); v != "" {
		return v
	}
	return c.C.FormValue(key)
}

// Get 获取路径参数
func (c *Controller) Get(key string) string {
	return c.C.Param(key)
}

// Query 获取 query 参数
func (c *Controller) Query(key string) string {
	return c.C.QueryParam(key)
}

// Post 获取 post 参数
func (c *Controller) Post(key string) string {
	return c.C.FormValue(key)
}

// --- 模板 ---

// Assign 模板赋值
func (c *Controller) Assign(key string, value interface{}) {
	c.data[key] = value
}

// getTemplate 获取或解析并缓存模板
func getTemplate(viewName string) (*template.Template, error) {
	// Debug 模式不缓存，方便开发时修改模板即时生效
	if config.GlobalConfig != nil && config.GlobalConfig.App.Debug {
		return template.ParseFiles(viewName)
	}

	if t, ok := templateCache.Load(viewName); ok {
		return t.(*template.Template), nil
	}

	t, err := template.ParseFiles(viewName)
	if err != nil {
		return nil, err
	}

	templateCache.Store(viewName, t)
	return t, nil
}

// ClearTemplateCache 清空模板缓存（模板文件变更后调用）
func ClearTemplateCache() {
	templateCache = sync.Map{}
}

// Display 模板渲染
func (c *Controller) Display(viewName string) error {
	if viewName == "" {
		viewName = c.resolveViewPath()
	}
	t, err := getTemplate(viewName)
	if err != nil {
		return c.C.String(500, err.Error())
	}
	return t.Execute(c.C.Response().Writer, c.data)
}

// ShowHtml 直接输出 HTML（不解析模板）
func (c *Controller) ShowHtml(viewName string) error {
	if viewName == "" {
		viewName = c.resolveViewPath()
	}
	content, err := os.ReadFile(viewName)
	if err != nil {
		return c.C.String(500, err.Error())
	}
	return c.C.HTML(200, string(content))
}

// resolveViewPath 根据约定解析视图文件路径
func (c *Controller) resolveViewPath() string {
	appName := c.AppName
	if appName == "" {
		appName = config.GlobalConfig.App.DefaultAppName
	}
	controllerName := c.ControllerName
	if controllerName == "" {
		controllerName = config.GlobalConfig.App.DefaultControllerName
	}
	actionName := c.ActionName
	if actionName == "" {
		actionName = config.GlobalConfig.App.DefaultActionName
	}
	return filepath.Join(
		config.RootPath,
		config.GlobalConfig.View.RootPath,
		appName,
		strings.ToLower(controllerName),
		strings.ToLower(actionName)+".html",
	)
}

// --- 响应（委托给 response 包，统一 API） ---

// Success 成功响应
func (c *Controller) Success(msg string, data interface{}) error {
	return response.Success(c.C, msg, data)
}

// Error 失败响应（HTTP 200 + code:1，兼容旧格式）
func (c *Controller) Error(msg string, data interface{}) error {
	return response.Error(c.C, msg, data)
}

// ErrorWithStatus 带 HTTP 状态码的错误响应
func (c *Controller) ErrorWithStatus(httpStatus int, code int, msg string, data interface{}) error {
	return response.ErrorWithStatus(c.C, httpStatus, code, msg, data)
}

// BadRequest 400 参数错误
func (c *Controller) BadRequest(msg string) error {
	return response.BadRequest(c.C, msg)
}

// Unauthorized 401 未认证
func (c *Controller) Unauthorized(msg string) error {
	return response.Unauthorized(c.C, msg)
}

// Forbidden 403 无权限
func (c *Controller) Forbidden(msg string) error {
	return response.Forbidden(c.C, msg)
}

// NotFound 404 资源不存在
func (c *Controller) NotFound(msg string) error {
	return response.NotFound(c.C, msg)
}

// ServerError 500 服务器错误
func (c *Controller) ServerError(msg string) error {
	return response.ServerError(c.C, msg)
}

// ResponseErrorMessage 使用系统消息码响应
func (c *Controller) ResponseErrorMessage(msg *response.Message, data interface{}) error {
	return response.ResponseMessage(c.C, msg, data)
}

// ResponseString 输出字符串
func (c *Controller) ResponseString(msg string) error {
	return c.C.String(200, msg)
}

// MethodNotAllowed 方法不允许
func (c *Controller) MethodNotAllowed() error {
	return c.C.String(405, "Method Not Allowed")
}

// --- 反射路由分发 ---

// isActionMethod 判断一个方法是否是用户定义的 action 方法
// 规则：无参数、返回一个 error、且不是 Controller 基类方法
func isActionMethod(typeInfo reflect.Type, method reflect.Method) bool {
	// 跳过 Init
	if method.Name == "Init" {
		return false
	}

	// 检查签名：func (c *XxxController) Method() error
	if method.Type.NumIn() != 1 { // 只有 receiver
		return false
	}
	if method.Type.NumOut() != 1 { // 只返回 error
		return false
	}
	if method.Type.Out(0).String() != "error" {
		return false
	}

	// 跳过基类方法（通过检查方法是否在 Controller 或 WebSocketController 上定义）
	baseControllerType := reflect.TypeOf((*Controller)(nil)).Elem()
	wsControllerType := reflect.TypeOf((*WebSocketController)(nil)).Elem()

	if method.Func.IsValid() {
		recvType := method.Type.In(0)
		if recvType.Kind() == reflect.Ptr {
			recvType = recvType.Elem()
		}
		if recvType == baseControllerType || recvType == wsControllerType {
			return false
		}
	}

	return true
}

// Router 自动路由（反射调用控制器方法，支持 GET/POST/PUT/DELETE）
func Router(e *echo.Echo, path string, appName string, controller interface{}) {
	typeInfo := reflect.TypeOf(controller)

	// 构建 action 映射（小写 -> 原名），自动识别用户定义的 action 方法
	actionMap := make(map[string]string)
	for i := 0; i < typeInfo.NumMethod(); i++ {
		m := typeInfo.Method(i)
		if isActionMethod(typeInfo, m) {
			actionMap[strings.ToLower(m.Name)] = m.Name
		}
	}

	handler := func(ctx echo.Context) error {
		newController := reflect.New(typeInfo.Elem()).Interface()

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

		// 设置上下文信息到 Controller 字段
		controllerName := typeInfo.Elem().Name()
		actionName := config.GlobalConfig.App.DefaultActionName
		if len(parts) > 0 && parts[0] != "" {
			actionName = parts[0]
		}

		// 设置 Controller 字段
		setFieldValue(newController, "AppName", appName)
		setFieldValue(newController, "ControllerName", controllerName)
		setFieldValue(newController, "ActionName", actionName)

		// 处理 URL 参数（/key/value/key/value 格式）
		for i := 1; i < len(parts); i += 2 {
			if i+1 < len(parts) {
				ctx.Set(parts[i], parts[i+1])
			} else {
				ctx.Set(parts[i], "")
			}
		}

		// 同时保留 echo.Context 上的设置（兼容旧代码）
		ctx.Set("appName", appName)
		ctx.Set("controllerName", controllerName)
		ctx.Set("actionName", actionName)

		// 查找 action 方法（去除连字符后匹配）
		cleanAction := strings.ReplaceAll(actionName, "-", "")
		if methodName, ok := actionMap[strings.ToLower(cleanAction)]; ok {
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
	}

	e.GET(path+"*", handler)
	e.POST(path+"*", handler)
	e.PUT(path+"*", handler)
	e.DELETE(path+"*", handler)
}

// setFieldValue 通过反射设置控制器字段值
func setFieldValue(controller interface{}, fieldName string, value string) {
	v := reflect.ValueOf(controller).Elem()
	field := v.FieldByName(fieldName)
	if field.IsValid() && field.CanSet() {
		field.SetString(value)
	}
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
		if val, ok := c.body[key]; ok {
			switch v := val.(type) {
			case float64:
				return int(v)
			case int:
				return v
			case string:
				if i, err := strconv.Atoi(v); err == nil {
					return i
				}
			}
		}
		return defaultVal
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return defaultVal
	}
	return i
}

// GetFloat64 获取浮点数参数
func (c *Controller) GetFloat64(key string, defaultVal float64) float64 {
	v := c.C.Param(key)
	if v == "" {
		v = c.C.QueryParam(key)
	}
	if v == "" {
		if val, ok := c.body[key]; ok {
			switch v := val.(type) {
			case float64:
				return v
			case int:
				return float64(v)
			case string:
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					return f
				}
			}
		}
		return defaultVal
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return defaultVal
	}
	return f
}

// GetBool 获取布尔参数（支持 true/false/1/0/"true"/"false"）
func (c *Controller) GetBool(key string, defaultVal bool) bool {
	v := c.GetParam(key)
	if v == "" {
		if val, ok := c.body[key]; ok {
			switch v := val.(type) {
			case bool:
				return v
			case float64:
				return v != 0
			case int:
				return v != 0
			case string:
				return parseBool(v)
			}
		}
		return defaultVal
	}
	return parseBool(v)
}

// parseBool 解析布尔值字符串
func parseBool(s string) bool {
	switch strings.ToLower(s) {
	case "true", "1", "yes":
		return true
	case "false", "0", "no", "":
		return false
	default:
		return false
	}
}

// GetString 获取字符串参数
func (c *Controller) GetString(key string, defaultVal string) string {
	v := c.C.Param(key)
	if v == "" {
		v = c.C.QueryParam(key)
	}
	if v == "" {
		if val, ok := c.body[key]; ok {
			if str, ok := val.(string); ok {
				return str
			}
		}
		return defaultVal
	}
	return v
}

// GetBody 获取解析后的JSON body
func (c *Controller) GetBody() map[string]interface{} {
	return c.body
}
