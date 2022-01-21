package tx

import (
	"github.com/goal-web/contracts"
	"github.com/jmoiron/sqlx"
)

type Tx struct {
	*sqlx.Tx
}

func (this *Tx) Query(query string, args ...interface{}) ([]contracts.Fields, error) {
	rows, err := this.Tx.Queryx(query, args...)

	if err != nil {
		return nil, err
	}

	results := make([]contracts.Fields, 0)

	for rows.Next() {
		row := make(map[string]interface{})
		err = rows.MapScan(row)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}

	return results, nil
}

func (this *Tx) Exec(query string, args ...interface{}) (contracts.Result, error) {
	return this.Tx.Exec(query, args...)
}
