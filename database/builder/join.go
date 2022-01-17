package builder

import "fmt"

type joinType string

const (
	LeftJoin    joinType = "LEFT"
	RightJoin   joinType = "RIGHT"
	InnerJoin   joinType = "INNER"
	FullOutJoin joinType = "FULL OUTER"
)

type Join struct {
	table      string
	join       joinType
	conditions *Wheres
}

func (this *Join) String() (result string) {
	result = fmt.Sprintf("%s JOIN %s", this.join, this.table)
	if this.conditions.IsEmpty() {
		return
	}
	result = fmt.Sprintf("%s ON (%s)", result, this.conditions.String())
	return
}
