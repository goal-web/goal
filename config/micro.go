package config

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/goal-web/contracts"
	"github.com/goal-web/micro"
	micro2 "go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func init() {
	configs["micro"] = func(env contracts.Env) any {
		return micro.Config{
			CustomOptions: []micro2.Option{
				micro2.Name("hello"),
				micro2.Version("latest"),
				micro2.Registry(etcd.NewRegistry(
					registry.Addrs(env.GetString("micro.etcd.address")),
				)),
			},
		}
	}
}
