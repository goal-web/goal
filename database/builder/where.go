package builder

import (
	"fmt"
)

type whereJoinType string

const (
	And whereJoinType = "AND"
	Or  whereJoinType = "OR"
)

type Where struct {
	field     string
	condition string
	arg       string
}

func (this *Where) String() string {
	if this == nil {
		return ""
	}
	return fmt.Sprintf("%s %s %s", this.field, this.condition, this.arg)
}

type Wheres struct {
	subWheres map[whereJoinType][]*Wheres
	wheres    map[whereJoinType][]*Where
}

func (this *Wheres) IsEmpty() bool {
	return len(this.subWheres) == 0 && len(this.wheres) == 0
}

func (this Wheres) getSubStringers(whereType whereJoinType) []fmt.Stringer {
	stringers := make([]fmt.Stringer, 0)
	for _, where := range this.subWheres[whereType] {
		stringers = append(stringers, where)
	}
	return stringers
}

func (this Wheres) getStringers(whereType whereJoinType) []fmt.Stringer {
	stringers := make([]fmt.Stringer, 0)
	for _, where := range this.wheres[whereType] {
		stringers = append(stringers, where)
	}
	return stringers
}

func (this *Wheres) getSubWheres(whereType whereJoinType) string {
	return JoinSubStringerArray(this.getSubStringers(whereType), string(whereType))
}

func (this *Wheres) getWheres(whereType whereJoinType) string {
	return JoinStringerArray(this.getStringers(whereType), string(whereType))
}

func (this *Wheres) String() (result string) {
	if this == nil || this.IsEmpty() {
		return ""
	}

	result = this.getSubWheres(And)
	andWheres := this.getWheres(And)

	if result != "" {
		if andWheres != "" {
			result = fmt.Sprintf("%s AND %s", result, andWheres)
		}
	} else {
		result = andWheres
	}

	orSubWheres := this.getSubWheres(Or)
	if result == "" {
		result = orSubWheres
	} else if orSubWheres != "" {
		result = fmt.Sprintf("%s OR %s", result, orSubWheres)
	}

	orWheres := this.getWheres(Or)
	if result == "" {
		result = orWheres
	} else if orWheres != "" {
		result = fmt.Sprintf("%s OR %s", result, orWheres)
	}

	return
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
func (this *Builder) WhereIn(field string, args interface{}) *Builder {
	return this.Where(field, "in", args)
}

func JoinStringerArray(arr []fmt.Stringer, sep string) (result string) {
	for index, stringer := range arr {
		if index == 0 {
			result = stringer.String()
		} else {
			result = fmt.Sprintf("%s %s %s", result, sep, stringer.String())
		}
	}

	return
}

func JoinSubStringerArray(arr []fmt.Stringer, sep string) (result string) {
	for index, stringer := range arr {
		if index == 0 {
			result = fmt.Sprintf("(%s)", stringer.String())
		} else {
			result = fmt.Sprintf("%s %s (%s)", result, sep, stringer.String())
		}
	}

	return
}
