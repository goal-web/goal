package config

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal/supports"
	"github.com/qbhy/goal/supports/utils"
	"os"
	"path/filepath"
)

type envProvider struct {
	supports.BaseFields
	Paths  []string
	Sep    string
	fields contracts.Fields
}

func NewEnv(paths []string, sep string) contracts.Env {
	provider := &envProvider{
		BaseFields: supports.BaseFields{Getter: func(key string) interface{} {
			return os.Getenv(key)
		}},
		Paths:  paths,
		Sep:    sep,
		fields: nil,
	}

	provider.BaseFields.FieldsProvider = provider
	return provider
}

func (this *envProvider) Fields() contracts.Fields {
	if this.fields != nil {
		return this.fields
	}

	this.fields = this.Load()

	return this.fields
}

func (this *envProvider) Load() contracts.Fields {
	var (
		files  []string
		fields = make(contracts.Fields)
	)
	for _, path := range this.Paths {
		tmpFiles, _ := filepath.Glob(path + "/*.env")
		files = append(files, tmpFiles...)
	}

	for _, file := range files {
		tempFields, _ := utils.LoadEnv(file, utils.StringOr(this.Sep, "="))
		if tempFields["env"] != nil { // 加载成功并且设置了 env
			newFields := make(contracts.Fields)
			envValue := tempFields["env"].(string)
			for key, field := range tempFields {
				if key != "env" {
					newFields[fmt.Sprintf("%s:%s", envValue, key)] = field
				}
			}
			tempFields = newFields
		}
		utils.MergeFields(fields, tempFields)
	}

	return fields
}
