package tests

import (
	"fmt"
	"reflect"
	"testing"
)

type DemoArg struct {
}

func TestReflectFunc(t *testing.T) {
	fn := func(param []DemoArg) interface{} {
		return nil
	}

	fnType := reflect.TypeOf(fn)
	argNum := fnType.NumIn()
	fmt.Println("参数个数:", fnType.NumIn())

	for i := 0; i < argNum; i++ {
		arg := fnType.In(i)
		fmt.Println(arg.Name(), arg.Kind())
	}

	fmt.Println(fnType)
}
