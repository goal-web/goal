package table

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/exceptions"
	"github.com/goal-web/supports/utils"
)

func (this *Table) UpdateOrInsert(attributes contracts.Fields, values ...contracts.Fields) bool {
	this.WhereFields(attributes)
	sql, bindings := this.UpdateSql(attributes)
	result, err := this.getExecutor().Exec(sql, bindings...)
	if err != nil {
		panic(UpdateException{exceptions.WithError(err, contracts.Fields{
			"attributes": attributes,
			"values":     values,
		})})
	}
	num, _ := result.RowsAffected()
	if num > 0 {
		return true
	}
	if len(values) > 0 {
		utils.MergeFields(attributes, values[0])
	}
	return this.Insert(attributes)
}

func (this *Table) UpdateOrCreate(attributes contracts.Fields, values ...contracts.Fields) interface{} {
	this.WhereFields(attributes)
	sql, bindings := this.UpdateSql(attributes)
	result, err := this.getExecutor().Exec(sql, bindings...)
	if err != nil {
		panic(UpdateException{exceptions.WithError(err, contracts.Fields{
			"attributes": attributes,
			"values":     values,
		})})
	}
	num, _ := result.RowsAffected()
	if num > 0 {
		return true
	}
	if len(values) > 0 {
		utils.MergeFields(attributes, values[0])
	}
	return this.Insert(attributes)
}

func (this *Table) Update(fields contracts.Fields) int64 {
	sql, bindings := this.UpdateSql(fields)
	result, err := this.getExecutor().Exec(sql, bindings...)
	if err != nil {
		panic(UpdateException{exceptions.WithError(err, fields)})
	}
	num, err := result.RowsAffected()
	if err != nil {
		panic(UpdateException{exceptions.WithError(err, fields)})
	}
	return num
}
