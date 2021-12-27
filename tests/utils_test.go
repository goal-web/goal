package tests

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"github.com/stretchr/testify/assert"
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
