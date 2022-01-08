package contracts

type DBConnector func(config Fields) DBConnection

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type DBFactory interface {
	Connection(key string) DBConnection
	Extend(name string, driver DBConnector)
}

type DBTx interface {
	SqlExecutor
	Commit() error
	Rollback() error
}

type SqlExecutor interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (Result, error)
}

type DBConnection interface {
	SqlExecutor
	Begin() (DBTx, error)
	Transaction(func(executor SqlExecutor) error) error
	DriverName() string
}
