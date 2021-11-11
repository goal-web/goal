package contracts

type DBConnectionProvider func(config Fields) DBConnection

type DBFactory interface {
	Connection(key string) DBConnection

	ExtendConnection(driver DBConnectionProvider)
}

type Executor interface {
	First() (map[string]interface{}, error)
	Get() ([]map[string]interface{}, error)
	OrderBy(field, order string)
	Select(fields ...string) Executor
	Where(fields ...string) Executor
	Insert(data ...interface{}) (int64, error)
	Delete() (int64, error)
	Update(data ...interface{}) (int64, error)
}

type DBConnection interface {
	Execute(statement string) (int64, error)
	Query(statement string) (interface{}, error)
	Table(name string) Executor
}
