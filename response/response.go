package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Success 成功响应
func Success(c echo.Context, msg string, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  msg,
		"data": data,
	})
}

// Error 失败响应（HTTP 200 + code:1，兼容旧格式）
func Error(c echo.Context, msg string, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 1,
		"msg":  msg,
		"data": data,
	})
}

// ErrorWithStatus 带 HTTP 状态码的错误响应
func ErrorWithStatus(c echo.Context, httpStatus int, code int, msg string, data interface{}) error {
	return c.JSON(httpStatus, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// BadRequest 400 参数错误
func BadRequest(c echo.Context, msg string) error {
	return ErrorWithStatus(c, http.StatusBadRequest, ParamsMissing.Code, msg, nil)
}

// Unauthorized 401 未认证
func Unauthorized(c echo.Context, msg string) error {
	return ErrorWithStatus(c, http.StatusUnauthorized, 200401, msg, nil)
}

// Forbidden 403 无权限
func Forbidden(c echo.Context, msg string) error {
	return ErrorWithStatus(c, http.StatusForbidden, 200403, msg, nil)
}

// NotFound 404 资源不存在
func NotFound(c echo.Context, msg string) error {
	return ErrorWithStatus(c, http.StatusNotFound, DataNone.Code, msg, nil)
}

// ServerError 500 服务器错误
func ServerError(c echo.Context, msg string) error {
	return ErrorWithStatus(c, http.StatusInternalServerError, SysError.Code, msg, nil)
}

// ResponseMessage 使用系统消息码响应
func ResponseMessage(c echo.Context, msg *Message, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": msg.Code,
		"msg":  msg.Msg,
		"data": data,
	})
}
