package tests

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/validator"
	"github.com/qbhy/goal/validator/checkers"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DemoParam struct {
	Id string
}

func TestValidator(t *testing.T) {
	demoCheckers := contracts.Checkers{
		"id": {checkers.StringLength{1, 5}},
	}
	assert.True(t, validator.Validate(DemoParam{Id: "55555"}, demoCheckers).IsSuccessful())
	assert.False(t, validator.Validate(DemoParam{Id: "666666"}, demoCheckers).IsSuccessful())

	assert.True(t, validator.Validate(contracts.Fields{
		"id": "55555",
	}, demoCheckers).IsSuccessful())

	assert.False(t, validator.Validate(contracts.Fields{
		"id": "666666",
	}, demoCheckers).IsSuccessful())
}
