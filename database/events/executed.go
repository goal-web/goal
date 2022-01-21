package events

type QueryExecuted struct {
	Sql        string
	Bindings   []interface{}
	Connection string
}

func (this *QueryExecuted) Event() string {
	return "QUERY_EXECUTED"
}

func (this *QueryExecuted) Sync() bool {
	return true
}
