package tx

import (
	"github.com/goal-web/collection"
	"github.com/goal-web/contracts"
	"github.com/jmoiron/sqlx"
	"github.com/goal-web/goal/database/events"
	"github.com/goal-web/goal/database/table"
)

type Tx struct {
	*sqlx.Tx
	events contracts.EventDispatcher
}

func New(tx *sqlx.Tx, events contracts.EventDispatcher) *Tx {
	return &Tx{
		Tx:     tx,
		events: events,
	}
}

func (this *Tx) Query(query string, args ...interface{}) (results contracts.Collection, err error) {
	defer func() {
		if err == nil {
			this.events.Dispatch(&events.QueryExecuted{Sql: query, Bindings: args})
		}
	}()
	rows, err := this.Tx.Query(query, args...)

	if err != nil {
		return nil, err
	}

	data, err := table.ParseRows(rows)
	results = collection.FromFieldsSlice(data)

	return
}

func (this *Tx) Exec(query string, args ...interface{}) (contracts.Result, error) {
	return this.Tx.Exec(query, args...)
}
