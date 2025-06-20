// Package consumer
// @Author fuzengyao
// @Date 2022-11-09 11:16:20
package consumer

import (
	"jihulab.com/guashu/gs-server/util/xnsq/consumer"
	"jihulab.com/guashu/gs-server/util/xnsq/server"
	"jihulab.com/guashu/gs-server/util/xnsq/service/registry"
)

func SocketConsumerHandler(opt registry.Options) server.ConsumerHandler {
	consumer.Options = opt
	return &SocketNsqConsumer{Options: opt}
}

type SocketNsqConsumer struct {
	registry.Options
}

func (l *SocketNsqConsumer) Run() {

}
