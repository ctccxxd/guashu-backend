package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/judwhite/go-svc"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"jihulab.com/guashu/gs-server/config"
	"jihulab.com/guashu/gs-server/lib/db/mysql"
	"jihulab.com/guashu/gs-server/lib/db/redis"
	"jihulab.com/guashu/gs-server/lib/helper"
	"jihulab.com/guashu/gs-server/lib/logger"
	appMiddleware "jihulab.com/guashu/gs-server/logic/http/middleware"
	"jihulab.com/guashu/gs-server/logic/nsq/producer"
	"jihulab.com/guashu/gs-server/routes"
	"jihulab.com/guashu/gs-server/util/xetcd"
)

var Echo = echo.New()

type logicProgram struct {
	once sync.Once
}

func main() {
	p := &logicProgram{}
	if err := svc.Run(p, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL); err != nil {
		logger.Sugar.Fatal(err)
	}
}

// svc 服务运行框架 程序启动时执行Init+Start, 服务终止时执行Stop
func (p *logicProgram) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}

func (p *logicProgram) Start() error {

	// 日志
	logger.New(
		logger.Base(config.AppConf.App.Env, config.AppConf.App.Path.LogPath),
		logger.FileName(helper.CurrentFileName()),
	)

	// 启动app
	go func() {
		newApp()
	}()

	return nil
}

/*
*
启动app
*/
func newApp() {
	e := Echo

	e.HTTPErrorHandler = appMiddleware.GlobalErrorHandler

	e.Binder = new(helper.Binder)
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	routes.Register(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.AppConf.Server.Http.Port)))
}

func (p *logicProgram) Stop() error {
	p.once.Do(func() {
		defer mysql.Disconnect()
		defer redis.Disconnect()
		defer xetcd.Close()
		defer routes.CancelRoute(Echo)
		defer producer.LogicProducer.Stop()
	})
	return nil
}
