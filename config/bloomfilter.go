package config

import (
	"github.com/goal-web/bloomfilter"
	"github.com/goal-web/contracts"
)

func init() {
	configs["bloomfilter"] = func(env contracts.Env) any {
		return bloomfilter.Config{
			Default: "default",
			Filters: bloomfilter.Filters{
				"default": contracts.Fields{
					"driver":   "file", // 将数据序列化到文件中，不支持分布式
					"size":     100000,
					"k":        .01,
					"filepath": "/Users/qbhy/project/go/goal-web/goal/storage/default", // 完整路径
				},
				"users": contracts.Fields{
					"driver": "redis", // 通过 redis bitmap 存储，支持分布式
					"size":   100000,
					"k":      .01,
					"key":    "bloomfilter:{name}", // {name} 表示该 filter 的key，这里是 users
					//"connection": "cache", // redis 连接
				},
			},
		}
	}
}
