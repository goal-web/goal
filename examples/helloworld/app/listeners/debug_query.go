package listeners

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/logs"
	"github.com/qbhy/goal/database/events"
)

type DebugQuery struct {
}

func (d DebugQuery) Handle(event contracts.Event) {
	if e, ok := event.(*events.QueryExecuted); ok {
		logs.WithFields(contracts.Fields{
			"sql":        e.Sql,
			"bindings":   e.Bindings,
			"connection": e.Connection,
		}).Info("sql executed")
	}
}
