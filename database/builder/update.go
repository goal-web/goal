package builder

import (
	"fmt"
	"strings"
)

func (this *Builder) UpdateSql(value map[string]interface{}) (sql string, bindings []interface{}) {
	if len(value) == 0 {
		return
	}
	valuesString := make([]string, 0)
	for name, field := range value {
		valuesString = append(valuesString, fmt.Sprintf("%s = ?", name))
		bindings = append(bindings, field)
	}

	sql = fmt.Sprintf("update %s set %s", this.table, strings.Join(valuesString, ","))

	if !this.wheres.IsEmpty() {
		sql = fmt.Sprintf("%s where %s", sql, this.wheres.String())
	}

	return
}
