package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
)

// Logger 请求日志中间件
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			duration := time.Since(start)

			// 使用 Echo 内置 logger
			c.Logger().Infof("[%s] %s %s %d %v",
				time.Now().Format("2006-01-02 15:04:05"),
				c.Request().Method,
				c.Request().URL.Path,
				c.Response().Status,
				duration,
			)
			return err
		}
	}
}
