package config

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"path/filepath"
)

type EnvFieldsProvider struct {
	Paths []string
	Sep   string
}

func (e EnvFieldsProvider) Get() contracts.Fields {
	var (
		files  []string
		fields = make(contracts.Fields)
	)
	for _, path := range e.Paths {
		tmpFiles, _ := filepath.Glob(path + "/*.env")
		files = append(files, tmpFiles...)
	}

	for _, file := range files {
		tempFields, _ := utils.LoadEnv(file, utils.StringOr(e.Sep, "="))
		if tempFields["env"] != nil { // 加载成功并且设置了 env
			newFields := make(contracts.Fields)
			env := tempFields["env"].(string)
			for key, field := range tempFields {
				if key != "env" {
					newFields[fmt.Sprintf("%s:%s", env, key)] = field
				}
			}
			tempFields = newFields
		}
		utils.MergeFields(fields, tempFields)
	}

	return fields
}
