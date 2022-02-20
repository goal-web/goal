package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/queue"
	"strings"
)

func init() {
	configs["queue"] = func(env contracts.Env) interface{} {
		return queue.Config{
			Defaults: queue.Defaults{
				Connection: env.StringOption("queue.connection", "default"),
				Queue:      env.StringOption("queue.default", "default"),
			},
			Connections: map[string]contracts.Fields{
				"default": {
					"driver":  "kafka",
					"delay":   "delay_queue", // 延迟队列名
					"brokers": strings.Split(env.GetString("queue.kafka.brokers"), ","),
					//"dialer": &kafka.Dialer{}, // 自定义 kafka.Dialer
					"queue": []string{
						"default", "slow", "high",
					},
				},
				"nsq": {
					"driver":           "nsq",
					"address":          env.GetString("queue.nsq.address"),
					"lookup_addresses": strings.Split(env.GetString("queue.nsq.lookup_addresses"), ","),
					//"config": &nsq.Config{}, // 自定义 nsq 设置
					"queue": []string{
						"default", "slow", "high",
					},
				},
			},
			Failed: queue.FailedJobs{
				Database: env.StringOption("db.connection", "mysql"),
				Table:    "failed_jobs",
			},
			Workers: map[string]queue.Workers{ // 相当于 laravel 的 horizon 配置
				"local": { // 本地环境
					"default": { // 工作组
						Connection: env.StringOption("queue.connection", "default"), // 指定连接
						Tries:      3,                                               // 最大尝试次数
						Queue:      []string{"default", "slow", "high"},             // 处理指定队列
						Processes:  10,                                              // 十个协程(工人)
					},
				},
				"production": { // 生产环境

				},
			},
		}
	}
}
