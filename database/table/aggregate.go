package table

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal/exceptions"
)

func (this *table) Count(columns ...string) int64 {
	sql, bindings := this.WithCount(columns...).SelectSql()
	var num int64
	err := this.getExecutor().Get(&num, sql, bindings...)
	if err != nil {
		panic(SelectException{exceptions.WithError(err, contracts.Fields{
			"columns":  columns,
			"sql":      sql,
			"bindings": bindings,
		})})
	}
	return num
}

func (this *table) Avg(column string, as ...string) int64 {
	sql, bindings := this.WithAvg(column, as...).SelectSql()
	var num int64
	err := this.getExecutor().Get(&num, sql, bindings...)
	if err != nil {
		panic(SelectException{exceptions.WithError(err, contracts.Fields{
			"column":   column,
			"sql":      sql,
			"bindings": bindings,
		})})
	}
	return num
}

func (this *table) Sum(column string, as ...string) int64 {
	sql, bindings := this.WithSum(column, as...).SelectSql()
	var num int64
	err := this.getExecutor().Get(&num, sql, bindings...)
	if err != nil {
		panic(SelectException{exceptions.WithError(err, contracts.Fields{
			"column":   column,
			"sql":      sql,
			"bindings": bindings,
		})})
	}
	return num
}

func (this *table) Max(column string, as ...string) int64 {
	sql, bindings := this.WithMax(column, as...).SelectSql()
	var num int64
	err := this.getExecutor().Get(&num, sql, bindings...)
	if err != nil {
		panic(SelectException{exceptions.WithError(err, contracts.Fields{
			"column":   column,
			"sql":      sql,
			"bindings": bindings,
		})})
	}
	return num
}

func (this *table) Min(column string, as ...string) int64 {
	sql, bindings := this.WithMin(column, as...).SelectSql()
	var num int64
	err := this.getExecutor().Get(&num, sql, bindings...)
	if err != nil {
		panic(SelectException{exceptions.WithError(err, contracts.Fields{
			"column":   column,
			"sql":      sql,
			"bindings": bindings,
		})})
	}
	return num
}
