package table

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/querybuilder"
	"github.com/qbhy/goal/application"
	"github.com/qbhy/goal/exceptions"
)

type table struct {
	contracts.QueryBuilder
	connection contracts.DBConnection
	tx         contracts.DBTx

	primaryKey string
}

// Query 将使用默认 connection
func Query(name string) *table {
	return &table{
		querybuilder.NewQuery(name),
		application.Get("db").(contracts.DBConnection),
		nil,
		"id",
	}
}

// WithConnection 使用指定链接
func WithConnection(name string, connection interface{}) *table {
	return (&table{querybuilder.NewQuery(name), nil, nil, "id"}).
		SetConnection(connection)
}

// WithTX 使用TX
func WithTX(name string, tx contracts.DBTx) *table {
	return &table{querybuilder.NewQuery(name), nil, tx, "id"}
}

// SetConnection 参数要么是 contracts.DBConnection 要么是 string
func (this *table) SetConnection(connection interface{}) *table {
	if conn, ok := connection.(contracts.DBConnection); ok {
		this.connection = conn
	} else {
		this.connection = application.Get("db.factory").(contracts.DBFactory).Connection(connection.(string))
	}
	return this
}

// SetPrimaryKey 设置主键
func (this *table) SetPrimaryKey(name string) *table {
	this.primaryKey = name
	return this
}

// getExecutor 获取 sql 语句的执行者
func (this *table) getExecutor() contracts.SqlExecutor {
	if this.tx != nil {
		return this.tx
	}
	return this.connection
}

// SetTX 参数必须是 contracts.DBTx 实例
func (this *table) SetTX(tx interface{}) contracts.QueryBuilder {
	this.tx = tx.(contracts.DBTx)
	return this
}

func (this *table) Paginate(perPage int64, current ...int64) (interface{}, int64) {
	//TODO implement me
	panic("implement me")
}

func (this *table) Create(fields contracts.Fields) interface{} {
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

func (this *table) Insert(values ...contracts.Fields) interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *table) Delete() int64 {
	//TODO implement me
	panic("implement me")
}

func (this *table) Update(fields contracts.Fields) int64 {
	//TODO implement me
	panic("implement me")
}

func (this *table) Get() interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *table) Find(key interface{}) interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *table) First() interface{} {
	//TODO implement me
	panic("implement me")
}
