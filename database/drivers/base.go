package drivers

import (
	"github.com/jmoiron/sqlx"
	"github.com/qbhy/goal/contracts"
	exceptions2 "github.com/qbhy/goal/database/exceptions"
	"github.com/qbhy/goal/database/tx"
	"github.com/qbhy/goal/exceptions"
)

type base struct {
	*sqlx.DB
}

func (this *base) Begin() (contracts.DBTx, error) {
	sqlxTx, err := this.DB.Beginx()
	if err != nil {
		return nil, err
	}
	return &tx.Tx{Tx: sqlxTx}, err
}

func (this *base) Transaction(fn func(tx contracts.SqlExecutor) error) error {
	sqlxTx, err := this.Begin()
	if err != nil {
		return exceptions2.BeginException{exceptions.WithError(err, nil)}
	}

	err = fn(sqlxTx)

	if err != nil {
		rollbackErr := sqlxTx.Rollback()
		if rollbackErr != nil {
			return exceptions2.RollbackException{exceptions.WithPrevious(rollbackErr, nil, err)}
		}
		return exceptions2.TransactionException{exceptions.WithError(err, nil)}
	}

	return nil
}

func (this *base) Exec(query string, args ...interface{}) (contracts.Result, error) {
	return this.DB.Exec(query, args...)
}
