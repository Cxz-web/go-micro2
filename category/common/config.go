package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	soucre := consul.NewSource(
		consul.WithAddress(host+"："+string(port)),
		consul.WithPrefix(prefix),
		// 是否移除前缀，
		consul.StripPrefix(true),
	)

	config, err := config.NewConfig()

	if err != nil {
		return config, err
	}

	return config, config.Load(soucre)
}
