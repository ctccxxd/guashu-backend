package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	applogger "jihulab.com/guashu/gs-server/lib/logger"
	"jihulab.com/guashu/gs-server/logic/http/model"
	"jihulab.com/guashu/gs-server/logic/http/service"
)

// Handler
func InterviewCoaching(ctx echo.Context) error {
	promotion := new(model.InterviewCoachingPromotion)
	if err := ctx.Bind(promotion); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "表单绑定失败: "+err.Error())
	}
	interviewCoachingService := service.NewInterviewCoachingService()
	err := interviewCoachingService.SendEmail(*promotion)
	if err != nil {
		applogger.Sugar.Error("调用SendEmail失败: %v\n", err)
		return err
	}
	return nil
}
