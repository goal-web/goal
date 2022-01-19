package builder

import "fmt"

type unionJoinType string

const (
	Union    unionJoinType = "union"
	UnionAll unionJoinType = "union all"
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
