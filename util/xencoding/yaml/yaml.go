// Package yaml
// @Author fuzengyao
// @Date 2022-11-09 11:17:16
package yaml

import (
	"gopkg.in/yaml.v3"
	"jihulab.com/guashu/gs-server/util/xencoding"
)

const Name = "yaml"

func init() {
	xencoding.RegisterCodec(codec{})
}

type codec struct{}

func (codec) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func (codec) Name() string {
	return Name
}
