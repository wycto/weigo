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

// Error 失败响应
func Error(c echo.Context, msg string, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 1,
		"msg":  msg,
		"data": data,
	})
}

// ResponseMessage 使用系统消息码响应
func ResponseMessage(c echo.Context, msg *Message, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": msg.Code,
		"msg":  msg.Msg,
		"data": data,
	})
}
