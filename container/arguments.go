package container

import (
	"github.com/qbhy/goal/supports/utils"
	"reflect"
)

type ArgumentsTypeMap map[string][]interface{}

func NewArgumentsTypeMap(args []interface{}) ArgumentsTypeMap {
	argsTypeMap := ArgumentsTypeMap{}
	for _, arg := range args {
		argTypeKey := utils.GetTypeKey(reflect.TypeOf(arg))
		argsTypeMap[argTypeKey] = append(argsTypeMap[argTypeKey], arg)
	}
	return argsTypeMap
}

func (this ArgumentsTypeMap) Pull(key string) (arg interface{}) {
	if item, exits := this[key]; exits && len(item) >= 1 {
		arg = item[0]
		this[key] = item[1:]
		return
	}
	return nil
}

// FindConvertibleArg 找到可转换的参数
func (this ArgumentsTypeMap) FindConvertibleArg(targetType reflect.Type) interface{} {
	for _, args := range this {
		for _, arg := range args {
			if reflect.TypeOf(arg).ConvertibleTo(targetType) {
				return arg
			}
		}
	}
	return nil
}
