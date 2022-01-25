package table

import (
	"github.com/goal-web/application"
	"github.com/goal-web/contracts"
	"github.com/goal-web/querybuilder"
	"github.com/goal-web/supports/class"
	"github.com/goal-web/supports/exceptions"
)

type Table struct {
	contracts.QueryBuilder
	executor contracts.SqlExecutor

	table      string
	primaryKey string
	class      *class.Class
}

func getTable(name string) *Table {
	builder := querybuilder.NewQuery(name)
	instance := &Table{
		QueryBuilder: builder,
		primaryKey:   "id",
		table:        name,
	}
	builder.Bind(instance)
	return instance
}

// Query 将使用默认 connection
func Query(name string) *Table {
	return getTable(name).SetConnection(application.Get("db").(contracts.DBConnection))
}

// WithConnection 使用指定链接
func WithConnection(name string, connection interface{}) *Table {
	return getTable(name).SetConnection(connection)
}

// WithTX 使用TX
func WithTX(name string, tx contracts.DBTx) contracts.QueryBuilder {
	return getTable(name).SetExecutor(tx)
}

// SetConnection 参数要么是 contracts.DBConnection 要么是 string
func (this *Table) SetConnection(connection interface{}) *Table {
	if conn, ok := connection.(contracts.DBConnection); ok {
		this.executor = conn
	} else {
		this.executor = application.Get("db.factory").(contracts.DBFactory).Connection(connection.(string))
	}
	return this
}

// SetClass 设置类
func (this *Table) SetClass(class *class.Class) *Table {
	this.class = class
	return this
}

// SetPrimaryKey 设置主键
func (this *Table) SetPrimaryKey(name string) *Table {
	this.primaryKey = name
	return this
}

// getExecutor 获取 sql 语句的执行者
func (this *Table) getExecutor() contracts.SqlExecutor {
	return this.executor
}

// SetExecutor 参数必须是 contracts.DBTx 实例
func (this *Table) SetExecutor(executor contracts.SqlExecutor) contracts.QueryBuilder {
	this.executor = executor
	return this
}

func (this *Table) Delete() int64 {
	sql, bindings := this.DeleteSql()
	result, err := this.getExecutor().Exec(sql, bindings...)
	if err != nil {
		panic(DeleteException{exceptions.WithError(err, contracts.Fields{"sql": sql, "bindings": bindings})})
	}
	num, err := result.RowsAffected()
	if err != nil {
		panic(DeleteException{exceptions.WithError(err, contracts.Fields{"sql": sql, "bindings": bindings})})
	}
	return num
}
