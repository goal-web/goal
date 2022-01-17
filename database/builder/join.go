package builder

import "fmt"

type joinType string

const (
	LeftJoin    joinType = "LEFT"
	RightJoin   joinType = "RIGHT"
	InnerJoin   joinType = "INNER"
	FullOutJoin joinType = "FULL OUTER"
	FullJoin    joinType = "FULL"
)

type Join struct {
	table      string
	join       joinType
	conditions *Wheres
}

func (this Join) String() (result string) {
	result = fmt.Sprintf("%s JOIN %s", this.join, this.table)
	if this.conditions.IsEmpty() {
		return
	}
	result = fmt.Sprintf("%s ON (%s)", result, this.conditions.String())
	return
}

type Joins []Join

func (this Joins) IsEmpty() bool {
	return len(this) == 0
}

func (this Joins) String() (result string) {
	if this.IsEmpty() {
		return
	}

	for index, join := range this {
		if index == 0 {
			result = join.String()
		} else {
			result = fmt.Sprintf("%s%s%s", result, ",", join.String())
		}
	}

	return
}
