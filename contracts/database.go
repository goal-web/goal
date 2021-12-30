package contracts

type DBConnector func(config Fields) DBConnection

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type DBFactory interface {
	Connection(key string) DBConnection

	ExtendConnection(name string, driver DBConnector)
}

type DBConnection interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (Result, error)
	DriverName() string
}
