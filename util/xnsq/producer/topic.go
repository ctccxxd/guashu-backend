// Package producer
// @Author fuzengyao
// @Date 2022-11-10 11:07:46
package producer

import (
	"strings"

	"jihulab.com/guashu/gs-server/lib/logger"
	"jihulab.com/guashu/gs-server/util/xnsq/api"
	"jihulab.com/guashu/gs-server/util/xnsq/service/registry"
)

var nsqApiClient *api.Client

func NewTopic(opt registry.Options) *Topic {
	return &Topic{opt: opt}
}

type Topic struct {
	opt registry.Options
}

// 删除
func (t *Topic) Delete(topic string) {
	err := nsqApiClient.Topic().Delete(topic)
	if err != nil {
		logger.Sugar.Error(err)
	}
}

// 删除类似的topic
func (t *Topic) DeleteByContain(value string) {

	// 获取所有的 topic
	nsqApiClient = api.NewClient(t.opt)
	topics, _ := nsqApiClient.Topic().QueryAll()

	for _, topic := range topics {
		if strings.Contains(topic, value) {
			t.Delete(topic)
		}
	}
}
