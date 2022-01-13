package tests

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMergeFields(t *testing.T) {
	fields1 := contracts.Fields{"a": "a"}
	utils.MergeFields(fields1, map[string]interface{}{
		"a":          "b",
		"int":        1,
		"bool":       true,
		"stringBool": "(true)",
	})

	assert.True(t, fields1["a"] == "b")
	assert.True(t, utils.GetStringField(fields1, "a") == "b")
	assert.True(t, utils.GetInt64Field(fields1, "int") == 1)
	assert.True(t, utils.GetBoolField(fields1, "bool"))
	assert.True(t, utils.GetBoolField(fields1, "stringBool"))

}

func TestRandStr(t *testing.T) {
	fmt.Println(utils.RandStr(50))
}

func TestSubString(t *testing.T) {
	subStr := utils.SubString("123456789", 1, 5)
	fmt.Println(subStr)

	assert.True(t, subStr == "23456")
	zhSubStr := utils.SubString("一二三四五六七八九", 1, 5)

	fmt.Println(zhSubStr)
	assert.True(t, zhSubStr == "二三四五六")

	overflowStr := utils.SubString("一二三四五六七八九", 1, 1000)
	fmt.Println(overflowStr)

	sufferStr := utils.SubString("一二三四五六七八九", 1, 0)
	fmt.Println(sufferStr)

	assert.True(t, overflowStr == sufferStr)

	midStr := utils.SubString("一二三四五六七八九", 1, -1)
	fmt.Println("midStr:", midStr)
	assert.True(t, midStr == "二三四五六七八")

}

func TestCarbonDay(t *testing.T) {
	fmt.Println(carbon.Now().Day(), carbon.Now().EndOfMonth().Day())
}

func TestOsHostname(t *testing.T) {
	fmt.Println(os.Hostname())
	fmt.Println(os.UserHomeDir())
}
