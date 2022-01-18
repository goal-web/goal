package builder

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/utils"
	"strings"
)

type whereJoinType string

const (
	And whereJoinType = "AND"
	Or  whereJoinType = "OR"
)

type Where struct {
	field     string
	condition string
	arg       interface{}
}

func (this *Where) String() string {
	if this == nil {
		return ""
	}
	var stringArg string
	lowerCaseCondition := strings.ToLower(this.condition)

	switch lowerCaseCondition {
	case "in", "not in", "between", "not between":
		isInGrammar := strings.Contains(lowerCaseCondition, "in")
		joinSymbol := utils.IfString(isInGrammar, ",", " AND ")
		switch arg := this.arg.(type) {
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
			stringArg = utils.JoinInterfaceArray(arg, joinSymbol)
		default:
			panic(exceptions.WithError(errors.New("不支持的参数类型"), contracts.Fields{
				"arg":       this.arg,
				"field":     this.field,
				"condition": this.condition,
			}))
		}
		if isInGrammar {
			stringArg = fmt.Sprintf("(%s)", stringArg)
		}
	default:
		stringArg = utils.ConvertToString(this.arg, "")
	}
	if this.condition == "" {
		return fmt.Sprintf("%s %s", this.field, stringArg)
	}
	return fmt.Sprintf("%s %s %s", this.field, this.condition, stringArg)
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
			result = fmt.Sprintf("%s And %s", result, andWheres)
		}
	} else {
		result = andWheres
	}

	orSubWheres := this.getSubWheres(Or)
	if result == "" {
		result = orSubWheres
	} else if orSubWheres != "" {
		result = fmt.Sprintf("%s Or %s", result, orSubWheres)
	}

	orWheres := this.getWheres(Or)
	if result == "" {
		result = orWheres
	} else if orWheres != "" {
		result = fmt.Sprintf("%s OR %s", result, orWheres)
	}

	return
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
