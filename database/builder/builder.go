package builder

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/utils"
	"strings"
)

type Callback func(*Builder) *Builder
type Provider func() *Builder
type whereFunc func(*Builder)

type bindingType string

type Builder struct {
	distinct bool
	table    string
	fields   []string
	wheres   *Wheres
	orderBy  OrderByFields
	groupBy  GroupBy
	joins    Joins
	unions   Unions
	having   *Wheres
	bindings map[bindingType][]interface{}
}

const (
	selectBinding  bindingType = "select"
	fromBinding    bindingType = "from"
	joinBinding    bindingType = "join"
	whereBinding   bindingType = "where"
	groupByBinding bindingType = "groupBy"
	havingBinding  bindingType = "having"
	orderBinding   bindingType = "order"
	unionBinding   bindingType = "union"
)

func NewQuery(table string) *Builder {
	return &Builder{
		table:    table,
		fields:   []string{"*"},
		orderBy:  OrderByFields{},
		bindings: map[bindingType][]interface{}{},
		joins:    Joins{},
		unions:   Unions{},
		groupBy:  GroupBy{},
		wheres: &Wheres{
			wheres:    map[whereJoinType][]*Where{},
			subWheres: map[whereJoinType][]*Wheres{},
		},
		having: &Wheres{
			wheres:    map[whereJoinType][]*Where{},
			subWheres: map[whereJoinType][]*Wheres{},
		},
	}
}

func FromSub(callback Provider, as string) *Builder {
	return NewQuery("").FromSub(callback, as)
}

func (this *Builder) getWheres() *Wheres {
	return this.wheres
}
func (this *Builder) Union(builder *Builder, unionType ...unionJoinType) *Builder {
	if builder != nil {
		if len(unionType) > 0 {
			this.unions[unionType[0]] = append(this.unions[unionType[0]], builder)
		} else {
			this.unions[Union] = append(this.unions[Union], builder)
		}
	}

	return this.addBinding(unionBinding, builder.GetBindings()...)
}

func (this *Builder) UnionAll(builder *Builder) *Builder {
	return this.Union(builder, UnionAll)
}

func (this *Builder) UnionByProvider(builder Provider, unionType ...unionJoinType) *Builder {
	return this.Union(builder(), unionType...)
}

func (this *Builder) UnionAllByProvider(builder Provider) *Builder {
	return this.Union(builder(), UnionAll)
}

func (this *Builder) WhereFunc(callback whereFunc, whereType ...whereJoinType) *Builder {
	subBuilder := &Builder{
		wheres: &Wheres{
			wheres:    map[whereJoinType][]*Where{},
			subWheres: map[whereJoinType][]*Wheres{},
		},
		bindings: map[bindingType][]interface{}{},
	}
	callback(subBuilder)
	if len(whereType) == 0 {
		this.wheres.subWheres[And] = append(this.wheres.subWheres[And], subBuilder.getWheres())
	} else {
		this.wheres.subWheres[whereType[0]] = append(this.wheres.subWheres[whereType[0]], subBuilder.getWheres())
	}
	this.addBinding(whereBinding, subBuilder.GetBindings()...)
	return this
}

func (this *Builder) OrWhereFunc(callback whereFunc) *Builder {
	return this.WhereFunc(callback, Or)
}

func (this *Builder) Where(field string, args ...interface{}) *Builder {
	var (
		arg       interface{}
		condition = "="
		whereType = And
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
		whereType = args[2].(whereJoinType)
	}

	raw, bindings := this.prepareArgs(condition, arg)

	this.wheres.wheres[whereType] = append(this.wheres.wheres[whereType], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})

	return this.addBinding(whereBinding, bindings...)
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
	raw, bindings := this.prepareArgs(condition, arg)

	this.wheres.wheres[Or] = append(this.wheres.wheres[Or], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})
	return this.addBinding(whereBinding, bindings...)
}

func (this *Builder) Having(field string, args ...interface{}) *Builder {
	var (
		arg       interface{}
		condition = "="
		whereType = And
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
		whereType = args[2].(whereJoinType)
	}

	raw, bindings := this.prepareArgs(condition, arg)

	this.having.wheres[whereType] = append(this.having.wheres[whereType], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})

	return this.addBinding(havingBinding, bindings...)
}

func (this *Builder) OrHaving(field string, args ...interface{}) *Builder {
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
	raw, bindings := this.prepareArgs(condition, arg)

	this.having.wheres[Or] = append(this.having.wheres[Or], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})
	return this.addBinding(havingBinding, bindings...)
}

func (this *Builder) prepareArgs(condition string, args interface{}) (raw string, bindings []interface{}) {
	condition = strings.ToLower(condition)
	switch condition {
	case "in", "not in", "between", "not between":
		isInGrammar := strings.Contains(condition, "in")
		joinSymbol := utils.IfString(isInGrammar, ",", " and ")
		var stringArg string
		switch arg := args.(type) {
		case string:
			stringArg = arg
		case fmt.Stringer:
			stringArg = arg.String()
		case []string:
			stringArg = strings.Join(arg, joinSymbol)
		case []int:
			stringArg = utils.JoinIntArray(arg, joinSymbol)
		case []int64:
			stringArg = utils.JoinInt64Array(arg, joinSymbol)
		case []float64:
			stringArg = utils.JoinFloat64Array(arg, joinSymbol)
		case []float32:
			stringArg = utils.JoinFloatArray(arg, joinSymbol)
		case []interface{}:
			bindings = arg
			return
		default:
			panic(exceptions.WithError(errors.New("不支持的参数类型"), contracts.Fields{
				"arg":       arg,
				"condition": condition,
			}))
		}
		bindings = utils.StringArray2InterfaceArray(strings.Split(stringArg, joinSymbol))
		if isInGrammar {
			raw = fmt.Sprintf("(%s)", strings.Join(utils.MakeSymbolArray("?", len(bindings)), ","))
		} else {
			raw = "? and ?"
		}
	case "is", "is not":
		raw = utils.ConvertToString(args, "")
	default:
		raw = "?"
		bindings = append(bindings, utils.ConvertToString(args, ""))
	}

	return
}

func (this *Builder) WhereIn(field string, args interface{}) *Builder {
	return this.Where(field, "in", args)
}

func (this *Builder) addBinding(bindType bindingType, bindings ...interface{}) *Builder {
	this.bindings[bindType] = append(this.bindings[bindType], bindings...)
	return this
}

func (this *Builder) GetBindings() (results []interface{}) {
	for _, bindings := range this.bindings {
		results = append(results, bindings...)
	}
	return
}

func (this *Builder) Distinct() *Builder {
	this.distinct = true
	return this
}

func (this *Builder) Join(table string, first, condition, second string, joins ...joinType) *Builder {
	join := InnerJoin
	if len(joins) > 0 {
		join = joins[0]
	}
	this.joins = append(this.joins, Join{table, join, &Wheres{wheres: map[whereJoinType][]*Where{
		And: {&Where{
			field:     first,
			condition: condition,
			arg:       second,
		}},
	}}})

	return this
}

func (this *Builder) JoinSub(provider Provider, as, first, condition, second string, joins ...joinType) *Builder {
	join := InnerJoin
	if len(joins) > 0 {
		join = joins[0]
	}
	subBuilder := provider()
	this.joins = append(this.joins, Join{fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as), join, &Wheres{wheres: map[whereJoinType][]*Where{
		And: {&Where{
			field:     first,
			condition: condition,
			arg:       second,
		}},
	}}})

	return this.addBinding(joinBinding, subBuilder.GetBindings()...)
}

func (this *Builder) FullJoin(table string, first, condition, second string) *Builder {
	return this.Join(table, first, condition, second, FullJoin)
}
func (this *Builder) FullOutJoin(table string, first, condition, second string) *Builder {
	return this.Join(table, first, condition, second, FullOutJoin)
}

func (this *Builder) LeftJoin(table string, first, condition, second string) *Builder {
	return this.Join(table, first, condition, second, LeftJoin)
}

func (this *Builder) RightJoin(table string, first, condition, second string) *Builder {
	return this.Join(table, first, condition, second, RightJoin)
}

func (this *Builder) OrWhereIn(field string, args interface{}) *Builder {
	return this.OrWhere(field, "in", args)
}

func (this *Builder) WhereBetween(field string, args interface{}, whereType ...whereJoinType) *Builder {
	if len(whereType) > 0 {
		return this.Where(field, "between", args, whereType[0])
	}

	return this.Where(field, "between", args)
}

func (this *Builder) OrWhereBetween(field string, args interface{}) *Builder {
	return this.OrWhere(field, "between", args)
}

func (this *Builder) WhereNotBetween(field string, args interface{}, whereType ...whereJoinType) *Builder {
	if len(whereType) > 0 {
		return this.Where(field, "not between", args, whereType[0])
	}

	return this.Where(field, "not between", args)
}

func (this *Builder) OrWhereNotBetween(field string, args interface{}) *Builder {
	return this.OrWhere(field, "not between", args)
}

func (this *Builder) WhereNotIn(field string, args interface{}) *Builder {
	return this.Where(field, "not in", args)
}

func (this *Builder) OrWhereNotIn(field string, args interface{}) *Builder {
	return this.OrWhere(field, "not in", args)
}

func (this *Builder) WhereIsNull(field string, whereType ...string) *Builder {
	if len(whereType) == 0 {
		return this.Where(field, "is", "null", And)
	}
	return this.Where(field, "is", "null", whereType[0])
}

func (this *Builder) OrWhereIsNull(field string) *Builder {
	return this.OrWhere(field, "is", "null")
}

func (this *Builder) OrWhereNotNull(field string) *Builder {
	return this.OrWhere(field, "is not", "null")
}

func (this *Builder) WhereNotNull(field string, whereType ...string) *Builder {
	if len(whereType) == 0 {
		return this.Where(field, "is not", "null", And)
	}
	return this.Where(field, "is not", "null", whereType[0])
}

func (this *Builder) From(table string, as ...string) *Builder {
	if len(as) == 0 {
		this.table = table
	} else {
		this.table = fmt.Sprintf("%s as %s", table, as[0])
	}
	return this
}

func (this *Builder) FromMany(tables ...string) *Builder {
	if len(tables) > 0 {
		this.table = strings.Join(tables, ",")
	}
	return this
}

func (this *Builder) FromSub(provider Provider, as string) *Builder {
	subBuilder := provider()
	this.table = fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as)
	return this.addBinding(fromBinding, subBuilder.GetBindings()...)
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

func (this *Builder) getSelect() string {
	if this.distinct {
		return "DISTINCT " + strings.Join(this.fields, ",")
	}
	return strings.Join(this.fields, ",")
}

func (this *Builder) ToSql() string {
	sql := fmt.Sprintf("SELECT %s FROM %s", this.getSelect(), this.table)

	if !this.joins.IsEmpty() {
		sql = fmt.Sprintf("%s %s", sql, this.joins.String())
	}

	if !this.wheres.IsEmpty() {
		sql = fmt.Sprintf("%s WHERE %s", sql, this.wheres.String())
	}

	if !this.groupBy.IsEmpty() {
		sql = fmt.Sprintf("%s GROUP BY %s", sql, this.groupBy.String())

		if !this.having.IsEmpty() {
			sql = fmt.Sprintf("%s HAVING %s", sql, this.having.String())
		}
	}

	if !this.orderBy.IsEmpty() {
		sql = fmt.Sprintf("%s ORDER BY %s", sql, this.orderBy.String())
	}

	if !this.unions.IsEmpty() {
		sql = fmt.Sprintf("(%s) %s", sql, this.unions.String())
	}

	return sql
}

func (this *Builder) CreateSql(value map[string]interface{}) (sql string, bindings []interface{}) {
	if len(value) == 0 {
		return
	}
	keys := make([]string, 0)

	valuesString := fmt.Sprintf("(%s)", strings.Join(utils.MakeSymbolArray("?", len(value)), ","))
	for name, field := range value {
		bindings = append(bindings, field)
		keys = append(keys, name)
	}

	sql = fmt.Sprintf("insert into %s %s values %s", this.table, fmt.Sprintf("(%s)", strings.Join(keys, ",")), valuesString)
	return
}

func (this *Builder) InsertSql(values []map[string]interface{}) (sql string, bindings []interface{}) {
	if len(values) == 0 {
		return
	}
	fields := make(map[string]interface{})
	valuesString := make([]string, 0)

	for _, value := range values {
		valuesString = append(valuesString, fmt.Sprintf("(%s)", strings.Join(utils.MakeSymbolArray("?", len(value)), ",")))
		for name, field := range value {
			fields[name] = true
			bindings = append(bindings, field)
		}
	}

	fieldsString := ""
	if len(fields) > 0 {
		fieldsString = fmt.Sprintf(" (%s)", strings.Join(utils.GetMapKeys(fields), ","))
	}

	sql = fmt.Sprintf("insert into %s%s values %s", this.table, fieldsString, strings.Join(valuesString, ","))
	return
}
