package middleware

import "github.com/labstack/echo/v4"

// CORS 跨域中间件
func CORS() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 只在需要时设置跨域头
			origin := c.Request().Header.Get("Origin")
			if origin != "" {
				c.Response().Header().Set("Access-Control-Allow-Origin", origin)
				c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			}

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(204)
			}
			return next(c)
		}
	}
}
