package tests

import (
	"fmt"
	"github.com/qbhy/goal/container"
	"github.com/qbhy/goal/utils"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestArgumentsTypeMap(t *testing.T) {
	args := container.NewArgumentsTypeMap([]interface{}{"啦啦啦", DemoParam{Id: "111"}})
	str := args.Pull("string")
	fmt.Println(str)
	assert.True(t, str == "啦啦啦")

	args = container.NewArgumentsTypeMap([]interface{}{})
	assert.True(t, args.Pull("string") == nil)
}

func TestBaseContainer(t *testing.T) {
	app := container.New()

	app.Instance("a", "a")
	assert.True(t, app.HasBound("a"))
	assert.True(t, app.Get("a") == "a")

	app.Alias("a", "A")

	assert.True(t, app.Get("A") == "a")
	assert.True(t, app.HasBound("A"))

	app.Provide(func() DemoParam {
		return DemoParam{Id: "测试一下"}
	})

	assert.True(t, app.Get(utils.GetTypeKey(reflect.TypeOf(DemoParam{}))).(DemoParam).Id == "测试一下")

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

type DemoStruct struct {
	Param  DemoParam `di:""`       // 注入对应类型的实例
	Config string    `di:"config"` // 注入指定 key 的实例
}

func TestContainerMake(t *testing.T) {
	app := container.New()

	app.Instance("config", "通过容器设置的配置")

	app.Provide(func() DemoParam {
		return DemoParam{Id: "没有外部参数的话，从容器中获取"}
	})

	demo := &DemoStruct{}

	app.DI(demo)

	fmt.Println(demo)
}

func TestReflectValue(t *testing.T) {
	var b interface{}
	b = struct {
		Name string
	}{}

	// v is the interface{}
	v := reflect.ValueOf(&b).Elem()

	// Allocate a temporary variable with type of the struct.
	//    v.Elem() is the vale contained in the interface.
	tmp := reflect.New(v.Elem().Type()).Elem()

	// Copy the struct value contained in interface to
	// the temporary variable.
	tmp.Set(v.Elem())

	// Set the field.
	tmp.FieldByName("Name").SetString("Hello")

	// Set the interface to the modified struct value.
	v.Set(tmp)

	fmt.Printf("%#v\n", b)
}
