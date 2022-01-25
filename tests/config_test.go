package tests

import (
	"github.com/goal-web/config"
	"github.com/goal-web/contracts"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBaseConfig(t *testing.T) {

	conf := config.NewConfig("testing")

	conf.Load(config.FieldsProvider{Data: map[string]interface{}{
		"a": "a",
	}})

	path, _ := os.Getwd()
	conf.Load(config.NewEnv([]string{path}, "="))

	// 测试从环境变量获取配置
	assert.Nil(t, os.Setenv("app.name", "ggboy"))
	assert.True(t, conf.GetString("app.name") == "ggboy")

	// 测试 env 覆盖配置文件配置
	assert.Nil(t, os.Setenv("int", "new_int"))
	assert.True(t, conf.GetString("int") == "new_int")
	conf.Unset("int")
	assert.Nil(t, os.Unsetenv("int"))

	conf.Merge("sub", config.WithFields(contracts.Fields{"a": "sub a"}))

	// 测试 get
	assert.True(t, conf.Get("a", "default a") == "a")
	assert.True(t, conf.Get("b", "default b") == "default b")

	// 测试多层 get
	assert.True(t, conf.Get("sub.a", "default a") == "sub a")
	assert.True(t, conf.Get("sub.b", "default sub b") == "default sub b")

	// 测试默认值
	assert.True(t, conf.Get("c") == nil)

	// 测试 set
	conf.Set("b", "b")
	assert.True(t, conf.Get("b", "default b") == "b")
	conf.Set("sub.b", "sub b")
	assert.True(t, conf.Get("sub.b", "default sub b") == "sub b")

	// 测试获取 fields 实例
	assert.True(t, conf.GetFields("do not exists fields")["tests"] == nil)
	conf.Set("field", contracts.Fields{"tests": "tests"})
	assert.True(t, conf.GetFields("field")["tests"] == "tests")

	// 测试获取 子配置 实例
	assert.True(t, conf.GetConfig("sub_config") == nil)
	assert.True(t, conf.GetConfig("sub").Get("a") == "sub a")

	// 测试基础类型 getter
	conf.Set("int", 10086)
	conf.Set("float", 100.86)
	conf.Set("string", "10086")
	conf.Set("bool", true)
	assert.True(t, conf.GetInt("int") == 10086)
	assert.True(t, conf.GetFloat("float") == 100.86)
	assert.True(t, conf.GetString("string") == "10086")
	assert.True(t, conf.GetBool("bool") == true)

	// 测试基础类型转换
	assert.True(t, conf.GetFloat("int") == 10086)
	assert.True(t, conf.GetInt("float") == 100)
	assert.True(t, conf.GetFloat("string") == 10086)
	assert.True(t, conf.GetInt("string") == 10086)

	conf.Set("string_true", "true")
	conf.Set("string_false", "false")
	assert.True(t, conf.GetBool("string_true"))
	assert.False(t, conf.GetBool("string_false"))

	conf.Set("int_true", 1)
	assert.True(t, conf.GetBool("int_true"))

	conf.Set("int_false", 0)
	assert.False(t, conf.GetBool("int_false"))

	conf.Set("float_true", 0.5)
	assert.True(t, conf.GetBool("int_true"))

	conf.Set("float_false", 0.49)
	assert.False(t, conf.GetBool("float_false"))
}

func TestDotEnv(t *testing.T) {
	path, _ := os.Getwd()

	env := config.NewEnv([]string{path + "/..", path}, "=")

	fields := env.Fields()

	assert.True(t, len(env.GetFields("redis")) > 0)

	assert.True(t, fields["testing:name"] == "testing goal")

	assert.Nil(t, os.Setenv("app.key", "goal"))
	assert.True(t, env.GetString("app.key") == "goal")
}
