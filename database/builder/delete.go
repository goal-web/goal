package builder

import (
	"fmt"
)

func (this *Builder) DeleteSql() (sql string, bindings []interface{}) {
	sql = fmt.Sprintf("delete from %s", this.table)

	if !this.wheres.IsEmpty() {
		sql = fmt.Sprintf("%s WHERE %s", sql, this.wheres.String())
	}
	bindings = this.GetBindings()
	return
}
