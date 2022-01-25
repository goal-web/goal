package table

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/exceptions"
	"github.com/goal-web/supports/utils"
)

func (this *Table) Create(fields contracts.Fields) interface{} {
	sql, bindings := this.CreateSql(fields)
	result, err := this.getExecutor().Exec(sql, bindings...)
	if err != nil {
		panic(CreateException{exceptions.WithError(err, fields)})
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(CreateException{exceptions.WithError(err, fields)})
	}

	if _, existsPrimaryKey := fields[this.primaryKey]; !existsPrimaryKey {
		fields[this.primaryKey] = id
	}
	return fields
}

func (this *Table) Insert(values ...contracts.Fields) bool {
	sql, bindings := this.InsertSql(values)
	result, err := this.getExecutor().Exec(sql, bindings...)

	if err != nil {
		panic(InsertException{exceptions.WithError(err, contracts.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		panic(InsertException{exceptions.WithError(err, contracts.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	return rowsAffected > 0
}

func (this *Table) InsertGetId(values ...contracts.Fields) int64 {
	sql, bindings := this.InsertSql(values)
	result, err := this.getExecutor().Exec(sql, bindings...)

	if err != nil {
		panic(InsertException{exceptions.WithError(err, contracts.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	id, err := result.LastInsertId()

	if err != nil {
		panic(InsertException{exceptions.WithError(err, contracts.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	return id
}

func (this *Table) InsertOrIgnore(values ...contracts.Fields) int64 {
	sql, bindings := this.InsertIgnoreSql(values)
	result, err := this.getExecutor().Exec(sql, bindings...)

	if err != nil {
		panic(InsertException{exceptions.WithError(err, contracts.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		panic(InsertException{exceptions.WithError(err, contracts.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	return rowsAffected
}

func (this *Table) InsertOrReplace(values ...contracts.Fields) int64 {
	sql, bindings := this.InsertReplaceSql(values)
	result, err := this.getExecutor().Exec(sql, bindings...)

	if err != nil {
		panic(InsertException{exceptions.WithError(err, contracts.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		panic(InsertException{exceptions.WithError(err, contracts.Fields{
			"values": values,
			"sql":    sql,
		})})
	}

	return rowsAffected
}

func (this *Table) FirstOrCreate(values ...contracts.Fields) interface{} {
	var attributes contracts.Fields
	argsLen := len(values)
	if argsLen > 0 {
		for field, value := range values[0] {
			attributes[field] = value
			this.Where(field, value)
		}
		if result := this.First(); result != nil {
			return result
		} else if argsLen > 1 {
			utils.MergeFields(attributes, values[1])
		}
		return this.Create(attributes)
	}

	return nil
}
