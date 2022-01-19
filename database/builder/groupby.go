package builder

import (
	"strings"
)

type GroupBy []string

func (this GroupBy) IsEmpty() bool {
	return len(this) == 0
}

func (this GroupBy) String() string {
	if this.IsEmpty() {
		return ""
	}

	return strings.Join(this, ",")
}

func (this *Builder) GroupBy(columns ...string) *Builder {
	this.groupBy = append(this.groupBy, columns...)

	return this
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
