// Package producer
// @Author fuzengyao
// @Date 2022-11-10 11:07:51
package producer

import (
	"time"

	"github.com/nsqio/go-nsq"
	"jihulab.com/guashu/gs-server/lib/logger"
	"jihulab.com/guashu/gs-server/util/xnsq/service/registry"
)

var Separator = "@"
var producer *nsq.Producer
var Options registry.Options

func StartNsqProducer(opt registry.Options) {

	Options = opt
	if producer != nil {
		return
	}

	var err error
	cfg := nsq.NewConfig()
	producer, err = nsq.NewProducer(opt.NsqAddress, cfg)
	if nil != err {
		logger.Sugar.Info(err)
		panic("nsq new panic")
	}

	err = producer.Ping()
	if nil != err {
		logger.Sugar.Info(err)
		panic("nsq ping panic")
	}
}

type Producer struct {
}

func (p *Producer) Publish(topic, data string) {
	err := producer.Publish(topic, []byte(Separator+data))
	if err != nil {
		logger.Sugar.Error(err)
	}
}

func (p *Producer) DelayPublish(topic, data string, delay time.Duration) {
	err := producer.DeferredPublish(topic, delay, []byte(Separator+data))
	if err != nil {
		logger.Sugar.Error(err)
	}
}

func (p *Producer) DeferredPublish(deviceId, serverAddress, topic, data string, delay time.Duration) {
	if serverAddress == "" {
		//logger.Sugar.Infow("device not online:", "device_id:", deviceId, "topic:", topic)
		return
	}

	err := producer.DeferredPublish(serverAddress+"."+topic, delay, []byte(deviceId+Separator+data))
	if err != nil {
		logger.Sugar.Error(err)
	}
}

// 指定服务发布
func (p *Producer) AssignServerPublish(serverAddress, topic, data string) {
	err := producer.Publish(serverAddress+"."+topic, []byte(Separator+data))
	if err != nil {
		logger.Sugar.Error(err)
	}
}

func (p *Producer) AssignUuidPublish(uuid, topic, data string) {
	if uuid == "" {
		logger.Sugar.Error("uuid为空")
		return
	}
	err := producer.Publish(topic, []byte(uuid+Separator+data))
	if err != nil {
		logger.Sugar.Error(err)
	}
}

func (p *Producer) StopProducer() {
	if producer != nil {
		producer.Stop()
	}

	logger.Sugar.Info("stop nsq producer")
}

func (p *Producer) DeleteTopicByStop() {
	topic := NewTopic(Options)
	topic.DeleteByContain(Options.LocalAddress)
}
