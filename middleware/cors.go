package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/wycto/weigo/config"
	"github.com/wycto/weigo/helper"
)

// CORS 跨域中间件
func CORS() echo.MiddlewareFunc {
	// 默认允许的请求头
	defaultHeaders := "Content-Type, Authorization"

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			origin := c.Request().Header.Get("Origin")
			if origin == "" {
				return next(c)
			}

			if !helper.IsOriginAllowed(origin) {
				return next(c)
			}

			c.Response().Header().Set("Access-Control-Allow-Origin", origin)
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

			// 使用配置的 AllowHeaders，未配置则用默认值
			headers := defaultHeaders
			if cfg := config.GlobalConfig; cfg != nil && len(cfg.App.AllowHeaders) > 0 {
				headers = strings.Join(cfg.App.AllowHeaders, ", ")
			}
			c.Response().Header().Set("Access-Control-Allow-Headers", headers)

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(204)
			}
			return next(c)
		}
	}
}
