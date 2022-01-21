package drivers

import (
	"github.com/goal-web/contracts"
	"github.com/jmoiron/sqlx"
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

func (this *base) Transaction(fn func(tx contracts.SqlExecutor) error) (err error) {
	sqlxTx, err := this.Begin()
	if err != nil {
		return exceptions2.BeginException{Exception: exceptions.WithError(err, nil)}
	}

	defer func() { // 处理 panic 情况
		if recoverErr := recover(); recoverErr != nil {
			rollbackErr := sqlxTx.Rollback()
			err = recoverErr.(error)
			if rollbackErr != nil {
				err = exceptions2.RollbackException{Exception: exceptions.WithPrevious(rollbackErr, nil, err)}
			} else {
				err = exceptions2.TransactionException{Exception: exceptions.WithError(err, nil)}
			}
		}
	}()

	err = fn(sqlxTx)

	if err != nil {
		rollbackErr := sqlxTx.Rollback()
		if rollbackErr != nil {
			return exceptions2.RollbackException{Exception: exceptions.WithPrevious(rollbackErr, nil, err)}
		}
		return exceptions2.TransactionException{Exception: exceptions.WithError(err, nil)}
	}

	return sqlxTx.Commit()
}

func (this *base) Exec(query string, args ...interface{}) (contracts.Result, error) {
	return this.DB.Exec(query, args...)
}
