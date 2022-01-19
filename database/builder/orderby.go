package builder

import (
	"fmt"
	"strings"
)

type orderType string

const (
	Desc orderType = "desc"
	Asc  orderType = "asc"
)

type OrderBy struct {
	field          string
	fieldOrderType orderType
}

type OrderByFields []OrderBy

func (this OrderByFields) IsEmpty() bool {
	return len(this) == 0
}

func (this OrderByFields) String() string {
	if this.IsEmpty() {
		return ""
	}

	columns := make([]string, 0)

	for _, orderBy := range this {
		columns = append(columns, fmt.Sprintf("%s %s", orderBy.field, orderBy.fieldOrderType))
	}

	return strings.Join(columns, ",")
}

func (this *Builder) OrderBy(field string, columnOrderType ...orderType) *Builder {
	if len(columnOrderType) > 0 {
		this.orderBy = append(this.orderBy, OrderBy{
			field:          field,
			fieldOrderType: columnOrderType[0],
		})
	} else {
		this.orderBy = append(this.orderBy, OrderBy{
			field:          field,
			fieldOrderType: Asc,
		})
	}

	return this
}

func (this *Builder) OrderByDesc(field string) *Builder {
	this.orderBy = append(this.orderBy, OrderBy{
		field:          field,
		fieldOrderType: Desc,
	})
	return this
}
