package builder

import "fmt"

func (this *Builder) Select(field string, fields ...string) *Builder {
	this.fields = append(fields, field)
	return this
}

func (this *Builder) AddSelect(fields ...string) *Builder {
	this.fields = append(this.fields, fields...)
	return this
}

func (this *Builder) SelectSub(provider Provider, as string) *Builder {
	subBuilder := provider()
	this.fields = []string{fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as)}
	return this.addBinding(selectBinding, subBuilder.GetBindings()...)
}
func (this *Builder) AddSelectSub(provider Provider, as string) *Builder {
	subBuilder := provider()
	this.fields = append(this.fields, fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as))
	return this.addBinding(selectBinding, subBuilder.GetBindings()...)
}

func (this *Builder) Count(fields ...string) *Builder {
	if len(fields) == 0 {
		return this.Select("count(*)")
	}
	return this.Select(fmt.Sprintf("count(%s) as %s_count", fields[0], fields[0]))
}

func (this *Builder) Avg(field string, as ...string) *Builder {
	if len(as) == 0 {
		return this.Select(fmt.Sprintf("avg(%s) as %s_avg", field, field))
	}
	return this.Select(fmt.Sprintf("avg(%s) as %s", field, as[0]))
}

func (this *Builder) Sum(field string, as ...string) *Builder {
	if len(as) == 0 {
		return this.Select(fmt.Sprintf("sum(%s) as %s_sum", field, field))
	}
	return this.Select(fmt.Sprintf("sum(%s) as %s", field, as[0]))
}

func (this *Builder) Max(field string, as ...string) *Builder {
	if len(as) == 0 {
		return this.Select(fmt.Sprintf("max(%s) as %s_max", field, field))
	}
	return this.Select(fmt.Sprintf("max(%s) as %s", field, as[0]))
}

func (this *Builder) Min(field string, as ...string) *Builder {
	if len(as) == 0 {
		return this.Select(fmt.Sprintf("min(%s) as %s_min", field, field))
	}
	return this.Select(fmt.Sprintf("min(%s) as %s", field, as[0]))
}
