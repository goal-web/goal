package tests

import (
	"github.com/qbhy/goal/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArgumentsTypeMap(t *testing.T) {
	args := container.NewArgumentsTypeMap([]interface{}{"啦啦啦", DemoParam{Id: "111"}})
	assert.True(t, args.Pull("string") == "啦啦啦")

	args = container.NewArgumentsTypeMap([]interface{}{})
	assert.True(t, args.Pull("string") == nil)
}

func TestBaseContainer(t *testing.T) {
	app := container.New()

	app.Instance("a", "a")
	assert.True(t, app.Get("a") == "a")

	app.Provide(func() DemoParam {
		return DemoParam{Id: "测试一下"}
	})

	app.Call(func(param DemoParam) {
		assert.True(t, param.Id == "测试一下")
	})

}

func TestContainer(t *testing.T) {
	app := container.New()

	app.Provide(func() DemoParam {
		return DemoParam{Id: "没有外部参数的话，从容器中获取"}
	})

	fn := func(param DemoParam) string {
		return param.Id
	}

	// 自己传参
	assert.True(t, app.Call(fn, DemoParam{Id: "优先使用外部参数"})[0] == "优先使用外部参数")

	// 不传参，使用容器中的实例
	assert.True(t, app.Call(fn)[0] == "没有外部参数的话，从容器中获取")

}
