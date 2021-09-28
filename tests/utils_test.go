package tests

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeFields(t *testing.T) {
	fields1 := contracts.Fields{"a": "a"}
	utils.MergeFields(fields1, map[string]interface{}{
		"a": "b",
	})

	assert.True(t, fields1["a"] == "b")
}
