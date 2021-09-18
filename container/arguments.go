package container

import (
	"github.com/qbhy/goal/utils"
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
