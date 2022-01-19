package builder

func (this *Builder) WhereIn(field string, args interface{}) *Builder {
	return this.Where(field, "in", args)
}
func (this *Builder) OrWhereIn(field string, args interface{}) *Builder {
	return this.OrWhere(field, "in", args)
}

func (this *Builder) WhereNotIn(field string, args interface{}) *Builder {
	return this.Where(field, "not in", args)
}

func (this *Builder) OrWhereNotIn(field string, args interface{}) *Builder {
	return this.OrWhere(field, "not in", args)
}
