// Package consumer
// @Author fuzengyao
// @Date 2022-11-09 11:15:37
package consumer

import (
	"jihulab.com/guashu/gs-server/util/xnsq/consumer"
	"jihulab.com/guashu/gs-server/util/xnsq/server"
	"jihulab.com/guashu/gs-server/util/xnsq/service/registry"
)

func LogicConsumerHandler(opt registry.Options) server.ConsumerHandler {
	consumer.Options = opt
	return &LogicNsqConsumer{opt}
}

type LogicNsqConsumer struct {
	registry.Options
}

func (l *LogicNsqConsumer) Run() {

}
