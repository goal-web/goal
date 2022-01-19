package builder

import "fmt"

func (this *Builder) WhereExists(provider Provider, where ...whereJoinType) *Builder {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())
	if len(where) == 0 {
		return this.Where("", "exists", subSql).
			addBinding(whereBinding, subBuilder.GetBindings()...)
	}

	return this.Where("", "exists", subSql, where[0]).
		addBinding(whereBinding, subBuilder.GetBindings()...)
}

func (this *Builder) OrWhereExists(provider Provider) *Builder {
	return this.WhereExists(provider, Or)
}

func (this *Builder) WhereNotExists(provider Provider, where ...whereJoinType) *Builder {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())
	if len(where) == 0 {
		return this.Where("", "not exists", subSql).
			addBinding(whereBinding, subBuilder.GetBindings()...)
	}

	return this.Where("", "not exists", subSql, where[0]).
		addBinding(whereBinding, subBuilder.GetBindings()...)
}

func (this *Builder) OrWhereNotExists(provider Provider) *Builder {
	return this.WhereNotExists(provider, Or)
}
