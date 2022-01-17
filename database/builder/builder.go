package builder

import (
	"fmt"
	"strings"
)

type Callback func(*Builder) *Builder
type BuilderProvider func() *Builder
type whereFunc func(*Builder)

type Builder struct {
	table   string
	fields  []string
	wheres  *Wheres
	orderBy OrderByFields
	groupBy GroupBy
}

func NewQueryBuilder(table string) *Builder {
	return &Builder{
		table:   table,
		fields:  []string{"*"},
		orderBy: OrderByFields{},
		groupBy: GroupBy{},
		wheres: &Wheres{
			wheres:    map[string][]*Where{},
			subWheres: map[string][]*Wheres{},
		},
	}
}

func (this *Builder) getWheres() *Wheres {
	return this.wheres
}

func (this *Builder) WhereFunc(callback whereFunc, whereType ...string) *Builder {
	subBuilder := NewQueryBuilder("")
	callback(subBuilder)
	if len(whereType) == 0 {
		this.wheres.subWheres[and] = append(this.wheres.subWheres[and], subBuilder.getWheres())
	} else {
		this.wheres.subWheres[whereType[0]] = append(this.wheres.subWheres[whereType[0]], subBuilder.getWheres())
	}
	return this
}

func (this *Builder) OrWhereFunc(callback whereFunc) *Builder {
	subBuilder := NewQueryBuilder("")
	callback(subBuilder)
	this.wheres.subWheres[or] = append(this.wheres.subWheres[or], subBuilder.getWheres())
	return this
}

func (this *Builder) Where(field string, args ...interface{}) *Builder {
	var (
		arg       interface{}
		condition = "="
		whereType = and
	)
	switch len(args) {
	case 1:
		arg = args[0]
	case 2:
		condition = args[0].(string)
		arg = args[1]
	case 3:
		condition = args[0].(string)
		arg = args[1]
		whereType = args[2].(string)
	}

	this.wheres.wheres[whereType] = append(this.wheres.wheres[whereType], &Where{
		field:     field,
		condition: condition,
		arg:       arg,
	})

	return this
}

func (this *Builder) WhereIn(field string, args interface{}) *Builder {
	return this.Where(field, "in", args)
}

func (this *Builder) OrWhereIn(field string, args interface{}) *Builder {
	return this.OrWhere(field, "in", args)
}

func (this *Builder) WhereNotIn(field string, args interface{}) *Builder {
	return this.Where(field, "not in", args)
}

func (this *Builder) OrWhereNotIn(field string, args interface{}) *Builder {
	return this.OrWhere(field, "not in", args)
}

func (this *Builder) OrWhere(field string, args ...interface{}) *Builder {
	var (
		arg       interface{}
		condition = "="
	)
	switch len(args) {
	case 1:
		arg = args[0]
	case 2:
		condition = args[0].(string)
		arg = args[1]
	default:
		condition = args[0].(string)
		arg = args[1]
	}

	this.wheres.wheres[or] = append(this.wheres.wheres[or], &Where{
		field:     field,
		condition: condition,
		arg:       arg,
	})
	return this
}

func (this *Builder) WhereIsNull(field string, whereType ...string) *Builder {
	if len(whereType) == 0 {
		return this.Where(field, "", "is null", and)
	}
	return this.Where(field, "", "is null", whereType[0])
}

func (this *Builder) OrWhereIsNull(field string) *Builder {
	return this.OrWhere(field, "", "is null")
}

func (this *Builder) OrWhereNotNull(field string) *Builder {
	return this.OrWhere(field, "", "is not null")
}

func (this *Builder) WhereNotNull(field string, whereType ...string) *Builder {
	if len(whereType) == 0 {
		return this.Where(field, "", "is not null", and)
	}
	return this.Where(field, "", "is not null", whereType[0])
}

func (this *Builder) From(table string, as ...string) *Builder {
	if len(as) == 0 {
		this.table = table
	} else {
		this.table = fmt.Sprintf("%s as %s", table, as[0])
	}
	return this
}

func (this *Builder) FromSub(callback BuilderProvider, as string) *Builder {
	this.table = fmt.Sprintf("(%s) as %s", callback().ToSql(), as)
	return this
}

func (this *Builder) Select(field string, fields ...string) *Builder {
	this.fields = append(this.fields, append(fields, field)...)
	return this
}

func (this *Builder) When(condition bool, callback Callback, elseCallback ...Callback) *Builder {
	if condition {
		return callback(this)
	} else if len(elseCallback) > 0 {
		return elseCallback[0](this)
	}
	return this
}

func (this *Builder) AddSelect(fields ...string) *Builder {
	this.fields = append(this.fields, fields...)
	return this
}

func (this *Builder) OrderBy(field string, columnOrderType ...orderType) *Builder {
	if len(columnOrderType) > 0 {
		this.orderBy = append(this.orderBy, OrderBy{
			field:          field,
			fieldOrderType: columnOrderType[0],
		})
	} else {
		this.orderBy = append(this.orderBy, OrderBy{
			field:          field,
			fieldOrderType: ASC,
		})
	}

	return this
}

func (this *Builder) GroupBy(columns ...string) *Builder {
	this.groupBy = append(this.groupBy, columns...)

	return this
}

func (this *Builder) OrderByDesc(field string) *Builder {
	this.orderBy = append(this.orderBy, OrderBy{
		field:          field,
		fieldOrderType: DESC,
	})
	return this
}

func (this *Builder) ToSql() string {
	sql := fmt.Sprintf("select %s from %s", strings.Join(this.fields, ","), this.table)

	if !this.wheres.IsEmpty() {
		sql = fmt.Sprintf("%s where %s", sql, this.wheres.String())
	}

	if !this.groupBy.IsEmpty() {
		sql = fmt.Sprintf("%s group by %s", sql, this.groupBy.String())
	}

	if !this.orderBy.IsEmpty() {
		sql = fmt.Sprintf("%s order by %s", sql, this.orderBy.String())
	}

	return sql
}
