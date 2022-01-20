package utils

import (
	"github.com/goal-web/contracts"
	"io/ioutil"
	"strings"
)

// LoadEnv 加载 .env 文件
func LoadEnv(envPath, sep string) (contracts.Fields, error) {
	envBytes, err := ioutil.ReadFile(envPath)
	if err != nil {
		return nil, err
	}

	fields := make(contracts.Fields)
	for _, line := range strings.Split(string(envBytes), "\n") {
		if strings.HasPrefix(line, "#") { // 跳过注释
			continue
		}
		values := strings.Split(line, sep)
		if len(values) > 1 {
			fields[values[0]] = strings.ReplaceAll(strings.Join(values[1:], sep), `"`, "")
		}
	}

	return fields, nil
}
