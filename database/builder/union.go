package builder

import "fmt"

type unionJoinType string

const (
	Union    unionJoinType = "UNION"
	UnionAll unionJoinType = "UNION ALL"
)

type Unions map[unionJoinType][]*Builder

func (this Unions) IsEmpty() bool {
	return len(this) == 0
}

func (this Unions) String() (result string) {
	if this.IsEmpty() {
		return
	}
	for unionType, builders := range this {
		for _, builder := range builders {
			result = fmt.Sprintf("%s %s (%s)", result, unionType, builder.ToSql())
		}
	}

	return
}
