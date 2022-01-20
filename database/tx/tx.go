package tx

import (
	"github.com/jmoiron/sqlx"
	"github.com/goal-web/contracts"
)

type Tx struct {
	*sqlx.Tx
}

func (this *Tx) Exec(query string, args ...interface{}) (contracts.Result, error) {
	return this.Tx.Exec(query, args...)
}
