package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/filesystem"
	"os"
)

func init() {
	configs["filesystem"] = func(env contracts.Env) any {
		return filesystem.Config{
			Default: env.GetString("filesystem.driver"),
			Disks: map[string]contracts.Fields{
				"local": {
					"driver": "local",
					"root":   env.GetString("filesystem.root"),
					"perm":   os.ModePerm,
				},
				"qiniu": {
					"driver":     "qiniu",
					"ttl":        3600, // 私有 url 有效期，单位秒
					"private":    env.GetBool("qiniu.private"),
					"domain":     env.GetBool("qiniu.domain"), // example: https://image.example.com"
					"bucket":     env.GetString("qiniu.bucket"),
					"access_key": env.GetString("qiniu.access.key"),
					"secret_key": env.GetString("qiniu.secret.key"),
				},
			},
		}
	}
}
