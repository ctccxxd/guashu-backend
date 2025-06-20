package routes

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"jihulab.com/guashu/gs-server/lib/logger"
	"jihulab.com/guashu/gs-server/logic/http/controller"
)

func Register(api *echo.Echo) {

	// 宣传推广相关接口
	orderGroup := api.Group("/api")
	orderGroup.POST("/interview-coaching", controller.InterviewCoaching)
}

// 结束router
func CancelRoute(e *echo.Echo) {
	if e == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Sugar.Fatal(err)
	}
	logger.Sugar.Info("stop router")
}
