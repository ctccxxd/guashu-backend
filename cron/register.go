package cron

import (
	"jihulab.com/guashu/gs-server/cron/base"
)

func Register(address, password string) {
	cron := base.StartCronTab(address, password)

	//注册定时任务
	go cron.Run()
}
