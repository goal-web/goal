package database

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
)

type Factory struct {
	config      contracts.Config
	connections map[string]contracts.DBConnection
	drivers     map[string]contracts.DBConnector
}

func (this *Factory) Connection(name string) contracts.DBConnection {
	if connection, existsConnection := this.connections[name]; existsConnection {
		return connection
	}

	this.connections[name] = this.make(name)

	return this.connections[name]
}

func (this *Factory) Extend(name string, driver contracts.DBConnector) {
	this.drivers[name] = driver
}

func (this *Factory) make(name string) contracts.DBConnection {
	config := this.config.Get("database").(Config)

	if connectionConfig, existsConnection := config.Connections[name]; existsConnection {
		driverName := utils.GetStringField(connectionConfig, "driver")
		if driver, existsDriver := this.drivers[driverName]; existsDriver {
			return driver(connectionConfig)
		}

		panic(DBConnectionException{
			error:  errors.New("该数据库驱动不存在：" + driverName),
			Code:   DB_DRIVER_DONT_EXIST,
			fields: connectionConfig,
		})
	}

	panic(DBConnectionException{
		error:      errors.New("数据库连接不存在：" + name),
		Code:       DB_CONNECTION_DONT_EXIST,
		Connection: name,
	})
}
