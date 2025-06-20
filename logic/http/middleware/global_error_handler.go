package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	applogger "jihulab.com/guashu/gs-server/lib/logger"
)

// GlobalErrorHandler 全局错误处理器
func GlobalErrorHandler(err error, c echo.Context) {
	// 跳过已响应的请求
	if c.Response().Committed {
		return
	}

	// 默认错误响应
	code := http.StatusInternalServerError
	message := "服务器内部错误"
	var internalError error

	// 处理 Echo 内置错误
	if httpErr, ok := err.(*echo.HTTPError); ok {
		code = httpErr.Code
		if msg, ok := httpErr.Message.(string); ok {
			message = msg
		} else {
			message = http.StatusText(code)
		}
	} else {
		internalError = err
	}

	// 构建响应
	response := map[string]interface{}{
		"code":    code,
		"message": message,
	}

	// 添加错误详情（调试模式或内部使用）
	if c.Echo().Debug && internalError != nil {
		response["details"] = internalError.Error()
	}

	// 记录错误日志
	if code >= http.StatusInternalServerError {
		applogger.Sugar.Error("错误: %v", internalError)
	} else {
		applogger.Sugar.Warnf("请求错误: %v", err)
	}

	// 返回 JSON 响应
	if err := c.JSON(code, response); err != nil {
		applogger.Sugar.Error("发送错误响应失败: %v", err)
	}
}
