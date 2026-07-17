package ctx

import (
	"github.com/labstack/echo/v4"
)

// Context 请求上下文封装
type Context struct {
	echo.Context
	AppName        string
	ControllerName string
	ActionName     string
}

// Has 参数是否存在（检查路径参数、query参数、form参数）
func (c *Context) Has(keys ...string) bool {
	for _, key := range keys {
		if c.Param(key) != "" {
			continue
		}
		if c.QueryParam(key) != "" {
			continue
		}
		if c.FormValue(key) != "" {
			continue
		}
		return false
	}
	return true
}

// HasAndNotEmpty 参数存在且不为空
func (c *Context) HasAndNotEmpty(keys ...string) bool {
	for _, key := range keys {
		v := c.GetParam(key)
		if v == "" || v == "false" || v == "null" || v == "0" {
			return false
		}
	}
	return true
}

// NotHas 参数不存在
func (c *Context) NotHas(keys ...string) bool {
	return !c.Has(keys...)
}

// NotHasOrEmpty 不存在或者为空
func (c *Context) NotHasOrEmpty(keys ...string) bool {
	return !c.HasAndNotEmpty(keys...)
}

// GetParam 获取参数（统一入口：路径 → query → form）
func (c *Context) GetParam(key string) string {
	if v := c.Param(key); v != "" {
		return v
	}
	if v := c.QueryParam(key); v != "" {
		return v
	}
	return c.FormValue(key)
}

// Get 获取路径参数
func (c *Context) Get(key string) string {
	return c.Param(key)
}

// Query 获取 query 参数
func (c *Context) Query(key string) string {
	return c.QueryParam(key)
}

// Post 获取 post 参数
func (c *Context) Post(key string) string {
	return c.FormValue(key)
}
