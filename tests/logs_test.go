package tests

import (
	"errors"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/logs"
	"testing"
)

func TestLogger(t *testing.T) {
	logs.WithError(errors.New("报错了")).Info("info数据")

	logs.WithFields(contracts.Fields{"id": "1"}).Warn("info数据")

	logs.WithException(exceptions.New("报错啦", contracts.Fields{"id": 1, "name": "qbhy"})).Info("info数据")
}
