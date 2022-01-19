package builder

func (this *Builder) WhereIsNull(field string, whereType ...string) *Builder {
	if len(whereType) == 0 {
		return this.Where(field, "is", "null")
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
		return this.Where(field, "is not", "null")
	}
	return this.Where(field, "is not", "null", whereType[0])
}
