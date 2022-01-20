package builder

import (
	"fmt"
	"github.com/goal-web/supports/utils"
	"strings"
)

type insertType string

const (
	Insert        insertType = "insert"
	InsertIgnore  insertType = "insert ignore"
	InsertReplace insertType = "replace"
)

func getInsertType(insertType2 ...insertType) insertType {
	if len(insertType2) > 0 {
		return insertType2[0]
	}
	return Insert
}

func (this *Builder) CreateSql(value map[string]interface{}, insertType2 ...insertType) (sql string, bindings []interface{}) {
	if len(value) == 0 {
		return
	}
	keys := make([]string, 0)

	valuesString := fmt.Sprintf("(%s)", strings.Join(utils.MakeSymbolArray("?", len(value)), ","))
	for name, field := range value {
		bindings = append(bindings, field)
		keys = append(keys, name)
	}

	sql = fmt.Sprintf("%s into %s %s values %s", getInsertType(insertType2...), this.table, fmt.Sprintf("(%s)", strings.Join(keys, ",")), valuesString)
	return
}

func (this *Builder) InsertSql(values []map[string]interface{}, insertType2 ...insertType) (sql string, bindings []interface{}) {
	if len(values) == 0 {
		return
	}
	fields := utils.GetMapKeys(values[0])
	valuesString := make([]string, 0)

	for _, value := range values {
		valuesString = append(valuesString, fmt.Sprintf("(%s)", strings.Join(utils.MakeSymbolArray("?", len(value)), ",")))
		for _, field := range fields {
			bindings = append(bindings, value[field])
		}
	}

	fieldsString := fmt.Sprintf(" (%s)", strings.Join(fields, ","))

	sql = fmt.Sprintf("%s into %s%s values %s", getInsertType(insertType2...), this.table, fieldsString, strings.Join(valuesString, ","))
	return
}

func (this *Builder) InsertIgnoreSql(values []map[string]interface{}) (sql string, bindings []interface{}) {
	return this.InsertSql(values, InsertIgnore)
}

func (this *Builder) InsertReplaceSql(values []map[string]interface{}) (sql string, bindings []interface{}) {
	return this.InsertSql(values, InsertReplace)
}
