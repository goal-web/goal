package table

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/exceptions"
)

func (this *table) Get() interface{} {
	sql, bindings := this.SelectSql()
	rows, err := this.getExecutor().Query(sql, bindings...)
	if err != nil {
		panic(SelectException{exceptions.WithError(err, contracts.Fields{"sql": sql, "bindings": bindings})})
	}
	return rows
}

func (this *table) Find(key interface{}) interface{} {
	return this.Where(this.primaryKey, key).First()
}

func (this *table) First() interface{} {
	if results := this.Take(1).Get().([]contracts.Fields); len(results) > 0 {
		return results[0]
	}
	return nil
}

func (this *table) FirstOrFail() interface{} {
	if result := this.First(); result != nil {
		return result
	}
	panic(NotFoundException{exceptions.WithError(errors.New("未找到"), contracts.Fields{})})
}
