package table

import (
	"github.com/goal-web/contracts"
)

type model struct {
	class      contracts.Class
	table      string
	connection string
}

func Model(class contracts.Class, table string, connection ...string) *Table {
	conn := ""
	if len(connection) > 0 {
		conn = connection[0]
	}
	return FromModel(model{class: class, table: table, connection: conn})
}

func (model model) GetClass() contracts.Class {
	return model.class
}

func (model model) GetTable() string {
	return model.table
}

func (model model) GetConnection() string {
	return model.connection
}
