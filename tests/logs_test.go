package tests

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/exceptions"
	"github.com/goal-web/supports/logs"
	"github.com/qbhy/parallel"
	"testing"
)

func TestLogger(t *testing.T) {
	logs.WithError(errors.New("报错了")).Info("info数据")

	logs.WithFields(contracts.Fields{"id": "1"}).Warn("info数据")

	logs.WithException(exceptions.New("报错啦", contracts.Fields{"id": 1, "name": "qbhy"})).Info("info数据")
}

func TestWithField(t *testing.T) {
	row := parallel.NewParallel(50)

	logs.WithError(errors.New("报错了")).WithField("field1", "1").Info("info数据")

	row.Add(func() interface{} {
		logs.WithError(errors.New("协程里面报错了")).WithField("field1", "1").Info("info数据")
		return nil
	})

	row.Run()
}
